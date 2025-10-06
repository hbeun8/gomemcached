package connectionhandler

import (
	"flag"
	"fmt"
	commandhandler "gomemc/command_handler"
	serializationhandler "gomemc/serialization_handler"
	datastorehandler "gomemc/datastore_handler"
	"io"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func Connect() {
	fmt.Println("Go Server")

	var host, port string
	flag.StringVar(&host, "host", "localhost", "Target server address/name")
	flag.StringVar(&port, "port", "4200", "Target server port")
	flag.Parse()

	// Create a new server instance
	s, err := NewServer(host, port)
	if err != nil {
		slog.Error("Unable to start server", "Error", err)
	}

	// Start the server
	s.Start()

	// Pause the main thread waiting for a SIGINT or SIGTERM signal to gracefully shut down the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	// Shut down the server
	slog.Info("Shutting down server...")
	s.Stop()
	slog.Info("Server stopped.")
}

type Server struct {
	Host     string
	Port     int
	listener net.Listener
	shutdown chan any
}

func NewServer(host string, port string) (*Server, error) {
	p, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf("%s:%d", host, p)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		Host:     host,
		Port:     p,
		listener: l,
		shutdown: make(chan interface{}),  // any is the same as interface{}
	}, nil
}

func (s *Server) Start() {
	// Start a goroutine to accept connections
	go s.acceptConnections()
}

func (s *Server) Stop() {
	// close the shutdown channel.
	close(s.shutdown)

	// Close the listener
	err := s.listener.Close()
	if err != nil {
		slog.Error("error closing listener", "Error", err)
	}
}

func (s *Server) acceptConnections() {
	for {
		c, err := s.listener.Accept()
		if err != nil {
			select {
			case <-s.shutdown:
				// end accepting connections when the shutdown channel is closed
				return
			default:
				log.Println("Accept error: ", err)
			}
		}
		go func() {
			s.handleConnection(c)
		}()
	}
}


func (s *Server) handleConnection(c net.Conn) {
	defer c.Close()

	// big buffer for assembling message
	buf := make([]byte, 0, 4096)

	// short read buffer
	rbuf := make([]byte, 10)

	//datastore
	m := make(map[string][][]byte)
	d := datastorehandler.Datastore{Datastore: m}
	for {
		select {
		case <-s.shutdown:
			// end handling the connection when the shutdown channel is closed
			log.Println("Closing client connection")
			return
		default:
			n, err := c.Read(rbuf)
			if n > 0 {
				buf = append(buf, rbuf[:n]...)

				// TODO Handle the protocol and remove consumed bytes from buf
				//a:= fmt.Sprintf("%q", buf)
				//fmt.Println(a)
				commandline, datablock:=commandhandler.Sifter(buf) 
				var block [][]byte
				if len(commandline) > 0 {
					block = append(block, commandline)
				}
				if len(datablock) > 0 {
					block = append(block, datablock)
				}
				if len(block) == 2 {
					message:=commandhandler.Command(commandline, datablock, &d)
					//preparing for the next block
					clear(commandline)
					clear(datablock)
					clear(block)
					c.Write(serializationhandler.Serializer2(message))
				}

			}
			if err != nil {
				if err == io.EOF {
					log.Println("Connection closed by client")
					return
				}
				log.Println("Error reading from connection:", err)
				return
			}
		}
	}
}
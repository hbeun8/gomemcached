package connectionhandler

import (
	"net"
	"log"
	"flag"
	"fmt"
)

func NaiveServer() {

	flagp := flag.Bool("p", false, "Port, defaults to 11211" )
	flag.Parse()
	port:=flag.CommandLine.Args()
	listener, err := net.Listen("tcp", "localhost:11211")
	if *flagp {
		listener, err = net.Listen("tcp", "localhost:"+port[0])	//127.0.0.1 
	}
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		
		go func (c net.Conn) {
			buf := make([]byte, 1024)
			_, err := c.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("starting on the default port")
			log.Print(string(buf))
			conn.Write([]byte("Hello from the TCP Server"))
			c.Close()
		}(conn)
	}
}
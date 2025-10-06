# Memcached Server in Go

A simple Memcached server implementation.

## Project Structure

The project is organized into the following components:

```
gomemc/
│── main.py                              
│   ├── command_handler/
│   │   │── command_test.go
│   │   │── command.go
│   ├── connection_handler/
│   │   │── connection.go
│   │   │── naiveserver.go    <-not being used>
│   ├── datastore_handler/
│   │   │── datastore_test.go
│   │   │── datastore.go
│   ├── protocol_handler/
│   │   │── protocol1_test.go <-not being used>
│   │   │── protocol1.go      <-not being used>
│   │   │── protocol2_test.go
│   │   │── protocol2.go
│   ├── serialization_handler/
│   │   │── serializer.go
└── go.mod

```
## Features:
1. Supported Commands: SET, GET, DEL, ADD, REPLACE, APPEND, PREPEND.
2. Set Key Expiry
3. Persistence through Append only File.
4. Concurrency through Goconcurrency and Channels
5. Tests


## Components

1. **Main Component** (`main.go`):
   - Entry point
   - Connects to Server

2. **Server Component** (`connection.go`):
   - Manages socket creation, listens, accept, reads and writes to client
   - Handles TCP
   - Coroutines for concurrent client connections
   - Handles server restart, shutdown
   - Coroutines for concurrent access to datastore

3. **Command Handler Component** (`command.go`):
   - Processes client data
   - Implements command handlers (set, get and etc)

4. **Datastore Handler Component** (`datastore.go`):
   - Stores data
   - Implements methods (set, get and etc)

5. **Protocol Handler Component** (`protocol.go`):
   - Processes client data
   - Implemts Parsers for each method (set, get and etc)

6. **Serialization Handler Component** (`serializer.go`):
   - Processes return client data
   - Decodes buffer and writes to commandline

## Testing Approach

The project uses pytest for testing with a hybrid approach:

1. **Unit Tests**:
   - Test individual components in isolation
   - Focus on specific functionality

## Running Tests

To run all tests:

```bash
go test -v
```

To run specific test categories:
TBC

## Running the Server

Start the server with various options:

```bash
# Run the Execuatble
./gomemc 

# Run main.go
go run main.go
# After go install
gomemc

```
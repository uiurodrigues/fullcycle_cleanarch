# Clean Arch

Welcome to CleanArch Application! This application its an example of an clean arch implementation.

## Usage

1. Run the command [docker-compose up -d] to build the environment (sql and rabbitMQ)
2. Run the command [cd cmd/ordersystem] to go to the system path
3. Run the command [go run main.go wire_gen.go] to execute the program


## Web Server
- It's running on port 8000
- Use the files oh the path /api to execute the commands


## GRPC Server
- It's running on port 50051
- Execute the command [evans -r repl] to access the application and see the Cr
- Execute the command [package pb]
- Execute the command [service OrderService]

1. Execute the command [call CreateOrder] to create a new Order
2. Execute the command [call ListOrder] to list all created Orders


## Graphqlql
- It's running on port 8080
- To access it call http://localhost:8080 and access the playground



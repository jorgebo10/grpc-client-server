# grpc-client-server
Demo for grpc client and server

Go to project root folder and run the code below to create server and client stubs

protoc -I ./proto --go_out=./server  \
    --go-grpc_out=./server \
    product_info.proto

protoc -I ./proto --go_out=./client  \
    --go-grpc_out=./client \
    product_info.proto

cd server
go build -v -o bin/server

cd client
go build -v -o bin/client

run
bin/server to start the server
and 
bin/client to start the client

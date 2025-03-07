protoc --go_out=. --go-grpc_out=. ./proto/auth.proto  

docker run -d -p 8123:8123 -p 9000:9000 --name clickhouse-server clickhouse/clickhouse-server
docker start clickhouse-server

go run main.go http_gateway.go grpc_server.go http.go handlers.go
go run main.go grpc_server.go handlers.go
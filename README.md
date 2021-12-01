# Miniproject3
Auction house using golang and grpc with passive replication
# Run project
To run, run commands in this order from this Miniproject3 directory

Start 5 replication nodes
```
go run server/server.go 6001
go run server/server.go 6002
go run server/server.go 6003
go run server/server.go 6004
go run server/server.go 6005
```

Start 2 clients
```
go run client/client.go client/frontend.go
go run client/client.go client/frontend.go
```
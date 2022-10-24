# go-grpc

## proto files
pb list: `common/model/`

## Generate files
from `.proto` files to `.pb.go` files
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./common/model/*.proto
```

## Run Server
```
go run ./services/service-user/main.go
```

## Run Client
```
go run ./client/main.go
```
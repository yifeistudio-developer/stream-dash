# Swift

**Swift** 是 StreamDash 的通信层，基于 Go 和 gRPC，负责数据源接入和推送。

## 文件说明

- **main.go**: gRPC 服务端，推送模拟数据。
- **proto/**: 从 Fox 生成的 Protobuf 和 gRPC 代码。

## 初始化

1. 初始化模块：
   ```bash
   go mod init streamdash/swift
   go get google.golang.org/grpc
   go get google.golang.org/protobuf

2. 生成 Protobuf 和 gRPC 代码：
   ```bash
   cd ../fox
   protoc --go_out=../swift/proto --go_opt=paths=source_relative --go-grpc_out=../swift/proto --go-grpc_opt=paths=source_relative data.proto
   ```
## 运行

```bash
  go run main.go
```
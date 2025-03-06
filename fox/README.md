# Fox

**Fox** 是 StreamDash 项目的 Protobuf 定义模块，负责跨组件共享的数据格式。

## 文件说明

- **data.proto**: 定义数据点 (DataPoint)、数据流 (DataStream) 和 gRPC 服务 (StreamService)。

## 生成代码

### Go

```bash
protoc --go_out=../swift/proto --go_opt=paths=source_relative --go-grpc_out=../swift/proto --go-grpc_opt=paths=source_relative data.proto
```

### Java/Kotlin (Flink)

```bash
protoc --java_out=../otter/src/main/java data.proto
```

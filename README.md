# StreamDash

**StreamDash** 是一款实时数据仪表盘生成器，旨在让用户通过简单配置即可从流式或静态数据源生成动态可视化仪表盘。它结合了高性能通信、分布式流处理和现代前端技术，追求各技术栈在特定领域的极致发挥。

## 项目概述

- **目标**: 提供一个高效、易用的工具，自动解析数据源并生成实时仪表盘。
- **适用场景**: 网站流量监控、传感器数据分析、财务指标展示等。
- **特性**:
  - 支持多种数据源（Kafka、CSV 等）。
  - 拖拽式界面，自动推荐图表类型。
  - 实时更新，支持过滤和聚合。
  - 可导出为独立 Web 组件。

## 技术栈

| **领域**       | **技术**                     | **优势**                          |
|----------------|------------------------------|-----------------------------------|
| 通信层         | Go + gRPC + Protobuf         | 高并发、低延迟、高效序列化        |
| 数据处理       | Flink Stateful Functions     | 实时流处理、状态管理              |
| 前端           | Vue 3 + WASM + Canvas        | 高性能渲染、响应式体验            |
| 存储           | Kafka                        | 分布式日志、数据流统一            |
| 数据格式       | Protobuf                     | 跨语言、高效传输                  |

## 模块代号

项目分为以下模块，每个模块以动物命名，便于管理和识别：

- **Swift（雨燕）**: 通信层，基于 Go + gRPC，负责数据源接入和推送。
- **Otter（水獭）**: 数据处理层，基于 Flink Stateful Functions，处理流数据和状态。
- **Hawk（鹰）**: 前端，基于 Vue 3 + WASM + Canvas，渲染实时仪表盘。
- **Beaver（河狸）**: 存储层，使用 Kafka，提供数据缓冲和传递。
- **Fox（狐狸）**: Protobuf 定义，跨组件共享的数据格式。

## 目录结构

StreamDash/
├── swift/         # Go 通信层 (Swift)
├── otter/         # Flink 数据处理层 (Otter)
├── hawk/          # Vue 3 前端 (Hawk)
├── beaver/        # Kafka 配置 (Beaver)
├── fox/           # Protobuf 定义 (Fox)
└── README.md

## 架构图

```
+----------------+       +----------------+       +----------------+
|   数据源       | ----> |   Swift (gRPC) | ----> |   Beaver       |
| (Kafka, CSV等) |       | (Protobuf)     |       | (Kafka)        |
+----------------+       +----------------+       +----------------+
                                    |
                                    v
+----------------+       +----------------+
|   Otter (Flink)| <---- |   Beaver       |
| (Stateful)     |       | (Protobuf)     |
+----------------+       +----------------+
                                    |
                                    v
+----------------+       +----------------+
|   Hawk (Vue 3) | <---- |   gRPC-Web     |
| (WASM + Canvas)|       | (Protobuf)     |
+----------------+       +----------------+
```

- **Swift**: 从数据源读取，推送 Protobuf 数据到 Kafka。
- **Beaver**: 存储和分发数据。
- **Otter**: 从 Kafka 消费，处理并输出结果。
- **Hawk**: 通过 gRPC-Web 接收数据，渲染仪表盘。

## 安装与配置

### 前置条件

- **Go**: 1.21+
- **Java**: 11+（Flink 和 Kotlin）
- **Node.js**: 18+（Vue 3）
- **Kafka**: 3.6+（带 ZooKeeper）
- **Rust**: 1.75+（WASM 可选）
- **protoc**: 3.25+（Protobuf 编译器）

### 安装步骤

1. **克隆仓库**:
   ```bash
   git clone https://github.com/your-repo/StreamDash.git
   cd StreamDash
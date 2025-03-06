# Beaver

**Beaver** 是 StreamDash 的存储层，使用 Kafka 提供数据缓冲和传递。

## 文件说明

- **kafka-config.yaml**: Kafka 服务端配置文件。

## 初始化

1. 下载 Kafka（例如 3.6.1）：
   ```bash
   wget https://archive.apache.org/dist/kafka/3.6.1/kafka_2.13-3.6.1.tgz
   tar -xzf kafka_2.13-3.6.1.tgz

2. 配置
    - 复制 kafka-config.yaml 到 Kafka 目录（如 kafka_2.13-3.6.1/config/server.properties）。

## 运行

1. 启动ZooKeeper
    ```bash
    cd kafka_2.13-3.6.1
    bin/zookeeper-server-start.sh config/zookeeper.properties &

2. 启动 Kafka：
    ```bash
    bin/kafka-server-start.sh config/server.properties &

3. 创建主题
    ```
    bin/kafka-topics.sh --create --topic input-topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1

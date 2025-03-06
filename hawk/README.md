# Hawk

**Hawk** 是 StreamDash 的前端，基于 Vue 3，使用 ECharts 渲染实时仪表盘。

## 文件说明

- **package.json**: Node.js 依赖配置文件。
- **vite.config.js**: Vite 构建配置。
- **src/App.vue**: 主组件，显示图表。
- **src/main.js**: Vue 入口文件。
- **src/proto/**: Protobuf 定义（从 Fox 复制）。

## 初始化

1. 安装依赖：
   ```bash
   npm install

2. 复制 Protobuf 文件：
    ```bash
    cp ../fox/data.proto src/proto/

## 运行

    ```bash
    yarn run dev
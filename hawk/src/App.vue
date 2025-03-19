<template>
  <div class="container">
    <h1>StreamDash - Hawk</h1>
    <v-chart :option="chartOption" style="height: 400px;" />
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import VChart from 'vue-echarts';
import { use } from 'echarts/core';
import { LineChart } from 'echarts/charts';
import { GridComponent, TooltipComponent } from 'echarts/components';
import { CanvasRenderer } from 'echarts/renderers';
import { grpc } from '@improbable-eng/grpc-web';
import protobuf from 'protobufjs';

use([LineChart, GridComponent, TooltipComponent, CanvasRenderer]);

const chartOption = ref({
  xAxis: { type: 'category', data: [] },
  yAxis: { type: 'value' },
  series: [{ type: 'line', data: [] }],
  tooltip: { trigger: 'axis' }
});

let client = null;

onMounted(async () => {
  // 加载 proto 文件
  const root = await protobuf.load('/src/proto/data.proto');
  const DataStream = root.lookupType('streamdash.DataStream');

  // 构造服务定义
  const StreamServiceDefinition = {
    serviceName: 'streamdash.StreamService', // 必须匹配 proto 中的服务名
    StreamData: {
      methodName: 'StreamData',
      service: { serviceName: 'streamdash.StreamService' }, // 服务名
      requestStream: false, // 单向流
      responseStream: true, // 服务器流
      requestType: DataStream,
      responseType: DataStream
    }
  };

  client = grpc.client(StreamServiceDefinition.StreamData, {
    host: 'http://localhost:8080',
    transport: grpc.WebsocketTransport()
  });

  client.onMessage((message) => {
    const points = message.getPointsList();
    points.forEach(point => {
      chartOption.value.xAxis.data.push(point.getTimestamp());
      chartOption.value.series[0].data.push(point.getValue());
    });
    if (chartOption.value.xAxis.data.length > 20) {
      chartOption.value.xAxis.data.shift();
      chartOption.value.series[0].data.shift();
    }
    chartOption.value = { ...chartOption.value };
  });

  client.onEnd((status, statusMessage) => {
    console.log('gRPC stream ended:', status, statusMessage);
  });

  client.

  // client.onError((err) => {
  //   console.error('gRPC error:', err);
  // });

  client.start();
});

onUnmounted(() => {
  // if (client) client.finish();
});
</script>

<style scoped>
.container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}
</style>

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
import * as protobuf from 'protobufjs';

use([LineChart, GridComponent, TooltipComponent, CanvasRenderer]);

const chartOption = ref({
  xAxis: { type: 'category', data: [] },
  yAxis: { type: 'value' },
  series: [{ type: 'line', data: [] }],
  tooltip: { trigger: 'axis' }
});

let ws = null;

onMounted(async () => {
  const root = await protobuf.load('/proto/data.proto');
  const DataStream = root.lookupType('streamdash.DataStream');

  ws = new WebSocket('ws://localhost:8080/data'); // 临时用 WebSocket，后续改为 gRPC-Web
  ws.binaryType = 'arraybuffer';
  ws.onmessage = (event) => {
    const buffer = new Uint8Array(event.data);
    const message = DataStream.decode(buffer);
    message.points.forEach(point => {
      chartOption.value.xAxis.data.push(point.timestamp);
      chartOption.value.series[0].data.push(point.value);
    });
    if (chartOption.value.xAxis.data.length > 20) {
      chartOption.value.xAxis.data.shift();
      chartOption.value.series[0].data.shift();
    }
    chartOption.value = { ...chartOption.value };
  };
});

onUnmounted(() => {
  if (ws) ws.close();
});
</script>

<style scoped>
.container {
  padding: 20px;
  max-width: 800px;
  margin: 0 auto;
}
</style>

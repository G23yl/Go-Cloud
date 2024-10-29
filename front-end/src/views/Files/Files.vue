<script setup lang="ts">
import Title from "@/components/Title/Title.vue"
import { ref } from "vue"

const title = ref("我的文件")
// 获取到uri中的path参数
const { query } = defineProps<{
  query: string
}>()
</script>

<template>
  <el-container style="margin-left: 20px">
    <el-header>
      <Title :title="title" :search="false" :bread="true" :path="query" />
    </el-header>
    <el-divider style="width: 99%"></el-divider>
    <el-main> </el-main>
  </el-container>
</template>

<style scoped lang="scss">
.el-main {
  --el-main-padding: 0;
}
</style>

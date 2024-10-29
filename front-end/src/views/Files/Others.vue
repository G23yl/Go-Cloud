<script setup lang="ts">
import { onBeforeMount, ref } from "vue"
import Title from "@/components/Title/Title.vue"
import type { DIVAOData } from "@/types/types"
import { usePageData } from "@/hooks/pages"
import Table from "@/components/Table/Table.vue"

onBeforeMount(async () => {
  const { docs } = usePageData()
  try {
    const res = await docs()
    if (res) {
      data = res.data
    }
  } catch (error) {}
})
let data: DIVAOData[] | undefined
const title = ref("其他文件")
</script>

<template>
  <el-container style="margin-left: 20px">
    <el-header>
      <Title :title="title" :search="false" />
    </el-header>
    <el-divider style="width: 99%"></el-divider>
    <el-main>
      <Table :data="data ? data : []" :icon="['fas', 'circle-question']" />
    </el-main>
  </el-container>
</template>

<style scoped lang="scss">
.el-main {
  --el-main-padding: 0;
}
</style>

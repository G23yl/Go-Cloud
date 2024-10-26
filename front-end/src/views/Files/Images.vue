<script setup lang="ts">
import { onBeforeMount, reactive, ref } from "vue"
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
const title = ref("我的图像")
// const test_data = reactive<DIVAOData[]>([
//   {
//     fileID: 1,
//     fileName: "test1",
//     fileSize: 100,
//     downloadNum: 5,
//     updateTime: "2014-09-12",
//     fileUrl: "xxx",
//   },
//   {
//     fileID: 2,
//     fileName: "test2",
//     fileSize: 100,
//     downloadNum: 0,
//     updateTime: "2014-09-13",
//     fileUrl: "xxx",
//   },
//   {
//     fileID: 2,
//     fileName: "test2",
//     fileSize: 100,
//     downloadNum: 0,
//     updateTime: "2014-09-13",
//     fileUrl: "xxx",
//   },
//   {
//     fileID: 2,
//     fileName: "test2",
//     fileSize: 100,
//     downloadNum: 0,
//     updateTime: "2014-09-13",
//     fileUrl: "xxx",
//   },
//   {
//     fileID: 2,
//     fileName: "test2",
//     fileSize: 100,
//     downloadNum: 0,
//     updateTime: "2014-09-13",
//     fileUrl: "xxx",
//   },
//   {
//     fileID: 2,
//     fileName: "test2",
//     fileSize: 100,
//     downloadNum: 0,
//     updateTime: "2014-09-13",
//     fileUrl: "xxx",
//   },
//   {
//     fileID: 2,
//     fileName: "test2",
//     fileSize: 100,
//     downloadNum: 0,
//     updateTime: "2014-09-13",
//     fileUrl: "xxx",
//   },
//   {
//     fileID: 2,
//     fileName: "test2",
//     fileSize: 100,
//     downloadNum: 0,
//     updateTime: "2014-09-13",
//     fileUrl: "xxx",
//   },
//   {
//     fileID: 2,
//     fileName: "test2",
//     fileSize: 100,
//     downloadNum: 0,
//     updateTime: "2014-09-13",
//     fileUrl: "xxx",
//   },
//   {
//     fileID: 2,
//     fileName: "test2",
//     fileSize: 100,
//     downloadNum: 0,
//     updateTime: "2014-09-13",
//     fileUrl: "xxx",
//   },
//   {
//     fileID: 2,
//     fileName: "test2",
//     fileSize: 100,
//     downloadNum: 0,
//     updateTime: "2014-09-13",
//     fileUrl: "xxx",
//   },
//   {
//     fileID: 2,
//     fileName: "test2",
//     fileSize: 100,
//     downloadNum: 0,
//     updateTime: "2014-09-13",
//     fileUrl: "xxx",
//   },
// ])
</script>

<template>
  <el-container style="margin-left: 20px">
    <el-header>
      <Title :title="title" />
    </el-header>
    <el-divider style="width: 99%"></el-divider>
    <el-main>
      <Table :data="data ? data : []" :icon="['fas', 'image']" />
      <!-- <Table :data="test_data" :icon="['fas', 'image']" /> -->
    </el-main>
  </el-container>
</template>

<style scoped lang="scss">
.el-main {
  --el-main-padding: 0;
}
</style>

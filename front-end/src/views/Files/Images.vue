<script setup lang="ts">
import { onBeforeMount, ref } from "vue"
import Title from "@/components/Title/Title.vue"
import type { DIVAOData } from "@/types/types"
import { usePageData } from "@/hooks/pages"
import Table from "@/components/Table/Table.vue"
import { getFileSizeStr } from "@/utils/util/util"

let data = ref<DIVAOData[]>()
onBeforeMount(async () => {
  const { images } = usePageData()
  try {
    const res = await images()
    if (res) {
      data.value = res.data
      data.value.forEach((item) => {
        item.fileSizeStr = getFileSizeStr(item.fileSize)
      })
    }
  } catch (error) {}
})

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
      <Title :title="title" :search="false" />
    </el-header>
    <el-divider style="width: 99%"></el-divider>
    <el-main>
      <Table :data="data ? data : []" :icon="['fas', 'image']" />
    </el-main>
  </el-container>
</template>

<style scoped lang="scss">
.el-main {
  --el-main-padding: 0;
}
</style>

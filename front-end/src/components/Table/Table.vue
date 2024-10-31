<script setup lang="ts">
import type { DIVAOData, FFData } from "@/types/types"
import { Delete, Download } from "@element-plus/icons-vue"
import { onMounted, ref } from "vue"

interface Props {
  data: DIVAOData[] | FFData[]
  size?: string
  color?: string
}
const { data, size = "lg", color = "#74c0fc" } = defineProps<Props>()
const emit = defineEmits<{
  deleteFile: [fileName: string, fileID: number]
}>()
// 实时更新表格高度
let tableHeight = ref(document.body.clientHeight - 135)
// 获取鼠标hover在哪一行
const hoverLine = ref<number>()

onMounted(() => {
  tableHeight.value = document.body.clientHeight - 135
})
window.addEventListener("resize", () => {
  tableHeight.value = document.body.clientHeight - 135
})
const hover = (row: any) => {
  hoverLine.value = row.fileID
}
const unhover = () => {
  hoverLine.value = undefined
}
const sortBySize = (a: any, b: any) => {
  if (a.fileSize === b.fileSize) {
    return 0
  }
  return a.fileSize > b.fileSize ? 1 : -1
}
const sortByCount = (a: any, b: any) => {
  if (a.downloadNum === b.downloadNum) {
    return 0
  }
  return a.downloadNum > b.downloadNum ? 1 : -1
}
const sortByDate = (a: any, b: any) => {
  const dateA = new Date(a.updateTime)
  const dateB = new Date(b.updateTime)
  if (dateA === dateB) {
    return 0
  }
  return dateA > dateB ? 1 : -1
}
// 删除文件操作
// 需要知道删除文件的路径，删除文件的文件名和文件ID
</script>

<template>
  <div class="content">
    <el-table
      :data="data"
      style="width: 98%"
      stripe
      :max-height="tableHeight"
      table-layout="auto"
      @cell-mouse-enter="hover"
      @cell-mouse-leave="unhover"
    >
      <!-- <el-table-column type="selection" width="30px" /> -->
      <el-table-column label="名称" show-overflow-tooltip>
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <font-awesome-icon :icon="scope.row.fileIcon" :size="size" :style="{ color: color }" />
            <span style="margin-left: 10px">{{ scope.row.fileName }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="大小" width="180px" sortable :sort-method="sortBySize">
        <template #default="scope">
          <el-tag type="primary" round>{{ scope.row.fileSizeStr }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="下载次数" width="180px" sortable :sort-method="sortByCount">
        <template #default="scope">
          <el-tag type="success" round>{{ scope.row.downloadNum }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="修改日期" sortable :sort-method="sortByDate">
        <template #default="scope">
          <el-tag type="info" round>{{ new Date(scope.row.updateTime).toLocaleString() }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <div v-show="scope.row.fileID === hoverLine">
            <el-button-group>
              <el-button
                type="danger"
                :icon="Delete"
                size="small"
                @click="emit('deleteFile', scope.row.fileName, scope.row.fileID)"
              ></el-button>
              <el-button type="primary" :icon="Download" size="small"></el-button>
            </el-button-group>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<style scoped lang="scss">
.icon {
  cursor: pointer;
  margin-right: 10px;
  &:focus {
    outline: none;
  }
}
</style>

<script setup lang="ts">
import useFileStore from "@/store/files"
import type { DIVAOData, FFData } from "@/types/types"
import { Delete, Download } from "@element-plus/icons-vue"
import { onMounted, ref, watch } from "vue"
import { useRoute, useRouter } from "vue-router"

interface Props {
  data: DIVAOData[] | FFData[] | undefined
  size?: string
  // 用于判断是FFData类型还是DIVAOData类型，是DIVAO类型就显示tag这一栏，用来标识该文件的父文件夹是谁
  isdorf?: string
}
const { data, size = "lg", isdorf = "D" } = defineProps<Props>()
const emit = defineEmits<{
  deleteFile: [filePath: string, fileName: string, fileID: number, type: string]
  download: [filePath: string, fileName: string]
}>()
// 实时更新表格高度
let tableHeight = ref(document.body.clientHeight - 135)
// 获取鼠标hover在哪一行
const hoverLine = ref<number>()
const hoverType = ref<string>()
const route = useRoute()
const router = useRouter()
const tableLoading = ref(true)
const { setFiles } = useFileStore()

// data还未获取的时候表格显示加载中
watch(
  () => data,
  () => {
    if (data) {
      tableLoading.value = false
    }
  }
)
// 实时更新表格高度
onMounted(() => {
  tableHeight.value = document.body.clientHeight - 135
})
window.addEventListener("resize", () => {
  tableHeight.value = document.body.clientHeight - 135
})
const hover = (row: any) => {
  hoverLine.value = row.fileID
  hoverType.value = row.type
}
const unhover = () => {
  hoverLine.value = undefined
  hoverType.value = undefined
}
// 排序函数
const sortByName = (a: any, b: any) => {
  if (a.fileName === b.fileName) {
    return 0
  }
  return a.fileName > b.fileName ? 1 : -1
}
const sortBySize = (a: any, b: any) => {
  if (a.fileSize === b.fileSize) {
    return 0
  }
  return a.fileSize > b.fileSize ? 1 : -1
}
// const sortByCount = (a: any, b: any) => {
//   if (a.downloadNum === b.downloadNum) {
//     return 0
//   }
//   return a.downloadNum > b.downloadNum ? 1 : -1
// }
const sortByDate = (a: any, b: any) => {
  const dateA = new Date(a.updateTime)
  const dateB = new Date(b.updateTime)
  if (dateA === dateB) {
    return 0
  }
  return dateA > dateB ? 1 : -1
}
// 双击文件夹
const doubleClickEnter = (row: any) => {
  if (data) {
    if (row.type === "dir") {
      let nextPath = route.fullPath
      if (data[0].filePath !== "/") {
        nextPath += "/"
      }
      nextPath += row.fileName
      router.push(nextPath)
    }
  }
}
// 选中文件
const selectChange = (selection: DIVAOData[]) => {
  setFiles(selection)
}
</script>

<template>
  <div class="content">
    <el-table
      :data="data"
      v-loading="tableLoading"
      element-loading-text="Loading..."
      style="width: 98%"
      stripe
      :max-height="tableHeight"
      table-layout="auto"
      @cell-mouse-enter="hover"
      @cell-mouse-leave="unhover"
      @row-dblclick="doubleClickEnter"
      @selection-change="selectChange"
    >
      <el-table-column v-if="isdorf === 'D'" type="selection" width="30" />
      <el-table-column
        label="名称"
        show-overflow-tooltip
        sortable
        :sort-method="sortByName"
        width="200px"
      >
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <font-awesome-icon
              :icon="scope.row.icon.icon"
              :size="size"
              :style="{ color: scope.row.icon.color }"
            />
            <span style="margin-left: 10px">{{ scope.row.fileName }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="大小" sortable :sort-method="sortBySize">
        <template #default="scope">
          <el-tag type="primary" round effect="light" v-if="scope.row.fileSizeStr">{{
            scope.row.fileSizeStr
          }}</el-tag>
        </template>
      </el-table-column>
      <!-- <el-table-column label="下载次数" width="150px" sortable :sort-method="sortByCount">
        <template #default="scope">
          <el-tag type="info" effect="light" v-if="scope.row.fileSizeStr">{{
            scope.row.downloadNum
          }}</el-tag>
        </template>
      </el-table-column> -->
      <el-table-column label="修改日期" sortable :sort-method="sortByDate">
        <template #default="scope">
          <el-tag type="info" round effect="dark">{{
            new Date(scope.row.updateTime).toLocaleString()
          }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column v-if="isdorf === 'D'" label="路径" show-overflow-tooltip>
        <template #default="scope">
          <!-- {{ scope.row.filePath }} -->
          <el-tag type="warning" round>{{ scope.row.filePath }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <div
            v-show="
              (isdorf === 'D' && scope.row.fileID === hoverLine) ||
              (isdorf === 'F' && scope.row.fileID === hoverLine && scope.row.type === hoverType)
            "
          >
            <el-button-group>
              <el-button
                type="danger"
                :icon="Delete"
                size="small"
                @click="
                  emit(
                    'deleteFile',
                    scope.row.filePath,
                    scope.row.fileName,
                    scope.row.fileID,
                    scope.row.type
                  )
                "
              ></el-button>
              <el-button
                v-if="scope.row.fileSizeStr"
                type="primary"
                :icon="Download"
                size="small"
                @click="emit('download', scope.row.filePath, scope.row.fileName)"
              ></el-button>
            </el-button-group>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<style scoped lang="scss">
* {
  font-family: "PingFang SC";
}
.icon {
  cursor: pointer;
  margin-right: 10px;
  &:focus {
    outline: none;
  }
}
</style>

<script setup lang="ts">
import { h, nextTick, onBeforeMount, ref } from "vue"
import Title from "@/components/Title/Title.vue"
import type { DIVAOData } from "@/types/types"
import { usePageData } from "@/hooks/pages"
import Table from "@/components/Table/Table.vue"
import { getFileSizeStr, getFileTypeIcon } from "@/utils/util/util"
import { ElMessageBox, ElNotification } from "element-plus"

let data = ref<DIVAOData[]>()
const title = ref("我的视频")
const { deleteF, downloadFile } = usePageData()

const getData = async () => {
  const { videos } = usePageData()
  try {
    const res = await videos()
    if (res) {
      data.value = res.data
      data.value.forEach((item) => {
        item.fileSizeStr = getFileSizeStr(item.fileSize)
        item.icon = getFileTypeIcon(3)
      })
    }
  } catch (error) {}
}
onBeforeMount(getData)
const deleteFile = (filePath: string, fileName: string, fileID: number) => {
  // 弹窗警告
  ElMessageBox({
    message: h("p", null, [
      h("span", null, "确认删除 "),
      h("div", { style: "font-size: 20px; color: #5170b8" }, fileName + " ?"),
    ]),
    showCancelButton: true,
    confirmButtonText: "删除",
    cancelButtonText: "取消",
    beforeClose: async (action, instance, done) => {
      if (action === "confirm") {
        // 如果确定删除，就开始加载
        instance.confirmButtonLoading = true
        instance.confirmButtonText = "删除中..."
        const res = await deleteF(filePath, fileName, fileID)
        if (res) {
          ElNotification({
            type: "success",
            message: "删除成功",
          })
        }
        // 不管删除失败还是成功都要关闭弹窗
        instance.confirmButtonLoading = false
        done()
      } else {
        done()
      }
    },
  })
    .then(() => {
      nextTick(getData)
    })
    .catch(() => {})
}
const patchDeleteFiles = (files: DIVAOData[]) => {
  // 弹窗警告
  ElMessageBox({
    message: h("p", null, [h("span", null, `确认删除 ${files.length} 个文件吗?`)]),
    showCancelButton: true,
    confirmButtonText: "删除",
    cancelButtonText: "取消",
    beforeClose: async (action, instance, done) => {
      if (action === "confirm") {
        // 如果确定删除，就开始加载
        instance.confirmButtonLoading = true
        instance.confirmButtonText = "删除中..."
        let flg = true
        for (let i = 0; i < files.length; i++) {
          const res = await deleteF(files[i].filePath, files[i].fileName, files[i].fileID)
          if (!res) {
            flg = false
            ElNotification({
              type: "error",
              message: `${files[i].fileName} 删除失败`,
            })
          }
        }
        if (flg) {
          ElNotification({
            type: "success",
            message: "删除成功",
          })
        }
        // 不管删除失败还是成功都要关闭弹窗
        instance.confirmButtonLoading = false
        done()
      } else {
        done()
      }
    },
  })
    .then(() => {
      nextTick(getData)
    })
    .catch(() => {})
}
// 触发下载
const download = async (filePath: string, fileName: string) => {
  const res = await downloadFile(filePath, fileName)
  if (res) {
    const blob = new Blob([res])
    const url = URL.createObjectURL(blob)
    const link = document.createElement("a")
    link.href = url
    link.download = fileName
    link.click()
    nextTick(getData)
  }
}
</script>

<template>
  <el-container style="margin-left: 20px">
    <el-header>
      <Title :title="title" :search="false" @patch-delete-files="patchDeleteFiles" />
    </el-header>
    <el-divider style="width: 99%"></el-divider>
    <el-main>
      <Table :data="data ? data : []" @delete-file="deleteFile" @download="download" />
    </el-main>
  </el-container>
</template>

<style scoped lang="scss">
.el-main {
  --el-main-padding: 0;
}
.el-header {
  --el-header-padding: 0;
}
</style>

<script setup lang="ts">
import { h, nextTick, onBeforeMount, ref } from "vue"
import Title from "@/components/Title/Title.vue"
import type { DIVAOData } from "@/types/types"
import { usePageData } from "@/hooks/pages"
import Table from "@/components/Table/Table.vue"
import { getFileSizeStr } from "@/utils/util/util"
import { ElMessageBox, ElNotification } from "element-plus"

let data = ref<DIVAOData[]>()
const title = ref("我的音频")
const { deleteF } = usePageData()

const getData = async () => {
  const { audios } = usePageData()
  try {
    const res = await audios()
    if (res) {
      data.value = res.data
      data.value.forEach((item) => {
        item.fileSizeStr = getFileSizeStr(item.fileSize)
        item.fileIcon = ["fas", "music"]
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
</script>

<template>
  <el-container style="margin-left: 20px">
    <el-header>
      <Title :title="title" :search="false" />
    </el-header>
    <el-divider style="width: 99%"></el-divider>
    <el-main>
      <Table :data="data ? data : []" @delete-file="deleteFile" />
    </el-main>
  </el-container>
</template>

<style scoped lang="scss">
.el-main {
  --el-main-padding: 0;
}
</style>

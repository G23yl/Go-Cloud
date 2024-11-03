<script setup lang="ts">
import Title from "@/components/Title/Title.vue"
import { usePageData } from "@/hooks/pages"
import { type FFData, type UploadFile } from "@/types/types"
import { getFileSizeStr, getFileTypeIcon } from "@/utils/util/util"
import { h, nextTick, onBeforeMount, ref, useTemplateRef } from "vue"
import Table from "@/components/Table/Table.vue"
import { ElMessage, ElMessageBox, ElNotification } from "element-plus"
import request from "@/utils/axios"
import { useRouter } from "vue-router"

const title = ref("我的文件")
// 上传页面ref
const pannelRef = useTemplateRef("pannelRef")
// 上传文件input ref
const inputRef = useTemplateRef("inputRef")
// 上传文件夹input ref
const dirInputRef = useTemplateRef("dirInputRef")
// 获取到uri中的path参数
const { query = "/" } = defineProps<{
  query: string
}>()
// 发请求获取当前页面的文件
let data = ref<FFData[]>()
const uploadFileList = ref<UploadFile[]>([])
const { deleteF, createFolder, deleteFolder, downloadFile } = usePageData()
const router = useRouter()

const getData = async () => {
  const { files } = usePageData()
  try {
    const res = await files(query)
    if (res) {
      data.value = res.data
      data.value.forEach((item) => {
        if (item.type === "file") {
          item.fileSizeStr = getFileSizeStr(item.fileSize)
        }
        item.icon = getFileTypeIcon(item.fileType)
      })
    } else {
      // 如果res为undefined说明没有这个文件夹，就跳转到404页面。因为如果有这个文件夹，但是没有数据，也会返回[]
      router.push("/404")
    }
  } catch (error) {}
}
onBeforeMount(getData)
const toggle = () => {
  if (pannelRef.value) {
    const width = pannelRef.value.style.width
    const height = pannelRef.value.style.height
    const padding = pannelRef.value.style.padding
    pannelRef.value.style.width = width === "0px" || width === "" ? "750px" : "0"
    pannelRef.value.style.height = height === "0px" || width === "" ? "500px" : "0"
    pannelRef.value.style.padding = padding === "0px" || padding === "" ? "10px" : "0"
  }
}
const toggleCreate = () => {
  ElMessageBox.prompt("文件夹名称", "创建文件夹", {
    confirmButtonText: "创建",
    cancelButtonText: "取消",
  })
    .then(async ({ value }) => {
      // 判断是否存在相同名称文件夹
      if (data.value) {
        for (let i = 0; i < data.value.length; i++) {
          if (data.value[i].type === "dir" && data.value[i].fileName === value) {
            ElNotification({
              type: "warning",
              message: "文件夹名称已存在",
            })
            return
          }
        }
      }
      // 没有data的话，必定不存在相同名称文件夹，所以直接上传
      const res = await createFolder(query, value)
      if (res) {
        ElNotification({
          type: "success",
          message: "创建文件夹成功",
        })
        nextTick(getData)
      }
    })
    .catch(() => {})
}
const handleSelect = () => {
  inputRef.value?.click()
}
const handleDirSelect = () => {
  dirInputRef.value?.click()
}
const handleChange = () => {
  uploadFileList.value = []
  if (inputRef.value && inputRef.value.files && data.value) {
    const fileList = inputRef.value.files
    if (fileList.length > 0) {
      for (let i = 0; i < fileList.length; i++) {
        const uploadFile: UploadFile = {
          file: fileList[i],
          progress: 0,
          successIconShow: false,
          existSameFile: false,
        }
        for (let j = 0; j < data.value.length; j++) {
          if (uploadFile.file.name === data.value[j].fileName) {
            uploadFile.existSameFile = true
            break
          }
        }
        uploadFileList.value.push(uploadFile)
      }
      handleUpload("file")
    }
  }
}
const handleDirChange = () => {
  uploadFileList.value = []
  if (dirInputRef.value && dirInputRef.value.files && data.value) {
    const fileList = dirInputRef.value.files
    if (fileList.length > 0) {
      const folderName = fileList[0].webkitRelativePath.split("/")[0]
      // 由于上传的是文件夹，因此我只需要判断当前目录下有没有同名文件夹即可
      for (let j = 0; j < data.value.length; j++) {
        if (data.value[j].type === "dir" && data.value[j].fileName === folderName) {
          ElNotification({
            type: "error",
            message: "文件夹名称已存在",
          })
          return
        }
      }
      for (let i = 0; i < fileList.length; i++) {
        const uploadFile: UploadFile = {
          file: fileList[i],
          progress: 0,
          successIconShow: false,
          existSameFile: false,
        }
        uploadFileList.value.push(uploadFile)
      }
      handleUpload("dir")
    }
  }
}
const handleRemove = (filename: string) => {
  uploadFileList.value = uploadFileList.value.filter((item) => item.file.name !== filename)
}
const handleUpload = async (uploadType: string) => {
  // 上传文件
  if (uploadType === "file") {
    let successFlg = true
    uploadFileList.value.forEach(async (item) => {
      // 不存在同名文件才发请求
      if (!item.existSameFile) {
        let formData = new FormData()
        formData.append("file", item.file)
        formData.append("filepath", query)
        const res = await request.post("/cloud/files", formData, {
          onUploadProgress: (p) => {
            if (p.total) {
              // 进度条显示出来
              item.progress = Math.round((p.loaded / p.total) * 100)
            }
            if (p.loaded === p.total) {
              item.successIconShow = true
            }
          },
          timeout: 120000,
        })
        if (!res) {
          successFlg = false
          ElNotification({
            type: "error",
            message: `上传文件失败: ${item.file.name}`,
          })
        }
      }
    })
    if (successFlg) {
      ElNotification({
        type: "success",
        message: "上传成功",
      })
    }
    nextTick(getData)
  } else {
    let successFlg = true
    // 上传文件夹
    // 找出所有不重复的文件夹
    const AllFoldersSet = new Set<string>()
    const folderPathAndName = []
    uploadFileList.value.forEach((item) => {
      let pathItems = item.file.webkitRelativePath.split("/")
      pathItems.splice(pathItems.length - 1, 1)
      for (let i = 0; i < pathItems.length; i++) {
        let p = pathItems.slice(0, i + 1)
        AllFoldersSet.add(p.join("/"))
      }
    })
    // 要先创建父文件夹，不然会报错
    const allFoldersArr = Array.from(AllFoldersSet).sort((a: string, b: string) => {
      return a.length - b.length
    })
    // 把所有需要创建的文件夹的名字和路径整理出来
    for (let i = 0; i < allFoldersArr.length; i++) {
      const items = allFoldersArr[i].split("/")
      const folderName = items[items.length - 1]
      let partialFolderPath = items.slice(0, items.length - 1).join("/")
      let folderPath = query
      if (partialFolderPath != "") {
        if (query != "/") {
          folderPath += "/"
        }
        folderPath += partialFolderPath
      }
      folderPathAndName.push({
        folderName: folderName,
        folderPath: folderPath,
      })
    }
    // 开始创建文件夹
    for (let i = 0; i < folderPathAndName.length; i++) {
      const res = await createFolder(
        folderPathAndName[i].folderPath,
        folderPathAndName[i].folderName
      )
      if (!res) {
        successFlg = false
        ElNotification({
          type: "error",
          message: `创建文件夹失败: ${folderPathAndName[i].folderName}`,
        })
        nextTick(getData)
        return
      }
    }
    // 创建文件
    uploadFileList.value.forEach(async (fileInfo) => {
      const items = fileInfo.file.webkitRelativePath.split("/")
      const partialFilePath = items.slice(0, items.length - 1).join("/")
      let filePath = query
      if (query !== "/") {
        filePath += "/"
      }
      filePath += partialFilePath
      let formData = new FormData()
      formData.append("file", fileInfo.file)
      formData.append("filepath", filePath)
      const res = await request.post("/cloud/files", formData, {
        onUploadProgress: (p) => {
          if (p.total) {
            // 进度条显示出来
            fileInfo.progress = Math.round((p.loaded / p.total) * 100)
          }
          if (p.loaded === p.total) {
            fileInfo.successIconShow = true
          }
        },
        timeout: 120000,
      })
      if (!res) {
        successFlg = false
        ElNotification({
          type: "error",
          message: `上传文件失败: ${fileInfo.file.name}`,
        })
      }
    })
    if (successFlg) {
      ElNotification({
        type: "success",
        message: "上传成功",
      })
    }
    nextTick(getData)
  }
}
const deleteFile = (filePath: string, fileName: string, fileID: number, type: string) => {
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
        if (type === "dir") {
          // 删除文件夹
          const res = await deleteFolder(filePath, fileName, fileID)
          if (res) {
            ElNotification({
              type: "success",
              message: "删除成功",
            })
          }
        } else {
          // 删除文件
          const res = await deleteF(filePath, fileName, fileID)
          if (res) {
            ElNotification({
              type: "success",
              message: "删除成功",
            })
          }
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
      <Title :title="title" :search="false" :bread="true" :path="query" />
    </el-header>
    <el-divider style="width: 99%"></el-divider>
    <el-main>
      <Table :data="data" isdorf="F" @delete-file="deleteFile" @download="download" />
      <div class="mul-btns">
        <div class="btn" @click="toggle">
          <font-awesome-icon :icon="['fas', 'plus']" style="color: #ffffff" size="lg" />
        </div>
        <el-tooltip content="创建文件夹" placement="left">
          <div class="folder" @click="toggleCreate">
            <font-awesome-icon :icon="['fas', 'folder-plus']" style="color: #ffffff" size="lg" />
          </div>
        </el-tooltip>
      </div>
    </el-main>
    <div ref="pannelRef" class="pannel">
      <el-scrollbar>
        <div class="pannel-content">
          <el-button type="primary" @click="handleSelect">上传文件</el-button>
          <el-button type="warning" @click="handleDirSelect">上传文件夹</el-button>
          <input ref="inputRef" type="file" style="display: none" multiple @change="handleChange" />
          <input
            ref="dirInputRef"
            type="file"
            style="display: none"
            webkitdirectory
            @change="handleDirChange"
          />
          <ul class="list">
            <li class="l" v-for="file in uploadFileList" :key="file.file.name">
              <div class="upper">
                <div class="upper-left">
                  <font-awesome-icon :icon="['fas', 'file-lines']" style="margin-right: 4px" />
                  <span>{{ file.file.name }}</span>
                  <div v-show="file.existSameFile">
                    <font-awesome-icon
                      :icon="['fas', 'circle-question']"
                      style="color: #ff8040; margin: 0 0 0 15px"
                    />
                    <span style="color: #ff8040"> 已存在同名文件 </span>
                  </div>
                </div>
                <div>
                  <font-awesome-icon
                    v-show="file.successIconShow"
                    :icon="['fas', 'circle-check']"
                    style="color: #63e6be"
                  />
                  <font-awesome-icon
                    class="x-icon"
                    :icon="['fas', 'xmark']"
                    @click="handleRemove(file.file.name)"
                  />
                </div>
              </div>
              <el-progress
                :percentage="file.progress"
                color="#5cb87a"
                style="width: 100%"
              ></el-progress>
            </li>
          </ul>
        </div>
      </el-scrollbar>
    </div>
  </el-container>
</template>

<style scoped lang="scss">
@use "../../style/_variables.scss" as *;
.el-main {
  --el-main-padding: 0;
}
.el-header {
  --el-header-padding: 0;
}
* {
  box-sizing: border-box;
}
.mul-btns {
  position: fixed;
  bottom: 20px;
  right: 20px;
  z-index: 100;
  display: flex;
  flex-direction: column-reverse;
  align-items: center;
}
.btn {
  width: 56px;
  height: 56px;
  border-radius: 50%;
  background-color: $aside-account-bg-color;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: all 0.3s ease;
  cursor: pointer;
  &:hover {
    transform: scale(1.05) rotate(180deg);
  }
}
.folder {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: $aside-account-bg-color;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  margin-bottom: 15px;
  transition: all 0.8s ease-out;
  &:hover {
    transform: scale(1.02);
    box-shadow: 0px 3px 5px -1px rgba(0, 0, 0, 0.2), 0px 6px 10px 0px rgba(0, 0, 0, 0.14),
      0px 1px 18px 0px rgba(0, 0, 0, 0.12);
  }
}
.pannel {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 0;
  height: 0;
  border-radius: 20px;
  box-shadow: 0 0 20px gray;
  z-index: 99;
  transition: all 0.5s cubic-bezier(0.22, 1.44, 0.82, 1.13);
  backdrop-filter: blur(10px);
  padding: 0;
}
.list {
  list-style-type: none;
  width: 100%;
  margin-top: 10px;
  color: #333;
}
.l {
  height: 40px;
  margin: 5px 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-between;
  padding: 5px 5px 0 5px;
  overflow: hidden;
  transition: all 0.2s ease-in-out;
  border-radius: 7px;
  border: 1px solid transparent;
  background-color: #fff;
  box-shadow: 0 0 10px transparent;
  &:hover {
    border: 1px solid #ddd;
    box-shadow: 0 0 10px #ddd;
    .x-icon {
      opacity: 100;
    }
  }
}
.upper {
  width: 100%;
  display: flex;
  justify-content: space-between;
}
.upper-left {
  display: flex;
}
.x-icon {
  transition: all 0.1s ease;
  opacity: 0;
  cursor: pointer;
  margin-right: 10px;
}
</style>

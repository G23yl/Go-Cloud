<script setup lang="ts">
import Title from "@/components/Title/Title.vue"
import { usePageData } from "@/hooks/pages"
import { type FFData, type UploadFile } from "@/types/types"
import { getFileSizeStr, getFileTypeIcon } from "@/utils/util/util"
import { nextTick, onBeforeMount, ref, useTemplateRef } from "vue"
import Table from "@/components/Table/Table.vue"
import { ElNotification } from "element-plus"
import request from "@/utils/axios"

const title = ref("我的文件")
const pannelRef = useTemplateRef("pannelRef")
const inputRef = useTemplateRef("inputRef")
// 获取到uri中的path参数
const { query = "/" } = defineProps<{
  query: string
}>()
// 发请求获取当前页面的文件
let data = ref<FFData[]>([])
const uploadFileList = ref<UploadFile[]>([])

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
        item.fileIcon = getFileTypeIcon(item.fileType)
      })
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
    pannelRef.value.style.padding = padding === "0px" || padding === "" ? "20px" : "0"
  }
}
const handleSelect = () => {
  inputRef.value?.click()
}
const handleChange = () => {
  uploadFileList.value = []
  if (inputRef.value && inputRef.value.files) {
    const fileList = inputRef.value.files
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
  }
}
const handleRemove = (filename: string) => {
  uploadFileList.value = uploadFileList.value.filter((item) => item.file.name !== filename)
}
const handleUpload = () => {
  uploadFileList.value.forEach(async (item) => {
    // 不存在同名文件才发请求
    if (!item.existSameFile) {
      let formData = new FormData()
      formData.append("file", item.file)
      formData.append("filepath", query)
      try {
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
        })
        if (res) {
          ElNotification({
            type: "success",
            message: "上传成功",
          })
        }
        // 上传成功后获取新数据
        nextTick(getData)
      } catch (error) {}
    }
  })
}
</script>

<template>
  <el-container style="margin-left: 20px">
    <el-header>
      <Title :title="title" :search="false" :bread="true" :path="query" />
    </el-header>
    <el-divider style="width: 99%"></el-divider>
    <el-main>
      <Table :data="data ? data : []" />
      <div class="btn" @click="toggle">
        <font-awesome-icon class="plus" :icon="['fas', 'plus']" style="color: #ffffff" size="lg" />
      </div>
    </el-main>
    <div ref="pannelRef" class="pannel">
      <el-scrollbar>
        <div class="pannel-content">
          <el-button type="primary" @click="handleSelect">选择文件</el-button>
          <el-button type="primary" @click="handleUpload">上传文件</el-button>
          <input ref="inputRef" type="file" style="display: none" multiple @change="handleChange" />
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
* {
  box-sizing: border-box;
}
.btn {
  width: 50px;
  height: 50px;
  position: fixed;
  bottom: 20px;
  right: 20px;
  border-radius: 50%;
  background-color: $aside-account-bg-color;
  z-index: 100;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: all 0.3s ease;
  cursor: pointer;
  &:hover {
    transform: scale(1.05) rotate(180deg);
  }
}
.pannel {
  box-sizing: border-box;
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

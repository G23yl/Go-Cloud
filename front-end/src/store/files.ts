import type { DIVAOData } from "@/types/types"
import { acceptHMRUpdate, defineStore } from "pinia"
import { ref } from "vue"

const useFileStore = defineStore("file", () => {
  // 多选选中的文件
  const selectedFiles = ref<DIVAOData[]>([])

  const setFiles = (fileList: DIVAOData[]) => {
    selectedFiles.value = fileList
  }

  const getFiles = () => {
    return selectedFiles
  }

  return {
    setFiles,
    getFiles,
  }
})

// 热重载
if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useFileStore, import.meta.hot))
}

export default useFileStore

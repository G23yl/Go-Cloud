<script setup lang="ts">
import Title from "@/components/Title/Title.vue"
import { useDataUI } from "@/hooks/data_UI"
import { usePageData } from "@/hooks/pages"
import type { OverallData } from "@/types/types"
import { onBeforeMount, reactive, ref } from "vue"
import { VueUiDonut, VueUiWheel } from "vue-data-ui"

onBeforeMount(async () => {
  const { index } = usePageData()
  try {
    const res = await index()
    if (res) {
      data = res.data
    }
  } catch (error) {}
})
let data: OverallData | undefined
const title = ref("主页")
const { donutConfig, wheelConfig } = useDataUI()
const donutDataset = reactive([
  { name: "文档", values: [data ? data.categoryNum.DocNum : 0], color: "#ff0000" },
  { name: "图像", values: [data ? data.categoryNum.ImgNum : 0], color: "#ff8000" },
  { name: "视频", values: [data ? data.categoryNum.VideoNum : 0], color: "#ffff00" },
  { name: "音频", values: [data ? data.categoryNum.MusicNum : 0], color: "#42d392" },
  { name: "其他", values: [data ? data.categoryNum.OtherNum : 0], color: "#8000ff" },
])
const wheelDataset = reactive({
  percentage: data ? (data.storeDetail.size / data.storeDetail.capacity) * 100 : 0,
})
</script>

<template>
  <el-container style="margin-left: 20px">
    <el-header>
      <Title :title="title" />
    </el-header>
    <el-divider></el-divider>
    <el-main>
      <div class="content">
        <div class="group1">
          <div class="upper folder-num">
            <div class="icon">
              <font-awesome-icon
                :icon="['fas', 'folder']"
                size="2xl"
                bounce
                style="color: #5f7578"
              />
            </div>
            <el-divider direction="vertical"></el-divider>
            <div class="num">
              <span>文件夹</span>
              <span>{{ data?.foldersNum }}</span>
            </div>
          </div>
          <div class="down category">
            <VueUiDonut :config="donutConfig" :dataset="donutDataset" style="padding: 0" />
          </div>
        </div>
        <div class="group2">
          <div class="upper file-num">
            <div class="icon">
              <font-awesome-icon :icon="['fas', 'file']" size="2xl" bounce style="color: #5f7578" />
            </div>
            <el-divider direction="vertical"></el-divider>
            <div class="num">
              <span>文件</span>
              <span>{{ data?.filesNum }}</span>
            </div>
          </div>
          <div class="down store">
            <VueUiWheel
              :config="wheelConfig"
              :dataset="wheelDataset"
              style="padding: 0; width: 63%"
            />
          </div>
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<style scoped lang="scss">
@use "../../style/_variables.scss" as *;

.content {
  display: flex;
  gap: 20px;
}

.upper {
  height: 95px;
  width: 100%;
  background-color: $aside-bg-color;
  border-radius: 10px;
  display: flex;
  justify-content: space-evenly;
  align-items: center;
  transition: all 0.5s ease;
  &:hover {
    background-color: #96e6a1;
    transform: scale(1.02);
    box-shadow: 0 5px 35px 0px rgba(0, 0, 0, 0.1);
  }
}

.down {
  height: 310px;
  width: 470px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.group1,
.group2 {
  display: flex;
  flex-direction: column;
  width: calc(100vw - 300px);
  gap: 20px;
}
.num {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 5px;
  & span:nth-child(1) {
    font-size: 20px;
    color: gray;
  }
  & span:nth-child(2) {
    font-size: 25px;
    font-weight: bold;
  }
}
</style>

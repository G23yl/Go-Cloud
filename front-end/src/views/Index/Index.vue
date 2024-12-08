<script setup lang="ts">
import Title from "@/components/Title/Title.vue"
import { useDataUI } from "@/hooks/data_UI"
import { usePageData } from "@/hooks/pages"
import { getFileSizeStr } from "@/utils/util/util"
import { onBeforeMount, reactive, ref } from "vue"
import { VueUiDonut, VueUiWheel } from "vue-data-ui"

const donutDataset = reactive([
  { name: "文档", values: [0], color: "#ff0000" },
  { name: "图像", values: [0], color: "#ff8000" },
  { name: "视频", values: [0], color: "#ffff00" },
  { name: "音频", values: [0], color: "#42d392" },
  { name: "其他", values: [0], color: "#8000ff" },
])
const wheelDataset = reactive({
  percentage: 0,
})
const folders = ref(0)
const files = ref(0)
const storeCapacity = ref<string>()
const storeSize = ref<string>()
const percentage = ref<number>()
const color = [
  { color: "#67c23a", percentage: 50 },
  { color: "#e6a23c", percentage: 80 },
  { color: "#f56c6c", percentage: 100 },
]

onBeforeMount(async () => {
  const { index } = usePageData()
  try {
    const res = await index()
    if (res) {
      folders.value = res.data.foldersNum
      files.value = res.data.filesNum
      storeCapacity.value = getFileSizeStr(res.data.storeDetail.capacity)
      storeSize.value = getFileSizeStr(res.data.storeDetail.size)
      percentage.value = (res.data.storeDetail.size / res.data.storeDetail.capacity) * 100
      donutDataset[0].values = [res.data.categoryNum.DocNum]
      donutDataset[1].values = [res.data.categoryNum.ImgNum]
      donutDataset[2].values = [res.data.categoryNum.VideoNum]
      donutDataset[3].values = [res.data.categoryNum.MusicNum]
      donutDataset[4].values = [res.data.categoryNum.OtherNum]
      wheelDataset.percentage = (res.data.storeDetail.size / res.data.storeDetail.capacity) * 100
    }
  } catch (error) {}
})
const title = ref("主页")
const { donutConfig, wheelConfig } = useDataUI()
</script>

<template>
  <el-container style="margin-left: 20px">
    <el-header>
      <Title :title="title" />
    </el-header>
    <el-divider style="width: 99%"></el-divider>
    <el-main>
      <div class="content">
        <div class="group1">
          <div class="upper folder-num">
            <div class="icon">
              <font-awesome-icon
                :icon="['fas', 'folder']"
                size="2xl"
                bounce
                style="color: #6683cc"
              />
            </div>
            <el-divider direction="vertical"></el-divider>
            <div class="num">
              <span>文件夹</span>
              <span>{{ folders }}</span>
            </div>
          </div>
          <div ref="donutRef" class="down category">
            <VueUiDonut
              :config="donutConfig"
              :dataset="donutDataset"
              style="padding: 0; max-width: 90%"
            />
          </div>
        </div>
        <div class="group2">
          <div class="upper file-num">
            <div class="icon">
              <font-awesome-icon :icon="['fas', 'file']" size="2xl" bounce style="color: #6683cc" />
            </div>
            <el-divider direction="vertical"></el-divider>
            <div class="num">
              <span>文件</span>
              <span>{{ files }}</span>
            </div>
          </div>
          <div class="down store">
            <VueUiWheel
              :config="wheelConfig"
              :dataset="wheelDataset"
              style="padding: 0; width: 55%"
            />
            <el-progress
              :stroke-width="25"
              :percentage="percentage ? percentage : 0"
              :color="color"
              :text-inside="true"
              :duration="10"
              striped
              striped-flow
              style="width: 100%; margin-top: 30px"
            >
              <span>{{ storeSize }} / {{ storeCapacity }}</span>
            </el-progress>
          </div>
        </div>
      </div>
    </el-main>
  </el-container>
</template>

<style scoped lang="scss">
@use "../../style/_variables.scss" as *;

* {
  box-sizing: border-box;
}

.el-main {
  --el-main-padding: 0 20px;
}

.content {
  display: flex;
  gap: 20px;
  font-family: "PingFang SC";
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
  margin-bottom: 10px;
  &:hover {
    background-color: hsl(222, 42%, 52%, 20%);
    transform: scale(1.02);
  }
}

.down {
  height: 320px;
  width: 470px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}

.group1,
.group2 {
  display: flex;
  flex-direction: column;
  width: calc(100vw - 300px);
  // gap: 20px;
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
.capacity {
  margin-top: 5px;
  width: 100%;
  height: 15px;
  background-color: #888;
}
</style>

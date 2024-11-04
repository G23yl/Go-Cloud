<script setup lang="ts">
import { ref } from "vue"
import Title from "@/components/Title/Title.vue"
import { ElMessage } from "element-plus"

const title = ref("帮助")
const imagePaths = ref([
  "@/assets/1.jpg",
  "../../assets/2.jpg",
  "../../assets/3.jpg",
  "../../assets/4.jpg",
  "../../assets/5.jpg",
])

const show = () => {
  ElMessage({
    type: "info",
    message: "Go Cloud~~",
  })
}
</script>

<template>
  <el-container style="margin-left: 20px">
    <el-header>
      <Title :title="title" />
    </el-header>
    <el-divider style="width: 99%"></el-divider>
    <el-main>
      <div class="container">
        <div class="box left">
          <div style="display: flex">
            <font-awesome-icon :icon="['fas', 'circle-info']" size="2xl" />
            <h2>用户须知</h2>
          </div>
          <el-divider style="margin: 20px 0" width="100%"></el-divider>
          <div class="items">
            <div class="wrap">
              <div class="header" @click="show">关于Go Cloud</div>
              <div class="content"></div>
            </div>
            <div class="wrap">
              <div class="header" @click="show">用户权益</div>
              <div class="content"></div>
            </div>
            <div class="wrap">
              <div class="header" @click="show">登录与注册</div>
              <div class="content"></div>
            </div>
            <div class="wrap">
              <div class="header" @click="show">上传与下载</div>
              <div class="content"></div>
            </div>
            <div class="wrap">
              <div class="header" @click="show">项目开发</div>
              <div class="content"></div>
            </div>
          </div>
        </div>
        <div class="box right">
          <el-carousel height="400px">
            <el-carousel-item v-for="item in 4" :key="item">
              <h3 class="small justify-center" text="2xl">{{ item }}</h3>
            </el-carousel-item>
          </el-carousel>
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
h2 {
  font-size: 30px;
  font-weight: 1000;
  margin-left: 10px;
}
.container {
  display: flex;
  height: calc(100vh - 135px);
  gap: 20px;
  font-family: "PingFang SC";
  .left {
    width: 50%;
    background-color: $aside-bg-color;
  }
  .right {
    width: 50%;
    // background-color: gray;
  }
}
.box {
  padding: 20px;
  border-radius: 20px;
}
.items {
  display: flex;
  flex-direction: column;
  gap: 20px;
  margin-top: 15px;
}
.wrap {
  width: 100%;
  background-color: hsl(222, 42%, 52%, 20%);
  border-radius: 10px;
  .header {
    padding: 10px 15px;
    font-size: 17px;
    font-weight: 600;
    cursor: pointer;
  }
  .content {
    display: none;
  }
}
.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  line-height: 400px;
  margin: 0;
  text-align: center;
}
.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}
.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
</style>

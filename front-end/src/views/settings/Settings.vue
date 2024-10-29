<script setup lang="ts">
import { ref } from "vue"
import Title from "@/components/Title/Title.vue"
import { getUser } from "@/utils/user"
import useUserStore from "@/store/user"
import { ElMessage } from "element-plus"

const title = ref("个人中心")
const input = ref<HTMLInputElement>()
const cardRef = ref<HTMLDivElement>()
let user = getUser()
const { changeAvatar } = useUserStore()

const handleChange = async (e: Event) => {
  const file = (e.target as HTMLInputElement).files?.[0]
  if (file) {
    // 把头像上传到服务器，服务器返回用户新信息，存在localstorage中
    try {
      const res = await changeAvatar(file)
      if (res) {
        ElMessage({
          type: "success",
          message: "头像更新成功",
        })
      }
    } catch (error) {}
  }
}
// 鼠标在card中移动事件
const move = (e: MouseEvent) => {
  if (cardRef.value) {
    const mouseParentLeft = e.clientX - cardRef.value.offsetLeft
    const mouseParentTop = e.clientY - cardRef.value.offsetTop
    const cardWidth = cardRef.value.getBoundingClientRect().width
    const cardHeight = cardRef.value.getBoundingClientRect().height
    const rotateX = (30 * (mouseParentTop - cardHeight / 2)) / (cardHeight / 2)
    const rotateY = (30 * -(mouseParentLeft - cardWidth / 2)) / (cardWidth / 2)
    cardRef.value.style.transform = `rotateX(${rotateX}deg) rotateY(${rotateY}deg)`
  }
}
// 鼠标移出时恢复原状
const leave = () => {
  if (cardRef.value) {
    cardRef.value.style.transform = `rotateX(0deg) rotateY(0deg)`
  }
}
</script>

<template>
  <el-container style="margin-left: 20px">
    <el-header>
      <Title :title="title" :search="false" />
    </el-header>
    <el-divider style="width: 99%"></el-divider>
    <el-main>
      <div class="content">
        <div ref="cardRef" class="card" @mousemove="move" @mouseleave="leave">
          <el-image class="avatar" :src="user?.avatar" @click="input?.click()"></el-image>
          <input ref="input" type="file" style="display: none" @change="handleChange" />
          <div class="username">
            <div class="icon-div">
              <font-awesome-icon :icon="['fas', 'user']" style="color: #ffffff" size="lg" />
            </div>
            <input class="text" type="text" :value="user?.username" disabled />
          </div>
          <div class="email">
            <div class="icon-div">
              <font-awesome-icon :icon="['fas', 'envelope']" style="color: #ffffff" size="lg" />
            </div>
            <input class="text" type="text" :value="user?.email" disabled />
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
  justify-content: center;
  align-items: center;
  .card {
    box-sizing: border-box;
    width: 300px;
    height: calc(100vh - 135px);
    // height: 100vh;
    border-radius: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 15px;
    background-image: url("../../assets/layered-waves-haikei.svg");
    background-size: cover;
    perspective: 500px;
    .avatar {
      width: 90px;
      border-radius: 50%;
      background-color: #ccc;
      cursor: pointer;
    }
  }
}
.icon-div {
  background-color: $aside-account-bg-color;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 50px;
  height: 50px;
}
.username,
.email {
  display: flex;
  height: 50px;
  margin: 10px 0;
}
.username {
  margin-top: 20px;
}
.text {
  outline: none;
  border: none;
  background-color: $aside-bg-color;
  padding: 10px;
  font-weight: 1000;
  color: #999;
}
</style>

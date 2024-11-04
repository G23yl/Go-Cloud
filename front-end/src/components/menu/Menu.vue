<script setup lang="ts">
import { deleteUser, getUser } from "@/utils/user"
import { ElMessage } from "element-plus"
import { onMounted, ref, watch } from "vue"
import { useRoute, useRouter } from "vue-router"

// 获取router
const router = useRouter()
// 获取route
const route = useRoute()
// 获取到用户登录信息用于显示头像，用户名，邮箱
const user = getUser()
let activeIndex = ref("")

// 实现menu刷新依然选中
onMounted(() => {
  activeIndex.value = route.path
})
watch(
  () => route.path,
  () => {
    activeIndex.value = route.path
  }
)
const logout = () => {
  // 删除用户信息
  deleteUser()
  router.push("/login")
  ElMessage({
    type: "success",
    message: "登出成功",
  })
}
const GotoHub = () => {
  window.open("https://github.com/G23yl/Go-Cloud")
}
</script>

<template>
  <div class="aside">
    <div class="logo">
      <img
        src="../../assets/logo.svg"
        alt="logo"
        style="height: 70px; cursor: pointer"
        @click="GotoHub"
      />
    </div>
    <el-scrollbar>
      <div class="menu">
        <el-menu
          :default-active="activeIndex"
          class="el-menu-vertical"
          style="border-right: none"
          router
        >
          <el-menu-item-group title="PREVIEW">
            <el-menu-item index="/dashboard/index">
              <font-awesome-icon class="icon" :icon="['fas', 'chart-simple']" size="lg" />
              主页
            </el-menu-item>
            <el-menu-item index="/dashboard/files">
              <font-awesome-icon class="icon" :icon="['fas', 'folder']" size="lg" />
              我的文件
            </el-menu-item>
            <el-sub-menu index="3">
              <template #title>
                <font-awesome-icon class="icon" :icon="['fas', 'layer-group']" size="lg" />
                分类文件
              </template>
              <el-menu-item index="/dashboard/docs">
                <font-awesome-icon class="icon" :icon="['fas', 'file-word']" size="lg" />
                我的文档
              </el-menu-item>
              <el-menu-item index="/dashboard/images">
                <font-awesome-icon class="icon" :icon="['fas', 'file-image']" size="lg" />
                我的图像
              </el-menu-item>
              <el-menu-item index="/dashboard/videos">
                <font-awesome-icon class="icon" :icon="['fas', 'file-video']" size="lg" />
                我的视频
              </el-menu-item>
              <el-menu-item index="/dashboard/audios">
                <font-awesome-icon class="icon" :icon="['fas', 'file-audio']" size="lg" />
                我的音频
              </el-menu-item>
              <el-menu-item index="/dashboard/others">
                <font-awesome-icon class="icon" :icon="['fas', 'file-circle-question']" size="lg" />
                其他文件
              </el-menu-item>
            </el-sub-menu>
          </el-menu-item-group>
          <el-menu-item-group title="ACCOUNT">
            <el-menu-item index="/dashboard/settings">
              <font-awesome-icon class="icon" :icon="['fas', 'gear']" size="lg" />
              个人中心
            </el-menu-item>
            <el-menu-item index="/dashboard/helps">
              <font-awesome-icon class="icon" :icon="['fas', 'circle-info']" size="lg" />
              帮助
            </el-menu-item>
          </el-menu-item-group>
        </el-menu>
      </div>
    </el-scrollbar>
    <div class="account">
      <el-avatar :src="user?.avatar"></el-avatar>
      <div class="account-info">
        <div class="account-name">{{ user?.username }}</div>
        <div class="account-email">{{ user?.email }}</div>
      </div>
      <font-awesome-icon
        class="logout-icon"
        :icon="['fas', 'right-to-bracket']"
        size="lg"
        @click="logout"
      />
    </div>
  </div>
</template>

<style scoped lang="scss">
@use "../../style/_variables.scss" as *;

.aside {
  height: calc(100vh - 40px);
  border: none;
  border-radius: $border-radius;
  background-color: $aside-bg-color;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  color: #fff;
  font-family: "PingFang SC";
}
.logo {
  display: flex;
  height: 70px;
  justify-content: center;
  align-items: center;
}
.account {
  display: flex;
  align-items: center;
  padding: 15px;
  margin-top: auto;
  background-color: $aside-account-bg-color;
  .account-info {
    margin-left: 5px;
    font-size: 12px;
    font-weight: 600;
    .account-name {
      font-size: 17px;
    }
    .account-email {
      margin-top: 5px;
      color: $aside-account-email-color;
    }
  }
}
.el-menu-vertical {
  --el-menu-bg-color: $aside-bg-color;
  --el-menu-active-color: #5170b8;
  --el-menu-hover-bg-color: hsl(222, 42%, 52%, 20%);
}

.icon {
  margin-right: 10px;
}

.logout-icon {
  margin-left: 9px;
  cursor: pointer;
}
</style>

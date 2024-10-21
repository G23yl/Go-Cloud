<script setup lang="ts">
import useUserStore from "@/store/user"
import type { LoginForm } from "@/types/types"
import { ElMessage, type FormInstance, type FormRules } from "element-plus"
import { reactive, ref, useTemplateRef } from "vue"
import { useRouter } from "vue-router"

const userInfo = reactive<LoginForm>({
  username: "",
  password: "",
})
const userStore = useUserStore()
const loading = ref(false)
const router = useRouter()
const loginFormRef = useTemplateRef<FormInstance>("loginFormRef")
const rules = reactive<FormRules<LoginForm>>({
  username: [
    {
      validator(rule, value, callback) {
        if (!value) {
          callback(new Error("Please enter your username"))
        } else if (value.length < 4 || value.length > 10) {
          callback(new Error("Username must be between 4 and 10 characters"))
        } else {
          callback()
        }
      },
      trigger: "blur",
    },
  ],
  password: [
    {
      validator(rule, value, callback) {
        if (!value) {
          callback(new Error("Please enter your password"))
        } else if (value.length < 4 || value.length > 16) {
          callback(new Error("Password must be between 4 and 16 characters"))
        } else {
          callback()
        }
      },
      trigger: "blur",
    },
  ],
})

const handleSubmit = async () => {
  loading.value = true
  if (!loginFormRef.value) {
    return
  }
  try {
    await loginFormRef.value.validate((valid) => {
      if (!valid) {
        return Promise.reject()
      }
      return Promise.resolve()
    })
    const response = await userStore.login(userInfo)
    // 登录成功提示信息以及跳转
    if (response) {
      router.push("/index")
      ElMessage({
        type: "success",
        message: "Login successful!",
      })
    }
  } catch (error) {
  } finally {
    loading.value = false
  }
}
// 设置未完成功能的提示
const experiment = () => {
  ElMessage({
    type: "info",
    message: "coming soon!!!",
  })
}
</script>

<template>
  <div class="container">
    <div class="card">
      <div class="side-panel">
        <div class="logo">
          <img src="@/assets/gocloud.svg" alt="Go Cloud Logo" class="logo-image" />
        </div>
        <h2>Manage Your Files Effortlessly</h2>
        <p>Welcome! Store and manage all your documents in one secure place.</p>
        <div class="image-placeholder"></div>
      </div>
      <div class="login-panel">
        <h2>Login</h2>
        <el-form
          ref="loginFormRef"
          :model="userInfo"
          status-icon
          label-width="auto"
          label-position="top"
          size="large"
          :rules="rules"
          style="max-width: 420px; width: 100%"
        >
          <el-form-item label="Username" prop="username">
            <el-input v-model="userInfo.username"></el-input>
          </el-form-item>
          <el-form-item label="Password" prop="password">
            <el-input v-model="userInfo.password" type="password" show-password></el-input>
          </el-form-item>
          <el-button
            type="primary"
            size="large"
            class="login-button"
            :loading="loading"
            @click="handleSubmit"
          >
            Login
          </el-button>
        </el-form>
        <div class="footer">
          <el-link
            href="#"
            type="primary"
            @click="experiment"
            :underline="false"
            style="color: #a3c2c7"
            >Forgot password?</el-link
          >
          <div class="social-login">
            <span>or login with</span>
            <img
              src="@/assets/QQ.svg"
              alt="qq"
              style="width: 25px; cursor: pointer; margin-right: 5px"
              @click="experiment"
            />
            <span>or</span>
            <img
              src="@/assets/wechat-fill.svg"
              alt="wechat"
              style="width: 25px; cursor: pointer"
              @click="experiment"
            />
          </div>
          <p>
            Don't have an account?
            <el-link href="/sign-up" type="primary" :underline="false" style="color: #a3c2c7"
              >Create Account</el-link
            >
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  width: 700px;
  height: 500px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  justify-content: center;
  align-items: center;
  font-family: "PingFang SC";
}

.card {
  display: flex;
  background-color: #f7f9fc;
  border-radius: 16px;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2);
  overflow: hidden;
  width: 100%;
  max-height: 500px;
  transition: transform 0.3s ease;
}

.card:hover {
  transform: scale(1.02);
}

.side-panel {
  background-color: #a3c2c7;
  color: #ffffff;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 40px;
  width: 35%;
  text-align: center;
}

.logo img.logo-image {
  width: 200px;
  height: auto;
  margin-bottom: 20px;
}

.side-panel h2 {
  font-size: 2em;
  font-weight: bold;
  margin-bottom: 20px;
}

.side-panel p {
  font-size: 0.9em;
  line-height: 1.5;
  max-width: 80%;
  margin-bottom: 30px;
}

.image-placeholder {
  background: center/contain no-repeat;
  width: 120px;
  height: 120px;
  margin-top: 20px;
}

.login-panel {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 65%;
  padding: 60px;
}

.login-panel h2 {
  font-size: 2em;
  font-weight: bold;
  color: #333;
  margin-bottom: 25px;
}

.login-button {
  width: 100%;
  background-color: #a3c2c7;
  color: #fff;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1.1em;
  padding: 15px 0;
  margin-top: 20px;
  transition: background-color 0.3s;
}

.login-button:hover {
  background-color: #7c9599;
}

.footer {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 25px;
  font-size: 0.9em;
  color: #555;
}

.social-login {
  display: flex;
  align-items: center;
  margin: 15px 0;
}

.social-login span {
  margin-right: 10px;
  color: #777;
}
</style>

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
          <a href="#" @click="experiment">Forgot password?</a>
          <div class="social-login">
            <span>or login with</span>
            <button class="social-button" @click="experiment">QQ</button>
            <button class="social-button" @click="experiment">WeChat</button>
          </div>
          <p>Don't have an account? <a href="/sign-up">Create Account</a></p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
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
  max-width: 900px;
  max-height: 700px;
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

form {
  display: flex;
  flex-direction: column;
  width: 100%;
  max-width: 420px;
}

label {
  display: flex;
  flex-direction: column;
  margin-bottom: 15px;
  font-size: 1em;
  color: #444;
}

input[type="text"],
input[type="password"] {
  padding: 12px;
  border-radius: 8px;
  border: 1px solid #ccc;
  margin-top: 5px;
  font-size: 1em;
  transition: border-color 0.3s ease-in-out;
  background-color: #f4f6f9;
}

input[type="text"]:focus,
input[type="password"]:focus {
  border-color: #6a6afc;
  outline: none;
}

.login-button {
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

.footer a {
  color: #a3c2c7;
  text-decoration: none;
  margin-bottom: 10px;
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

.social-button {
  background-color: #ffffff;
  border: 1px solid #ddd;
  padding: 10px 20px;
  border-radius: 5px;
  margin: 0 8px;
  cursor: pointer;
  font-size: 0.9em;
  transition: background-color 0.3s;
}

.social-button:hover {
  background-color: #f1f1f1;
}
</style>

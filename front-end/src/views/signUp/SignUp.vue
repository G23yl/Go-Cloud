<script setup lang="ts">
import useUserStore from "@/store/user"
import type { SignUpForm, SignUpFormRules, SignUpInfo } from "@/types/types"
import { ElMessage, type FormInstance, type FormItemInstance, type FormRules } from "element-plus"
import { ref, reactive } from "vue"
import { useRouter } from "vue-router"

const signUpInfo = reactive<SignUpInfo>({
  username: "",
  password: "",
  code: "",
  confirmPassword: "",
  email: "",
})
const rules = reactive<FormRules<SignUpFormRules>>({
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
  confirmPassword: [
    {
      validator(rule, value, callback) {
        if (!value) {
          callback(new Error("Please confirm your password"))
        } else if (value !== signUpInfo.password) {
          callback(new Error("Password do not match"))
        } else {
          callback()
        }
      },
      trigger: "change",
    },
  ],
  email: [
    {
      validator(rule, value: string, callback) {
        // 邮箱正则
        const pattern = /^\d{5,12}@qq+\.com$/
        if (!value) {
          callback(new Error("Please enter your email"))
        } else if (!pattern.test(value)) {
          callback(new Error("Only support qq email"))
        } else {
          callback()
        }
      },
      trigger: "blur",
    },
  ],
})
const userStore = useUserStore()
// 表单ref
const signUpFormRef = ref<FormInstance>()
// email的ref
const emailRef = ref<FormItemInstance>()
// 控制发送验证码按钮的状态
let sendCodeBtn = ref(false)
// 控制发送验证码按钮加载状态
let sendCodeLoading = ref(false)
// 控制注册按钮的状态
let signUpLoading = ref(false)
// 倒计时——1min
let countDown = ref(60)
// 计时器
let timer: number | undefined
const router = useRouter()

// 发送验证码
const sendVerificationCode = async () => {
  // 邮箱检验不合格
  if (!emailRef.value || emailRef.value.validateState != "success") {
    return
  }
  try {
    // 开始加载
    sendCodeLoading.value = true
    const response = await userStore.sendCode(signUpInfo.email)
    if (response) {
      ElMessage({
        type: "success",
        message: response.msg,
      })
      // 让按钮不能点击
      sendCodeBtn.value = true
      // 60s之后才能点击
      timer = setInterval(() => {
        if (countDown.value === 0) {
          clearInterval(timer)
          // 可以点击
          sendCodeBtn.value = false
          // 恢复计数器
          countDown.value = 10
          timer = undefined
        }
        countDown.value -= 1
      }, 1000)
    }
  } catch (error) {
  } finally {
    // 加载结束
    sendCodeLoading.value = false
  }
}

const handleSignUp = async () => {
  if (!signUpFormRef.value) {
    return
  }
  try {
    // 验证表单
    await signUpFormRef.value.validate((valid) => {
      if (!valid) {
        return Promise.reject()
      }
      return Promise.resolve()
    })
    const signUpForm: SignUpForm = {
      username: signUpInfo.username,
      password: signUpInfo.password,
      code: signUpInfo.code,
      email: signUpInfo.email,
    }
    // 开始加载
    signUpLoading.value = true
    // 发送请求
    const response = await userStore.signUp(signUpForm)
    if (response) {
      ElMessage({
        type: "success",
        message: response.msg,
      })
      router.push("/login")
    }
  } catch (error) {
  } finally {
    // 取消加载
    signUpLoading.value = false
  }
}
</script>

<template>
  <div class="container">
    <div class="card">
      <div class="side-panel">
        <div class="logo">
          <img src="../../assets/login-logo.svg" alt="Go Cloud Logo" class="logo-image" />
        </div>
        <h2>Join Go Cloud Today</h2>
        <p>Sign up to start securely storing and managing your files.</p>
      </div>
      <el-scrollbar height="500px">
        <div class="signup-panel">
          <h2>Sign Up</h2>
          <el-form
            ref="signUpFormRef"
            :model="signUpInfo"
            :rules="rules"
            status-icon
            label-width="auto"
            label-position="top"
            style="max-width: 420px; width: 100%"
            size="large"
          >
            <el-form-item label="Username" prop="username">
              <el-input v-model="signUpInfo.username" placeholder="请输入用户名" />
            </el-form-item>
            <el-form-item ref="emailRef" label="Email" prop="email">
              <el-input v-model="signUpInfo.email" placeholder="请输入邮箱" />
            </el-form-item>
            <div class="verification-section">
              <el-form-item>
                <el-input v-model="signUpInfo.code" placeholder="验证码" />
              </el-form-item>
              <el-button
                type="primary"
                @click="sendVerificationCode"
                class="send-code-button"
                :disabled="sendCodeBtn"
                :loading="sendCodeLoading"
              >
                {{ timer ? `${countDown}s` : "Send" }}
              </el-button>
            </div>
            <el-form-item label="Password" prop="password">
              <el-input
                type="password"
                show-password
                v-model="signUpInfo.password"
                placeholder="请输入密码"
              />
            </el-form-item>
            <el-form-item prop="confirmPassword">
              <el-input
                type="password"
                show-password
                v-model="signUpInfo.confirmPassword"
                placeholder="确认密码"
              />
            </el-form-item>
            <el-button
              type="primary"
              size="large"
              class="signup-button"
              @click="handleSignUp"
              :loading="signUpLoading"
            >
              Sign Up
            </el-button>
          </el-form>
          <div class="footer">
            <p>Already have an account? <router-link to="/">Login</router-link></p>
          </div>
        </div>
      </el-scrollbar>
    </div>
  </div>
</template>

<style scoped lang="scss">
.container {
  width: 700px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  justify-content: center;
  align-items: center;
  height: 500px;
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

.signup-panel {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  flex-grow: 1;
  padding: 40px;
  overflow-y: auto;
}

.signup-panel h2 {
  font-size: 2em;
  font-weight: bold;
  color: #333;
  margin-bottom: 25px;
}

.el-form-item--large {
  margin-bottom: 19px;
}

/* 验证码输入框和按钮的样式 */
.verification-section {
  display: flex;
  align-items: center;
}

.send-code-button {
  width: 97px;
  background-color: #a3c2c7;
  color: #fff;
  border: none;
  padding: 8px 12px;
  border-radius: 8px;
  font-size: 1.1em;
  cursor: pointer;
  font-size: 0.9em;
  transition: background-color 0.3s;
  margin-left: 20px;
  margin-bottom: 19px;
}

.send-code-button:hover {
  background-color: #7c9599;
}

.signup-button {
  width: 100%;
  background-color: #a3c2c7;
  color: #fff;
  border: none;
  padding: 15px 0;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1.1em;
  margin-top: 10px;
  transition: background-color 0.3s;
}

.signup-button:hover {
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
</style>

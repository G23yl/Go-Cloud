import { createApp } from "vue"
import App from "./App.vue"
import router from "./router/router"
import "./style/reset.css"
import pinia from "./store"
import ElementPlus from "element-plus"
import "element-plus/dist/index.css"
import * as ElementPlusIconsVue from "@element-plus/icons-vue"

const app = createApp(App)
// 引入路由
app.use(router)
app.use(pinia)
app.use(ElementPlus)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
app.mount("#app")

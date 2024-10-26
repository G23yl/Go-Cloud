import { createApp } from "vue"
import App from "./App.vue"
import router from "./router/router"
import "./style/reset.css"
import pinia from "./store"
import ElementPlus from "element-plus"
import "element-plus/dist/index.css"
import * as ElementPlusIconsVue from "@element-plus/icons-vue"
import { library } from "@fortawesome/fontawesome-svg-core"
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome"
import {
  faChartSimple,
  faCircleInfo,
  faCircleQuestion,
  faCloudArrowDown,
  faEye,
  faFile,
  faFileAudio,
  faFileCircleQuestion,
  faFileImage,
  faFileVideo,
  faFileWord,
  faFilm,
  faFolder,
  faGear,
  faImage,
  faLayerGroup,
  faMusic,
  faPenToSquare,
  faRightToBracket,
  faShareNodes,
  faTrashCan,
} from "@fortawesome/free-solid-svg-icons"
import "vue-data-ui/style.css"
import { VueUiRadar } from "vue-data-ui"

library.add(
  faChartSimple,
  faFolder,
  faLayerGroup,
  faFileWord,
  faFileImage,
  faFileVideo,
  faFileAudio,
  faFileCircleQuestion,
  faRightToBracket,
  faCircleInfo,
  faGear,
  faFile,
  faImage,
  faTrashCan,
  faCloudArrowDown,
  faShareNodes,
  faPenToSquare,
  faEye,
  faFilm,
  faMusic,
  faCircleQuestion
)

const app = createApp(App)
// 引入路由
app.use(router)
app.use(pinia)
app.use(ElementPlus)
app.component("font-awesome-icon", FontAwesomeIcon)
app.component("VueUiRadar", VueUiRadar)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
app.mount("#app")

import { createApp } from 'vue'
// 引入清除样式
import "@/style/reset.scss"
//引入element自定义样式
import "@/style/element.scss"

//elementPlus相关配置
import ElementPlus from "element-plus";
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import "element-plus/dist/index.css"
import "element-plus/theme-chalk/dark/css-vars.css"
import * as ElIcon from "@element-plus/icons-vue"

//引入pinia仓库
import pinia from "@/store/index"
//导入路由组件
import router from "@/router/index"

// svg icons
import 'virtual:svg-icons-register'

//引入自定义插件：注册全局组件
import GlobalComponents  from "@/components/index";

//引入阿里图标 多色图标需要引入iconfont.js
import "@/assets/icons/iconfont/iconfont.scss"
import "@/assets/icons/iconfont/iconfont"

import App from './App.vue'
const app = createApp(App)
// 注册element Icons组件
for (const [key, component] of Object.entries(ElIcon)) {
    app.component(key, component)
  }
app.use(GlobalComponents).use(pinia).use(router).use(ElementPlus,{locale:zhCn}).mount("#app")


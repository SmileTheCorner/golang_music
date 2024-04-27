//引入需要注册的全局对象
import SvgIcon from "@/components/SvgIcon/index.vue";
//引入自定义的Message消息提示组件
import Message from "@/components/Message/index"
import { App } from "vue";

const allGlobalComponents = { SvgIcon};

//对外暴露插件对象
export default {
  //务必叫install方法
  install(app: App<Element>) {
    Object.keys(allGlobalComponents).forEach((key) => {
      const allGlobalComponents = { SvgIcon};
      app.component(key,allGlobalComponents[key as keyof typeof allGlobalComponents]
      );
    });

    //全局挂载Message消息提示组件
    app.config.globalProperties.$message = Message
  },
};

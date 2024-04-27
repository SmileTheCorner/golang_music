import { AuthStore } from "@/store/module/auth";
import router from "@/router/index"
import GlobalStore from "@/store/module/global";
import { LOGIN_URL } from "@/config/index";
import {RouteRecordRaw} from "vue-router"


// 引入 views 文件夹下所有 vue 文件
const views = import.meta.glob("@/views/**/*.vue");

/**
 * 初始化动态路由
 */
export const initDynamicRouter = async () => {
  const useAuthStore = AuthStore();
  const useGlobalStore = GlobalStore();
  try {
    //获取菜单列表和按钮权限列表
    await useAuthStore.getAuthMenuList();
    await useAuthStore.getAuthButtonList();
   //添加动态路由
    useAuthStore.flatMenuListGet.forEach((item) => {
      item.children && delete item.children;
      if (item.component && typeof item.component == "string") {
        item.component = views["/src/views" + item.component + ".vue"];
      }
      if (item.meta && item.meta.isFull) {
        router.addRoute(item as unknown as RouteRecordRaw);
      } else {
        router.addRoute("layout", item as unknown as RouteRecordRaw);
      }
    });
  } catch (error) {
     // 当按钮 || 菜单请求出错时，重定向到登陆页
     useGlobalStore.setToken("");
     router.replace(LOGIN_URL);
     return Promise.reject(error);
  }
};

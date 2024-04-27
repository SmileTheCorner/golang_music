//静态路由
import {RouteRecordRaw} from "vue-router"
import {HOME_URL,LOGIN_URL} from "@/config/index"
/**
 * staticRouter(静态路由)
 */
export const staticRouter: RouteRecordRaw[] = [
	{
		path: "/",
		redirect: "/layout"
	},
	{
		path: LOGIN_URL,
		name: "login",
		component: () => import("@/views/login/index.vue"),
		meta: {
			title: "登录"
		}
	},
	{
		path: "/layout",
		name: "layout",
		component: () => import("@/views/layout/index.vue"),
		redirect:HOME_URL,
		meta: {
			title: "布局"
		},
		children:[
			{
				path: HOME_URL,
				name: "home",
				component: () => import("@/views/home/index.vue"),
				meta: {
					title: "首页",
				}
			},
			{
				path: "/editCode",
				name: "editCode",
				component: () => import("@/views/system/codeGeneration/editCodeInfo.vue"),
				meta: {
					title: "代码生成信息编辑",
					isFull: 1
				}
			},
		]
	},
	// {
	// 	path: "/:pathMatch(.*)*",
	// 	redirect:"/404",
	// 	name: "any",
	// }
];

/**
 * errorRouter(错误页面路由)
 */
export const errorRouter = [
	{
		path: "/403",
		name: "403",
		component: () => import("@/components/ErrorMessage/403.vue"),
		meta: {
			title: "403页面"
		}
	},
	{
		path: "/404",
		name: "404",
		component: () => import("@/components/ErrorMessage/404.vue"),
		meta: {
			title: "404页面"
		}
	},
	{
		path: "/500",
		name: "500",
		component: () => import("@/components/ErrorMessage/500.vue"),
		meta: {
			title: "500页面"
		}
	},
];
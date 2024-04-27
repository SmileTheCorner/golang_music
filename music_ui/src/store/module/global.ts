import { defineStore } from "pinia";
import { GlobalState, ThemeConfigProps, AssemblySizeType } from "../interface";
import { DEFAULT_PRIMARY } from "@/config/index";
import piniaPersistConfig from "@/config/piniaPersist";
import {LayoutType} from "@/store/interface/index"


// defineStore 调用后返回一个函数，调用该函数获得 Store 实体
const GlobalStore = defineStore({
	// id: 必须的，在所有 Store 中唯一
	id: "GlobalState",
	// state: 返回对象的函数
	state: (): GlobalState => ({
		//是否收起菜单
		isCollapse: false,
		// token
		token: "",
		// userInfo
		userInfo: "",
		// element组件大小
		assemblySize: "default",
		// language
		language: "",
		// themeConfig
		themeConfig: {
			// 当前页面是否全屏
			maximize: false,
			// 布局切换 ==>  纵向：vertical | 经典：classic | 横向：crosswise | 分栏：columns
			layout: "vertical",
			// 默认 primary 主题颜色
			primary: DEFAULT_PRIMARY,
			// 深色模式
			isDark: false,
			// 灰色模式
			isGrey: false,
			// 色弱模式
			isWeak: false,
			// 折叠菜单
			isCollapse: false,
			// 面包屑导航
			breadcrumb: true,
			// 面包屑导航图标
			breadcrumbIcon: true,
			// 标签页
			tabs: true,
			// 标签页图标
			tabsIcon: true,
			// 页脚
			footer: true
		}
	}),
	getters: {},
	actions: {
		//setIsCollapse
		setIsCollapse(collapse: boolean) {
			this.isCollapse = collapse
		},
		// setToken
		setToken(token: string) {
			this.token = token;
		},
		// setUserInfo
		setUserInfo(userInfo: any) {
			this.userInfo = userInfo;
		},
		// setAssemblySizeSize
		setAssemblySizeSize(assemblySize: AssemblySizeType) {
			this.assemblySize = assemblySize;
		},
		// updateLanguage
		updateLanguage(language: string) {
			this.language = language;
		},
		// setThemeConfig
		setThemeConfig(themeConfig: ThemeConfigProps) {
			this.themeConfig = themeConfig;
		},
		//设置主题颜色
		setThemeColor(color: string | null) {
			if (!color) this.themeConfig.primary = DEFAULT_PRIMARY;
			else this.themeConfig.primary = color;
		},
		//设置深色模式
		setThemeDark(dark: boolean) {
			this.themeConfig.isDark = dark;
		},
		//设置灰色模式
		setThemeGreyOrWeak(type: "isGrey" | "isWeak", value: boolean) {
			//根据type设置颜色
			this.themeConfig[type] = value;
			//设置isGrey或isWeak只能有一个为真
			(value && type === "isGrey") ? this.themeConfig.isWeak = false : this.themeConfig.isGrey = false
		},
		//设置页脚的显示掩藏
		setFooter(show: boolean) {
			this.themeConfig.footer = show;
		},
		//设置菜单栏的展开和折叠
		setMenuCollapse(collapse: boolean) {
			this.themeConfig.isCollapse = collapse;
		},
		//设置布局模式
		setLayoutModel(layout: LayoutType) {
			this.themeConfig.layout = layout;
		},
	},
	persist: piniaPersistConfig("GlobalState")
});

export default GlobalStore


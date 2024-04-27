import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "path";
import path from "path";
import { createSvgIconsPlugin } from "vite-plugin-svg-icons";
import VueDevTools from "vite-plugin-vue-devtools";
import vueJsx from "@vitejs/plugin-vue-jsx"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    VueDevTools(),
    vueJsx(),
    createSvgIconsPlugin({
      // 指定需要缓存的图标文件夹
      iconDirs: [path.resolve(process.cwd(), "src/assets/icons/svg")],
      // 指定symbolId格式
      symbolId: "icon-[dir]-[name]",
    }),
  ],
  resolve: {
    alias: {
      "@": resolve(__dirname, "./src"),
    },
  },
  server: {
    open: true,
    cors: true,
    // 跨域代理配置
    proxy: {
      "/api": {
        target: "http://localhost:8090",
        changeOrigin: true,
        //rewrite: path => path.replace(/^\/api/, "")
      },
    },
  },
  css: {
	preprocessorOptions: {
	  scss: {
		additionalData: `@import "@/style/var.scss";`
	  }
	}
  },
});

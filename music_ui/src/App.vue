<template>
  <el-config-provider :size="assemblySize" :button="buttonConfig">
    <router-view></router-view>
  </el-config-provider>
</template>

<script setup lang="ts">
import { computed, reactive, onMounted } from "vue";
import GlobalStore from "@/store/module/global";
import { setThemeColor, setThemeDark ,setGreyOrWeak} from "@/utils/globalSettings";

const globalStore = GlobalStore();
// element assemblySize
const assemblySize = computed(() => globalStore.assemblySize);

// element button config
const buttonConfig = reactive({ autoInsertSpace: false });

//应用启动时加载本地的主题颜色、字体大小等
onMounted(() => {
  let themeConfig = JSON.parse(localStorage.getItem("GlobalState") || "{}").themeConfig;
  //判断globalState是否有值
  if (themeConfig) {
    //设置主题
    globalStore.setThemeConfig(themeConfig);

    // 主题颜色
    setThemeColor(themeConfig.primary, themeConfig.isDark);
    //暗黑模式
    themeConfig.isDark ? setThemeDark(true) : setThemeDark(false);
    //设置灰色和弱色模式
    if(themeConfig.isGrey) setGreyOrWeak('grey',true)
    if(themeConfig.isWeak) setGreyOrWeak('weak',true)
  }
});
</script>

<style scoped></style>

<template>
  <div class="menu-content">
    <el-scrollbar>
      <el-menu :mode="props.mode" :unique-opened="true" :router="false" :collapse="useGlobalStore.themeConfig.isCollapse">
        <SubMenu :menu-list="menuList"></SubMenu>
      </el-menu>
    </el-scrollbar>
  </div>
</template>
<script setup lang="ts" name="MyMenu">
import SubMenu from "./SubMenu.vue"
import {computed,defineProps} from "vue"
import {AuthStore} from "@/store/module/auth"
import GlobalStore from "@/store/module/global"

const props = defineProps({
  mode:{
    type:String,
    default:"vertical"
  }
})

const useAuthStore = AuthStore()
const useGlobalStore = GlobalStore()
const menuList = computed(() => useAuthStore.showMenuListGet);
</script>

<style scoped lang="scss">
.menu-content{
  width: 100%;
  height: 100%;
}
</style>

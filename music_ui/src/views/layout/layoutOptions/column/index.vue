<template>
  <div class="common-layout">
    <el-container class="container">
      <div class="aside-prefix">
        <el-scrollbar>
          <div class="card">
            <div
              v-for="item in menuList"
              :key="item.path"
              class="split-item"
              :class="{ 'split-active': splitActive === item.path || `/${splitActive.split('/')[1]}` === item.path }"
              @click="changeSubMenu(item)"
            >
              <SvgIcon class="svg-icon" :iconClass="item.meta?.icon"></SvgIcon>
              <span class="title">{{ item.meta?.title }}</span>
            </div>
          </div>
        </el-scrollbar>
      </div>
      <el-aside
        class="card twoGrad"
        v-show="subMenuList.length"
        :class="{ 'not-aside': !subMenuList.length }"
        :style="{ width: isCollapse ? '110px' : '180px' }"
      >
        <el-scrollbar>
          <el-menu
            class="menu"
            :default-active="activeMenu"
            :router="false"
            :collapse="isCollapse"
            :collapse-transition="false"
            :unique-opened="true"
          >
            <SubMenu :menu-list="subMenuList" />
          </el-menu>
        </el-scrollbar>
      </el-aside>
      <el-container class="content">
        <el-header class="header">
          <div class="header-lf">
            <SvgIcon iconClass="bicycle" class="bicycle"></SvgIcon>
            <span>管理员</span>
            <ToolBarLeft></ToolBarLeft>
          </div>
          <div class="header-rf">
            <ToolBarRight></ToolBarRight>
          </div>
        </el-header>
        <el-main class="main">
          <Main></Main>
        </el-main>
        <el-footer class="footer" v-show="GlobalStore().themeConfig.footer">
          <Footer></Footer>
        </el-footer>
      </el-container>
    </el-container>
  </div>
</template>
<script setup name="layout" lang="ts">
import { computed, ref } from "vue";
import SubMenu from "@/views/layout/components/SubMenu.vue";
import ToolBarLeft from "@/views/layout/components/ToolBarLeft.vue";
import ToolBarRight from "@/views/layout/components/ToolBarRight.vue";
import Footer from "@/views/layout/components/Footer.vue";
import Main from "@/views/layout/components/Main.vue";
import GlobalStore from "@/store/module/global";
import { AuthStore } from "@/store/module/auth";
import { useRoute, useRouter } from "vue-router";

const globalStore = GlobalStore();
const useAuthStore = AuthStore();
const route = useRoute();
const router = useRouter();
//是否折叠
const isCollapse = computed(() => globalStore.themeConfig.isCollapse);
//激活的菜单
const activeMenu = computed(
  () => (route.meta.activeMenu ? route.meta.activeMenu : route.path) as string
);
//获取菜单数据
const menuList = computed(() => useAuthStore.showMenuListGet);

const splitActive = ref("");
const subMenuList = ref<Menu.MenuOptions[]>([]);
//点击一级菜单改变
const changeSubMenu = (item: Menu.MenuOptions) => {
  splitActive.value = item.path || "";
  if (item.children?.length) {
    return (subMenuList.value = item.children);
  }
  subMenuList.value = [];
  router.push(item.path || "");
};
</script>
<style scoped lang="scss">
@import "./index.scss";
</style>

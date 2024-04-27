<template>
  <div class="common-layout">
    <el-container class="container">
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
      <el-container>
        <el-aside class="menu">
          <MyMenu></MyMenu>
        </el-aside>
        <el-container>
          <el-main class="main">
            <Main></Main>
          </el-main>
          <el-footer class="footer" v-show="GlobalStore().themeConfig.footer">
            <Footer></Footer>
          </el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>
<script setup name="layout" lang="ts">
import MyMenu from "@/views/layout/components/MyMenu.vue"
import ToolBarLeft from "@/views/layout/components/ToolBarLeft.vue"
import ToolBarRight from "@/views/layout/components/ToolBarRight.vue"
import Footer from "@/views/layout/components/Footer.vue"
import Main from "@/views/layout/components/Main.vue"
import GlobalStore from "@/store/module/global";


</script>
<style scoped lang="scss">
@import "./index.scss";
</style>

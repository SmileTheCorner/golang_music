<template>
  <div class="expand">
    <el-icon
      style="vertical-align: middle"
      :size="25"
      @click="changeIsCollapse"
    >
    <component :is="useGlobalStore.themeConfig.isCollapse ? 'Expand': 'Fold'"></component>
    </el-icon>
    <div class="breadcrumb">
      <el-breadcrumb separator-icon="ArrowRight">
        <el-breadcrumb-item :to="{ path: '/' }">homepage</el-breadcrumb-item>
        <el-breadcrumb-item>首页</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
  </div>
</template>
<script setup lang="ts">
import GlobalStore from "@/store/module/global";
const useGlobalStore = GlobalStore();
const changeIsCollapse = () => {
  useGlobalStore.setMenuCollapse(!useGlobalStore.themeConfig.isCollapse);
};
</script>
<style scoped lang="scss">
.expand {
  margin-left: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  .breadcrumb {
    margin-left: 20px;
  }
}
</style>

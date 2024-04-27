<template>
    <template v-for="subItem in menuList" :key="subItem.path">
      <el-sub-menu v-if="subItem.children?.length" :index="subItem.path">
        <template #title>
          <SvgIcon :iconClass="subItem.meta?.icon"></SvgIcon>
          <span class="sle">{{ subItem.meta?.title }}</span>
        </template>
        <SubMenu :menu-list="subItem.children" />
      </el-sub-menu>
      <el-menu-item v-else :index="subItem.path" @click="handleClickMenu(subItem)">
        <SvgIcon :iconClass="subItem.meta?.icon"></SvgIcon>
        <template #title>
          <span class="sle">{{ subItem.meta?.title }}</span>
        </template>
      </el-menu-item>
    </template>
  </template>

<script setup lang="ts" name="SubMenu">
import { useRouter } from "vue-router";
defineProps<{ menuList: Menu.MenuOptions[] }>();

const router = useRouter();
const handleClickMenu = (subItem: Menu.MenuOptions) => {
  if (subItem.meta?.isLink) return window.open(subItem.meta?.isLink, "_blank");
  if (typeof subItem.path === 'string') {
    router.push(subItem.path);
  }
};
</script>
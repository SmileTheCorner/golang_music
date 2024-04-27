<template>
    <component :is="LayoutComponents[layout]" />
</template>
<script lang="ts" setup name="layout">
import LayoutVertical from "./layoutOptions/classic/index.vue";
import LayoutClassic from "./layoutOptions/vertical/index.vue";
import LayoutCrosswise from "./layoutOptions/crosswise/index.vue";
import LayoutColumn from "./layoutOptions/column/index.vue";
import { computed, type Component } from "vue";
import {LayoutType} from "@/store/interface/index"
import GlobalStore from "@/store/module/global";


const globalStore = GlobalStore();

const LayoutComponents: Record<LayoutType, Component> = {
  vertical: LayoutVertical,
  classic: LayoutClassic,
  crosswise: LayoutCrosswise,
  columns: LayoutColumn
};

const layout = computed(() => globalStore.themeConfig.layout);
</script>
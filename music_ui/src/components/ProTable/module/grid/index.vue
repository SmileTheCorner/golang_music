<template>
  <div :style="style">
    <slot></slot>
  </div>
</template>
<script setup lang="ts" name="Grid">
import { BreakPoint } from "@/components/ProTable/interface/index";
import {
  computed,
  ref,
  watch,
  provide,
  useSlots,
  onBeforeMount,
  onMounted,
  onActivated,
  onUnmounted,
  onDeactivated,
  VNodeArrayChildren,
  VNode,
} from "vue";
//定义接收的类型
type Props = {
  cols?: number | Record<BreakPoint, number>; //列宽的配置
  collapsed?: boolean; //搜索栏是否展开
  collapsedRows?: number; //搜索栏的行数
  gap?: [number, number] | number; //grid布局的间距
};
//接收参数
const props = withDefaults(defineProps<Props>(), {
  cols: () => ({ xs: 1, sm: 2, md: 2, lg: 3, xl: 4 }),
  collapsed: false,
  collapsedRows: 1,
  gap: 0,
});
//DOM即将挂载
onBeforeMount(() => props.collapsed && findIndex());
//DOM挂载完成
onMounted(() => {
  resize({ target: { innerWidth: window.innerWidth } } as unknown as UIEvent);
  window.addEventListener("resize", resize);
});
//被包含在<keep-alive>中的组件，会多出两个生命周期钩子函数。被激活时执行
onActivated(() => {
  resize({ target: { innerWidth: window.innerWidth } } as unknown as UIEvent);
  window.addEventListener("resize", resize);
});
//组件卸载完成后执行的函数
onUnmounted(() => {
  window.removeEventListener("resize", resize);
});
//比如从 A 组件，切换到 B 组件，A 组件消失时执行。
onDeactivated(() => {
  window.removeEventListener("resize", resize);
});
// 监听屏幕变化
const resize = (e: UIEvent) => {
  let width = (e.target as Window).innerWidth;
  switch (!!width) {
    case width < 768:
      breakPoint.value = "xs";
      break;
    case width >= 768 && width < 992:
      breakPoint.value = "sm";
      break;
    case width >= 992 && width < 1200:
      breakPoint.value = "md";
      break;
    case width >= 1200 && width < 1920:
      breakPoint.value = "lg";
      break;
    case width >= 1920:
      breakPoint.value = "xl";
      break;
  }
};

// 注入 gap 间距
provide("gap", Array.isArray(props.gap) ? props.gap[0] : props.gap);

// 注入响应式断点
let breakPoint = ref<BreakPoint>("xl");
provide("breakPoint", breakPoint);

// 注入要开始折叠的 index
const hiddenIndex = ref(-1);
provide("shouldHiddenIndex", hiddenIndex);

// 注入 cols
const gridCols = computed(() => {
  if (typeof props.cols === "object")
    return props.cols[breakPoint.value] ?? props.cols;
  return props.cols;
});
provide("cols", gridCols);

// 寻找需要开始折叠的字段 index
const slots = useSlots().default!();

const findIndex = () => {
  let fields: VNodeArrayChildren = [];
  let suffix: VNode | null = null;
  //遍历插槽数据
  slots.forEach((slot: any) => {
    // suffix
    if (
      typeof slot.type === "object" &&
      slot.type.name === "GridItem" &&
      slot.props?.suffix !== undefined
    )
      suffix = slot;
    // slot children
    if (typeof slot.type === "symbol" && Array.isArray(slot.children))
      fields.push(...slot.children);
  });
  // 计算 suffix 所占用的列
  let suffixCols = 0;
  if (suffix) {
    suffixCols =
      ((suffix as VNode).props![breakPoint.value]?.span ??
        (suffix as VNode).props?.span ??
        1) +
      ((suffix as VNode).props![breakPoint.value]?.offset ??
        (suffix as VNode).props?.offset ??
        0);
  }

  try {
    let find = false;
    fields.reduce((prev = 0, current, index) => {
      prev +=
        ((current as VNode)!.props![breakPoint.value]?.span ??
          (current as VNode)!.props?.span ??
          1) +
        ((current as VNode)!.props![breakPoint.value]?.offset ??
          (current as VNode)!.props?.offset ??
          0);
      if (Number(prev) > props.collapsedRows * gridCols.value - suffixCols) {
        hiddenIndex.value = index;
        find = true;
        throw "find it";
      }
      return prev;
    }, 0);
    if (!find) hiddenIndex.value = -1;
  } catch (e) {
    console.warn(e);
  }
};

// 断点变化时 执行 findIndex
watch(
  () => breakPoint.value,
  () => {
    if (props.collapsed) findIndex();
  }
);

// 监听 collapsed
watch(
  () => props.collapsed,
  (value) => {
    if (value) return findIndex();
    hiddenIndex.value = -1;
  }
);

//设置间距
const gridGap = computed(() => {
  if (typeof props.gap === "number") return `${props.gap}px`;
  if (Array.isArray(props.gap)) `${props.gap[1]}px ${props.gap[0]}px`;
  return "unset";
});
// 设置 style
const style = computed(() => {
  return {
    display: "grid",
    gridGap: gridGap.value,
    gridTemplateColumns: `repeat(${gridCols.value}, minmax(0, 1fr))`,
  };
});

defineExpose({ breakPoint });
</script>

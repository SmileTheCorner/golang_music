<template>
  <div class="filter">
    <div class="filter-header">
      <h4 class="title">{{ title }}</h4>
      <el-input
        v-model="filterText"
        placeholder="输入关键字进行过滤"
        clearable
      />
    </div>
    <div class="tree-content">
      <el-scrollbar height="100%">
        <el-tree
          ref="treeRef"
          class="filter-tree"
          default-expand-all
          :node-key="id"
          :show-checkbox="multiple"
          :data="multiple ? treeData : treeAllData"
          :props="defaultProps"
          :current-node-key="!multiple ? selected : ''"
          :highlight-current="!multiple"
          :check-on-click-node="multiple"
          :default-checked-keys="multiple ? selected : []"
          :filter-node-method="filterNode"
          @node-click="handleNodeClick"
          @check="handleCheckChange"
        />
      </el-scrollbar>
    </div>
  </div>
</template>
<script setup lang="ts" name="TreeFilter">
import { ref, watch, onBeforeMount,nextTick } from "vue";
import { ElTree } from "element-plus";

// 接收父组件参数并设置默认值
interface TreeFilterProps {
  requestApi?: (data?: any) => Promise<any>; // 请求分类数据的 api ==> 非必传
  data?: { [key: string]: any }[]; // 分类数据，如果有分类数据，则不会执行 api 请求 ==> 非必传
  title: string; // treeFilter 标题 ==> 非必传
  id: string | number; // 选择的id ==> 非必传，默认为 “id”
  label: string; // 显示的label ==> 非必传，默认为 “label”
  multiple: boolean; // 是否为多选 ==> 非必传，默认为 false
  defaultValue?: any; // 默认选中的值 ==> 非必传
}

//设置默认值
const props = withDefaults(defineProps<TreeFilterProps>(), {});

const defaultProps = {
  children: "children",
  label: props.label,
};

interface Tree {
  [key: string]: any;
}

const treeRef = ref<InstanceType<typeof ElTree>>();
const treeData = ref<{ [key: string]: any }[]>([]);
const treeAllData = ref<{ [key: string]: any }[]>([]);
const filterText = ref<string>("");

//多选的时候进行处理
const selected = ref();
const setSelected = () => {
  if (props.multiple) {
    selected.value = Array.isArray(props.defaultValue)
      ? props.defaultValue
      : [props.defaultValue];
  } else {
    selected.value =
      typeof props.defaultValue === "string" ? props.defaultValue : "";
  }
};

//页面挂载完成后获取数据
onBeforeMount(async () => {
  //对默认的值进行处理
  setSelected();

  //判断是否有请求的接口，如果有则发送请求获取数据
  if (props.requestApi) {
    try {
      const data  = await props.requestApi();
      //将值赋给treeData
      treeData.value = data;
      treeAllData.value = [{ id: "", [props.label]: "全部" }, ...data];
    } catch (error) {
      console.log(error);
    }
  }
});

// 使用 nextTick 防止打包后赋值不生效，开发环境是正常的
watch(
  () => props.defaultValue,
  () => nextTick(() => setSelected()),
  { deep: true, immediate: true }
);

//监听输入框的输入
watch(filterText, (val) => {
  treeRef.value!.filter(val);
});

//监听数据的变化
watch(
  () => props.data,
  () => {
    if (props.data?.length) {
      treeData.value = props.data;
      treeAllData.value = [{ id: "", [props.label]: "全部" }, ...props.data];
    }
  },
  { deep: true, immediate: true }
);

//过滤数据 value--输入的值 data--树节点
const filterNode = (value: string, data: Tree) => {
  if (!value) return true;
  return data.label.includes(value);
};

interface FilterEmits {
  (e: "change", value: any): void;
}
const emit = defineEmits<FilterEmits>();

//当点击节点的时候触发该函数
const handleNodeClick = (nodeData: { [key: string]: any }) => {
  //如果是多选则不需要进行返回id
  if (props.multiple) return;
  emit("change", nodeData[props.id]);
};

//多选时点击节点复选框之后触发
const handleCheckChange = () => {
  emit("change", treeRef.value?.getCheckedKeys());
};

// 暴露给父组件使用
defineExpose({ treeData, treeAllData, treeRef });
</script>
<style scoped lang="scss">
.filter {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
  .filter-header {
    margin-bottom: 10px;
    .title {
      margin: 0 0 15px;
      font-size: 18px;
      font-weight: bold;
      color: var(--el-color-info-dark-2);
      letter-spacing: 0.5px;
    }
  }
  .tree-content {
    flex: 1;

    :deep(.el-tree) {
      overflow: auto;
    }
  }
}
</style>

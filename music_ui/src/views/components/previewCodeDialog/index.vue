<template>
  <el-dialog
    v-model="centerDialogVisible"
    :title="dialogProps.title"
    width="80%"
    align-center
  >
    <div class="table-content">
      <el-tabs v-model="activeName" class="tabs">
        <el-tab-pane label="controllerCode" name="controllerCode">
          <pre class="code-pre">{{ dialogProps.rowData.controllerCode }}</pre>
        </el-tab-pane>
        <el-tab-pane label="entityCode" name="entityCode">
          <pre class="code-pre">{{ dialogProps.rowData.entityCode }}</pre>
        </el-tab-pane>
        <el-tab-pane label="serviceCode" name="serviceCode">
          <pre class="code-pre">{{ dialogProps.rowData.serviceCode }}</pre>
        </el-tab-pane>
        <el-tab-pane label="daoCode" name="daoCode">
          <pre class="code-pre">{{ dialogProps.rowData.daoCode }}</pre>
        </el-tab-pane>
        <el-tab-pane label="routerCode" name="routerCode">
          <pre class="code-pre">{{ dialogProps.rowData.routerCode }}</pre>
        </el-tab-pane>
        <el-tab-pane label="axiosCode" name="axiosCode">
          <pre class="code-pre">{{ dialogProps.rowData.axiosCode }}</pre>
        </el-tab-pane>
        <el-tab-pane label="indexCode" name="indexCode">
          <pre class="code-pre">{{ dialogProps.rowData.indexCode }}</pre>
        </el-tab-pane>
        <el-tab-pane label="typeCode" name="typeCode">
          <pre class="code-pre">{{ dialogProps.rowData.typeCode }}</pre>
        </el-tab-pane>
        <el-tab-pane label="sqlCode" name="sqlCode">
          <pre class="code-pre">{{ dialogProps.rowData.sqlCode }}</pre>
        </el-tab-pane>
        <div class="copy-icon" @click="copyCode">
          <el-tooltip effect="dark" content="复制" placement="top">
            <el-icon><CopyDocument /></el-icon>
          </el-tooltip>
        </div>
      </el-tabs>
    </div>
  </el-dialog>
</template>
<script setup lang="ts" name="codeGeneration">
import { ref } from "vue";
import useClipboard from "vue-clipboard3";
import { ElMessage } from "element-plus";

//定义对话框类型
interface DialogProps {
  title: string;
  rowData: any;
}

//显示参数
const dialogProps = ref<DialogProps>({
  title: "",
  rowData: {},
});
// 控制对话框的显示和隐藏
const centerDialogVisible = ref<boolean>(false);
// 接收父组件传过来的参数
const acceptParams = (params: DialogProps) => {
  dialogProps.value = params;
  centerDialogVisible.value = true;
};

const activeName = ref("controllerCode");
//复制代码
// 使用插件
const { toClipboard } = useClipboard();
const copyCode = async () => {
  try {
    let content = dialogProps.value.rowData[activeName.value];
    // 复制
    await toClipboard(content);
    ElMessage({
      message: "已复制",
      type: "success",
    });
  } catch (error) {
    console.log("error==>", error);
  }
};

//将需要暴露出去的数据与方法都可以暴露出去
defineExpose({
  acceptParams,
});
</script>
<style lang="scss" scoped>
.table-content {
  height: 70vh;
  overflow: auto;
  .tabs {
    position: relative;
    .code-pre {
      overflow-x: auto;
    }
    .copy-icon {
      position: absolute;
      right: 30px;
      top: 10px;
      font-size: 20px;
    }
  }
}
</style>

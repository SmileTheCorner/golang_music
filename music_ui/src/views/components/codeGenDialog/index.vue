<template>
  <el-dialog
    v-model="centerDialogVisible"
    :title="dialogProps.title"
    width="80%"
    align-center
  >
    <div class="table-content">
      <ProTable
        ref="proTable"
        title="数据表"
        rowKey="id"
        :operate="false"
        :indent="30"
        :columns="columns"
        :requestApi="getAllTableList"
        :requestAuto="true"
        :searchCol="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 4 }"
      >
      </ProTable>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="makeSure()"> 确定 </el-button>
        <el-button @click="cancel()">取消</el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script setup lang="ts" name="codeGeneration">
import { ref } from "vue";
import moment from "moment";
import ProTable from "@/components/ProTable/index.vue";
import { ColumnProps, ProTableInstance } from "@/components/ProTable/interface";
import { CodeGen, GenTable } from "@/type/codeGeneration/index";
import { getAllTableList, importGenTable } from "@/api/system/codeGen/index";

// 获取 ProTable DOM
const proTable = ref<ProTableInstance>();

//定义对话框类型
interface DialogProps {
  title: string;
  proRef: ProTableInstance;
}

//显示参数
const dialogProps = ref<DialogProps>({
  title: "",
  proRef:undefined,
});
// 控制对话框的显示和隐藏
const centerDialogVisible = ref<boolean>(false);
// 接收父组件传过来的参数
const acceptParams = (params: DialogProps) => {
  dialogProps.value = params;
  centerDialogVisible.value = true;
};

//数据列
const columns: ColumnProps<CodeGen>[] = [
  { type: "selection", width: "60" },
  { prop: "tableSchema", label: "数据库名" },
  { prop: "tableName", label: "表名", search: { el: "input" } },
  { prop: "tableComment", label: "表描述", search: { el: "input" } },
  {
    prop: "createTime",
    label: "创建时间",
    render: (scope) => {
      return moment(scope.row.createTime).format("YYYY-MM-DD HH:mm:ss");
    },
  },
  {
    prop: "updateTime",
    label: "修改时间",
    render: (scope) => {
      return moment(scope.row.updateTime).format("YYYY-MM-DD HH:mm:ss");
    },
  },
];

//确定按钮
const makeSure = async () => {
  let params: GenTable[] = [];
  //获取多选到的数据id
  let selects = proTable.value.selectedList as CodeGen;

  for (let value of selects) {
    let obj: GenTable = {
      tableSchema: value.tableSchema,
      tableName: value.tableName,
      tableComment: value.tableComment,
    };
    params.push(obj);
  }

  try {
    await importGenTable(params);
    centerDialogVisible.value = false;
    dialogProps.value.proRef.getTableList()
  } catch (error) {
    console.log("error=====>", error);
  }
};
//取消按钮
const cancel = () => {
  centerDialogVisible.value = false;
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
}
</style>

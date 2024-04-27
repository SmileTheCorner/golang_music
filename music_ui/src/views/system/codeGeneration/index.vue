<template>
  <div class="app-container">
    <ProTable
      ref="proTableFather"
      title="数据表"
      rowKey="id"
      :indent="30"
      :columns="columns"
      :requestApi="getGenTableList"
      :requestAuto="true"
      :searchCol="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 4 }"
    >
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button type="primary" icon="CirclePlus" @click="generateCode('')"
          >生成</el-button
        >
        <el-button icon="CirclePlus" @click="importTable()">导入</el-button>
        <el-button type="success" icon="EditPen" @click="edit('')"
          >修改</el-button
        >
        <el-button
          type="danger"
          icon="Delete"
          @click="drop(proTableFather?.selectedListIds)"
          >删除</el-button
        >
      </template>
      <!-- 表格操作按钮 -->
      <template #operation="scope">
        <el-button link icon="View" @click="preview(scope.row)">预览</el-button>
        <el-button
          type="success"
          link
          icon="EditPen"
          @click="edit(scope.row.id, scope.row.tableName)"
          >编辑</el-button
        >
        <el-button
          type="primary"
          link
          icon="Download"
          @click="download(scope.row.id)"
          >生成代码</el-button
        >
        <el-button type="danger" link icon="Delete" @click="drop(scope.row.id)"
          >删除</el-button
        >
      </template>
    </ProTable>
    <CodeGenDialog ref="dialogRef" />
    <PreviewCodeDialog ref="previewCodeDialogRef" />
  </div>
</template>
<script setup lang="ts">
import { ref } from "vue";
import moment from "moment";
import ProTable from "@/components/ProTable/index.vue";
import { ColumnProps, ProTableInstance } from "@/components/ProTable/interface";
import { GenTable, CodeInfo } from "@/type/codeGeneration/index";
import {
  getGenTableList,
  getColumnsByName,
  oneKeyGenerateCode,
  dropGenTableById,
  downloadCode,
} from "@/api/system/codeGen/index";
import CodeGenDialog from "@/views/components/codeGenDialog/index.vue";
import PreviewCodeDialog from "@/views/components/previewCodeDialog/index.vue";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";

const router = useRouter();
// 获取 ProTable 元素，调用其获取刷新数据方法（还能获取到当前查询参数，方便导出携带参数）
const proTableFather = ref<ProTableInstance>();

//数据列
const columns: ColumnProps<GenTable>[] = [
  { type: "selection", width: "60" },
  { prop: "tableSchema", label: "数据库名" },
  { prop: "tableName", label: "表名", search: { el: "input" } },
  { prop: "structName", label: "结构体名", search: { el: "input" } },
  { prop: "tableComment", label: "表描述", search: { el: "input" } },
  {
    prop: "createTime",
    label: "创建时间",
    render: (scope) => {
      return moment(scope.row.createdAt).format("YYYY-MM-DD HH:mm:ss");
    },
  },
  {
    prop: "updateTime",
    label: "修改时间",
    render: (scope) => {
      return moment(scope.row.updatedAt).format("YYYY-MM-DD HH:mm:ss");
    },
  },
  { prop: "operation", label: "操作", width: "310", fixed: "right" },
];
// 打开 dialog(新增、查看、编辑)
const dialogRef = ref<InstanceType<typeof CodeGenDialog> | null>(null);
const previewCodeDialogRef = ref<InstanceType<typeof PreviewCodeDialog> | null>(
  null
);
//生成代码
const generateCode = (id: string | number) => {
  console.log(id);
};
//下载生成的代码
const download = async (
  id: string | number | string[] | number[] | undefined
) => {
  try {
    let ids = Array.isArray(id) ? id.map(String) : [String(id)];
    let result = await downloadCode(ids);
  } catch (error) {
    console.log("error==>", error);
  }
};
//导入数据表
const importTable = () => {
  const params = {
    title: "导入数据表",
  };
  dialogRef.value?.acceptParams(params);
};
//预览生成的代码
const preview = async (row: GenTable) => {
  try {
    let tableName = row.tableName!;
    let data = await getColumnsByName(tableName);
    let param: CodeInfo = {
      moduleTitle: row.tableComment,
      author: "钱",
      moduleName: row.structName,
      fields: data,
    };
    let result = await oneKeyGenerateCode(param);
    previewCodeDialogRef.value?.acceptParams({
      title: "预览代码",
      rowData: result,
    });
  } catch (error) {
    console.log("error==>", error);
  }
};
//删除数据
const drop = async (id: string | number | string[] | number[] | undefined) => {
  let ids = Array.isArray(id) ? id.map(String) : [String(id)];
  await dropGenTableById(ids);
  ElMessage({
    message: "删除成功",
    type: "success",
  });
  //从新获取数据
  proTableFather.value?.getTableList();
};

//编辑代码生成的基本信息
const edit = (id: string | number, tableName: string) => {
  let params = {
    id,
    tableName,
  };
  router.push({
    name: "editCode",
    state: {
      params,
    },
  });
};
</script>

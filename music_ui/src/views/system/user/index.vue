<template>
  <div class="app-container">
    <ProTable
      ref="proTableFather"
      title="数据表"
      rowKey="id"
      :indent="30"
      :columns="columns"
      :requestApi="list"
      :treeFilterApi="departmentTree"
      treeFilterLabel="name"
      :requestAuto="true"
      :treeFilter="true"
      treeFilterTitle="部门列表"
      :searchCol="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 4 }"
    >
      <!-- 表格 header 按钮 -->
      <template #tableHeader>
        <el-button type="primary" icon="Download">新增</el-button>
        <el-button type="danger" icon="Delete">删除</el-button>
      </template>
      <!-- 表格操作按钮 -->
      <template #operation="scope">
        <el-button type="success" link icon="EditPen">编辑</el-button>
        <el-button type="danger" link icon="Delete">删除</el-button>
      </template>
    </ProTable>
  </div>
</template>
<script setup lang="ts">
import { ref } from "vue";
import moment from "moment";
import ProTable from "@/components/ProTable/index.vue";
import { ColumnProps, ProTableInstance } from "@/components/ProTable/interface";
import { User } from "@/type/user/index";
import { list } from "@/api/system/user/index";
import { departmentTree } from "@/api/system/department/index";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";

// 获取 ProTable 元素，调用其获取刷新数据方法（还能获取到当前查询参数，方便导出携带参数）
const proTableFather = ref<ProTableInstance>();

//数据列
const columns: ColumnProps<User>[] = [
  { type: "selection", width: "60" },
  { prop: "name", label: "昵称" },
  { prop: "realName", label: "名称", search: { el: "input" } },
  { prop: "avatar", label: "头像" },
  { prop: "mobile", label: "手机" },
  { prop: "email", label: "邮箱" },
  { prop: "gender", label: "性别" },
  { prop: "idCard", label: "身份证" },
  { prop: "address", label: "地址" },
  { prop: "status", label: "状态" },
  { prop: "age", label: "年龄" },
  {
    prop: "CreatedAt",
    label: "创建时间",
    width: "180",
    render(scope) {
      return moment(scope.row.CreatedAt).format("YYYY-MM-DD HH:mm:ss");
    },
  },
  { prop: "operation", label: "操作", width: "150", fixed: "right" },
];
</script>

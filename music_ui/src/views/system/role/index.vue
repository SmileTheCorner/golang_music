<template>
  <div class="app-container">
    <div class="table-box">
      <ProTable ref="proTable" title="角色列表" rowKey="id" :indent="30" :columns="columns" :requestApi="getRoleList"
        :requestAuto="true" :searchCol="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 4 }">
        <!-- 表格 header 按钮 -->
        <template #tableHeader>
          <el-button type="primary" icon="CirclePlus" @click="openDrawer('新增',defaultData)">新增</el-button>
        </template>
        <!-- 表格操作 -->
        <template #operation="scope">
          <el-button type="primary" link icon="View" @click="openDrawer('查看', scope.row)">查看</el-button>
          <el-button type="success" link icon="EditPen" @click="openDrawer('编辑', scope.row)">编辑</el-button>
          <el-button type="danger" link icon="Delete" @click="deleteData(scope.row)">删除</el-button>
        </template>
      </ProTable>
      <RoleDrawer ref="drawerRef" />
    </div>
  </div>
</template>
  
<script lang="tsx" setup name="department">
import { ref } from "vue";
import { ColumnProps, ProTableInstance } from "@/components/ProTable/interface";
import moment from "moment";
import ProTable from "@/components/ProTable/index.vue";
import { useHandleData } from "@/hooks/useHandleData";
import { getRoleList, addRole, updateRole, dropRole } from "@/api/system/role/index";
import RoleDrawer from "@/views/components/roleDrawer/index.vue";
import { Role } from "@/type/role/index"

// 表格配置项
const columns: ColumnProps<Role>[] = [
  { type: "index", label: "序号", width: "60" },
  { prop: "name", label: "角色名称" },
  { prop: "code", label: "角色标识" },
  { prop: "sort", label: "显示顺序" },
  { prop: "dataScope", label: "数据范围" },
  { prop: "dataScopeDept", label: "指定部门数据", showOverflowTooltip: true },
  {
    prop: "status", label: "角色状态", render: (scope) => {
      return scope.row.status == 0 ? <el-tag type='success'>正常</el-tag> : <el-tag type='danger'>停用</el-tag>
    }
  },
  { prop: "remark", label: "备注" },
  {
    prop: "createdAt", label: "创建时间", render: (scope) => {
      return moment(scope.row.createdAt).format("YYYY-MM-DD HH:mm:ss");
    }
  },
  { prop: "operation", label: "操作", width: "220", fixed: "right" },
];

// 打开 dialog(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof RoleDrawer> | null>(null);
// 获取 ProTable 元素，调用其获取刷新数据方法（还能获取到当前查询参数，方便导出携带参数）
const proTable = ref<ProTableInstance>();

//默认数据
const defaultData = ref<Role>({
  sort:1,
  status:0
})

//查看编辑
const openDrawer = (title: string, rowData: Partial<Role> = {}) => {
  const params = {
    title,
    rowData: { ...rowData },
    isView: title === "查看",
    api:
      title === "新增"
        ? addRole
        : title === "编辑"
          ? updateRole
          : undefined,
    getTableList: proTable.value?.getTableList,
  };
  drawerRef.value?.acceptParams(params);
};

//删除部门信息
const deleteData = async (row: Role) => {
  await useHandleData(dropRole,[row.id], `删除${row.name}`)
  //重新获取数据
  proTable.value?.getTableList();
};
</script>
<style lang="scss" scoped>
.table-box {
  width: 100%;
  height: 100%;
}
</style>
  
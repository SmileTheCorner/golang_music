<template>
  <div class="app-container">
    <div class="table-box">
      <ProTable ref="proTable" title="部门列表" rowKey="id" :indent="30" :columns="columns" :requestApi="getDepartmentList"
        :requestAuto="true" :searchCol="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 4 }">
        <!-- 表格 header 按钮 -->
        <template #tableHeader>
          <el-button type="primary" icon="CirclePlus" @click="openDrawer('新增', defaultParams, '')">新增</el-button>
        </template>
        <!-- 表格操作 -->
        <template #operation="scope">
          <el-button  link icon="CirclePlus"
            @click="openDrawer('新增', defaultParams, scope.row.id)">新增</el-button>
          <el-button type="primary" link icon="View" @click="openDrawer('查看', scope.row)">查看</el-button>
          <el-button type="success" link icon="EditPen" @click="openDrawer('编辑', scope.row)">编辑</el-button>
          <el-button type="danger" link icon="Delete" @click="deleteData(scope.row)">删除</el-button>
        </template>
      </ProTable>
      <DepartmentDialog ref="dialogRef" />
    </div>
  </div>
</template>

<script lang="tsx" setup name="department">
import { ref, reactive } from "vue";
import { Department } from "@/type/department/index";
import { ColumnProps, ProTableInstance } from "@/components/ProTable/interface";
import moment from "moment";
import ProTable from "@/components/ProTable/index.vue";
import { useHandleData } from "@/hooks/useHandleData";
import {
  getDepartmentList,
  addDepartment,
  dropDepartment,
  updateDepartment,
} from "@/api/system/department/index";
import DepartmentDialog from "@/views/components/departmentDialog/index.vue";

// 表格配置项
const columns: ColumnProps<Department>[] = [
  { type: "index", label: "序号", width: "60" },
  { prop: "name", label: "部门名称", search: { el: "input" } },
  { prop: "contacts", label: "部门联系人" },
  { prop: "email", label: "邮箱" },
  { prop: "phone", label: "电话" },
  {
    prop: "status",
    label: "部门状态",
    search: { el: "select" },
    enum: [
      { label: "正常", value: 0 },
      { label: "停用", value: 1 },
    ],
    render: (scope) => {
      return (
        <el-switch
          v-model={scope.row.status}
          inline-prompt
          active-text="正常"
          inactive-text="停用"
          active-value={0}
          inactive-value={1}
          onClick={() => changeDepartmentStatus(scope.row)}
        />
      );
    },
  },
  {
    prop: "orderNum",
    label: "排序序号"
  },
  {
    prop: "createdAt",
    label: "创建时间",
    width: "180",
    render: (scope) => {
      return moment(scope.row.createdAt).format("YYYY-MM-DD HH:mm:ss");
    },
  },
  { prop: "operation", label: "操作", width: "300", fixed: "right" },
];

// 打开 dialog(新增、查看、编辑)
const dialogRef = ref<InstanceType<typeof DepartmentDialog> | null>(null);
// 获取 ProTable 元素，调用其获取刷新数据方法（还能获取到当前查询参数，方便导出携带参数）
const proTable = ref<ProTableInstance>();

const defaultParams = reactive<Department>({
  id: "",
  parentId: "",
  name: "",
  contacts: "",
  phone: "",
  email: "",
  orderNum: 1,
  status: 0,
});

const openDrawer = (title: string, rowData: Partial<Department> = {}, id?: string) => {
  if (title == '新增') {
    rowData.parentId = id
  }
  const params = {
    title,
    rowData: { ...rowData },
    isView: title === "查看",
    api:
      title === "新增"
        ? addDepartment
        : title === "编辑"
          ? updateDepartment
          : undefined,
    getTableList: proTable.value?.getTableList,
  };
  dialogRef.value?.acceptParams(params);
};

//删除部门信息
const deleteData = async (row: Department) => {
  let ids: Array<string> = []
  recursiveSearchAllId(row, ids)
  if (ids.length > 1) {
    await useHandleData(dropDepartment, ids, `删除操作将${row.name}的子项也一并删除`)
  } else {
    await useHandleData(dropDepartment, ids, `删除${row.name}`)
  }
  //重新获取数据
  proTable.value?.getTableList();
};

//递归寻找出某个部门或公司下面的所以子公司或部门的id
const recursiveSearchAllId = (row: Department, ids: Array<string>) => {
  if (row.id) {
    ids.push(row.id)
  }
  if (row.children && row.children.length) {
    row.children.map(item => {
      recursiveSearchAllId(item, ids)
    })
  }

}

//修改部门的状态
const changeDepartmentStatus = async (row: Department) => {
  row.status = row.status == 0 ? 0 : 1
  await useHandleData(updateDepartment, row, `切换【${row.name}】状态`)
}
</script>
<style lang="scss" scoped>
.table-box {
  width: 100%;
  height: 100%;
}
</style>

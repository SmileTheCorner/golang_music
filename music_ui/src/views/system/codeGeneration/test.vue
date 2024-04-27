/** * 字典数值表-index.vue页面 * @author 钱 * @date 2024-01-05 15:24:29 *
@moduleName : DictData */

<template>
  <div class="app-container">
    <div class="table-box">
      <ProTable
        ref="proTable"
        title="字典数值表"
        rowKey="id"
        :indent="30"
        :columns="columns"
        :requestApi="getDictDataList"
        :requestAuto="true"
        :searchCol="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 4 }"
      >
        <!-- 表格 header 按钮 -->
        <template #tableHeader>
          <el-button
            type="primary"
            icon="CirclePlus"
            @click="openDrawer('新增', defaultParams, '')"
            >新增</el-button
          >
        </template>
        <!-- 表格操作 -->
        <template #operation="scope">
          <el-button
            link
            icon="CirclePlus"
            @click="openDrawer('新增', defaultParams, scope.row.id)"
            >新增</el-button
          >
          <el-button
            type="primary"
            link
            icon="View"
            @click="openDrawer('查看', scope.row)"
            >查看</el-button
          >
          <el-button
            type="success"
            link
            icon="EditPen"
            @click="openDrawer('编辑', scope.row)"
            >编辑</el-button
          >
          <el-button
            type="danger"
            link
            icon="Delete"
            @click="deleteData(scope.row)"
            >删除</el-button
          >
        </template>
      </ProTable>
      <DepartmentDialog ref="dialogRef" />
    </div>
  </div>
</template>

<script lang="tsx" setup name="department">
import { ref, reactive } from "vue";
import { DictData } from "@/type/dictData/index";
import { ColumnProps, ProTableInstance } from "@/components/ProTable/interface";
import moment from "moment";
import ProTable from "@/components/ProTable/index.vue";
import { useHandleData } from "@/hooks/useHandleData";
import {
  addDictData,
  dropDictData,
  editDictData,
  getDictDataList,
} from "@/api/system/dictData/index";
import DepartmentDialog from "@/views/components/departmentDialog/index.vue";

// 表格配置项
const columns: ColumnProps<DictData>[] = [
  { type: "selection", width: "60" },

  { prop: "create_by", label: "创建者" },

  { prop: "created_at", label: "创建时间" },

  { prop: "deleted_at", label: "删除时间" },

  { prop: "dict_label", label: "字典标签" },

  { prop: "dict_sort", label: "字典排序" },

  { prop: "dict_type", label: "字典类型" },

  { prop: "dict_value", label: "字典值" },

  { prop: "id", label: "主键id" },

  { prop: "remark", label: "备注" },

  { prop: "status", label: "状态（0正常 1停用）" },

  { prop: "update_by", label: "修改者" },

  { prop: "updated_at", label: "更新时间" },

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

const openDrawer = (
  title: string,
  rowData: Partial<Department> = {},
  id?: string
) => {
  if (title == "新增") {
    rowData.parentId = id;
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
  let ids: Array<string> = [];
  recursiveSearchAllId(row, ids);
  if (ids.length > 1) {
    await useHandleData(
      dropDepartment,
      ids,
      `删除操作将${row.name}的子项也一并删除`
    );
  } else {
    await useHandleData(dropDepartment, ids, `删除${row.name}`);
  }
  //重新获取数据
  proTable.value?.getTableList();
};

//递归寻找出某个部门或公司下面的所以子公司或部门的id
const recursiveSearchAllId = (row: Department, ids: Array<string>) => {
  if (row.id) {
    ids.push(row.id);
  }
  if (row.children && row.children.length) {
    row.children.map((item) => {
      recursiveSearchAllId(item, ids);
    });
  }
};

//修改部门的状态
const changeDepartmentStatus = async (row: Department) => {
  row.status = row.status == 0 ? 0 : 1;
  await useHandleData(updateDepartment, row, `切换【${row.name}】状态`);
};
</script>
<style lang="scss" scoped>
.table-box {
  width: 100%;
  height: 100%;
}
</style>

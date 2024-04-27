<template>
  <div class="app-container">
    <div class="table-box">
      <ProTable
        ref="proTable"
        title="菜单列表"
        rowKey="id"
        :indent="30"
        :columns="columns"
        :requestApi="getAuthMenuListApi"
        :requestAuto="true"
        :pagination="false"
        :searchCol="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 4 }"
      >
        <!-- 表格 header 按钮 -->
        <template #tableHeader>
          <el-button
            type="primary"
            icon="CirclePlus"
            @click="openDrawer('新增', defaultParameters, '')"
            >新增</el-button
          >
        </template>
        <!-- 表格操作 -->
        <template #operation="scope">
          <el-button
            type="primary"
            link
            icon="CirclePlus"
            @click="openDrawer('新增', defaultParameters, scope.row.meta.id)"
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
            type="primary"
            link
            icon="EditPen"
            @click="openDrawer('编辑', scope.row)"
            >编辑</el-button
          >
          <el-button
            type="primary"
            link
            icon="Delete"
            @click="deleteData(scope.row)"
            >删除</el-button
          >
        </template>
      </ProTable>
      <MenuDialog ref="drawerRef" />
    </div>
  </div>
</template>

<script lang="tsx" setup name="menu">
import {
  addMenu,
  delMenu,
  editMenu,
  getAuthMenuListApi,
} from "@/api/system/menu/index";
import { ref, reactive } from "vue";
import { ElMessageBox, ElMessage } from "element-plus";
import { ColumnProps, ProTableInstance } from "@/components/ProTable/interface";
import ProTable from "@/components/ProTable/index.vue";
import MenuDialog from "@/views/components/menuDialog/index.vue";
import SvgIcon from "@/components/SvgIcon/index.vue";
import { useHandleData } from "@/hooks/useHandleData";

// 获取 ProTable 元素，调用其获取刷新数据方法（还能获取到当前查询参数，方便导出携带参数）
const proTable = ref<ProTableInstance>();

//新增菜单默认选项
const defaultParameters = reactive<Menu.MenuOptions>({
  meta: {
    isFrame: 1,
    isFull: 0,
    isCache: 0,
    visible: 0,
    status: 0,
    orderNum: 1,
  },
});
// 表格配置项
const columns: ColumnProps<Menu.MenuOptions>[] = [
  { type: "index", label: "序号", width: "60" },
  {
    prop: "name",
    label: "名称",
    width: "200",
    align: "left",
  },
  { prop: "meta.title", label: "标题", width: "150", search: { el: "input" } },
  {
    prop: "meta.icon",
    label: "图标",
    render: (scope) => {
      return <SvgIcon iconClass={scope.row.meta?.icon}></SvgIcon>;
    },
  },
  { prop: "path", label: "路径", width: "200", align: "left" },
  { prop: "component", label: "组件路径", width: "250", align: "left" },
  {
    prop: "meta.isFrame",
    label: "外部链接",
    render: (scope) => {
      return scope.row.meta?.isFrame == 0 ? (
        <el-tag class="ml-2" type="warning">
          是
        </el-tag>
      ) : (
        <el-tag class="ml-2" type="danger">
          否
        </el-tag>
      );
    },
  },
  {
    prop: "meta.menuType",
    label: "类型",
    search: { el: "select" },
    enum: [
      { label: "目录", value: "M" },
      { label: "菜单", value: "C" },
      { label: "按钮", value: "F" },
    ],
    render: (scope) => {
      return scope.row.meta?.menuType == "M" ? (
        <el-tag>目录</el-tag>
      ) : scope.row.meta?.menuType == "C" ? (
        <el-tag class="ml-2" type="success">
          菜单
        </el-tag>
      ) : (
        <el-tag class="ml-2" type="info">
          按钮
        </el-tag>
      );
    },
  },
  {
    prop: "meta.isCache",
    label: "缓存",
    width: "100",
    render: (scope) => {
      return (
        <el-switch
          v-model={scope.row.meta.isCache}
          inline-prompt
          active-text="缓存"
          inactive-text="不缓存"
          active-value={0}
          inactive-value={1}
          onClick={() => handleIsCacheChange(scope.row)}
        />
      );
    },
  },
  {
    prop: "meta.visible",
    label: "显示掩藏",
    render: (scope) => {
      return (
        <el-switch
          v-model={scope.row.meta.visible}
          inline-prompt
          active-text="显示"
          inactive-text="掩藏"
          active-value={0}
          inactive-value={1}
          onClick={() => handleVisibleChange(scope.row)}
        />
      );
    },
  },
  {
    prop: "meta.status",
    label: "状态",
    search: { el: "select" },
    enum: [
      { label: "正常", value: 0 },
      { label: "停用", value: 1 },
    ],
    render: (scope) => {
      return (
        <el-switch
          v-model={scope.row.meta.status}
          inline-prompt
          active-text="正常"
          inactive-text="停用"
          active-value={0}
          inactive-value={1}
          onClick={() => handleStatus(scope.row)}
        />
      );
    },
  },
  { prop: "meta.perms", label: "权限", width: "100" },
  { prop: "operation", label: "操作", width: "300", fixed: "right" },
];

//改变菜单掩藏状态
const handleVisibleChange = async (row: Menu.MenuOptions) => {
  await useHandleData(
    editMenu,
    { id: row.meta?.id, visible: row.meta?.visible == 0 ? 0 : 1 },
    `切换【${row.name}】菜单掩藏状态`
  );
  proTable.value?.getTableList();
};

//改变菜单正常和停用状态
const handleStatus = async (row: Menu.MenuOptions) => {
  await useHandleData(
    editMenu,
    { id: row.meta?.id, status: row.meta?.status == 0 ? 0 : 1 },
    `切换【${row.name}】菜单启用状态`
  );
};

//修改菜单缓存状态
const handleIsCacheChange = async (row: Menu.MenuOptions) => {
  await useHandleData(
    editMenu,
    { id: row.meta?.id, isCache: row.meta?.isCache == 0 ? 0 : 1 },
    `切换【${row.name}】菜单缓存状态`
  );
};

// 打开 drawer(新增、查看、编辑)
const drawerRef = ref<InstanceType<typeof MenuDialog> | null>(null);

const openDrawer = (
  title: string,
  rowData: Partial<Menu.MenuOptions> = {},
  id?: string
) => {
  if (rowData.meta) {
    rowData.meta.parentId = id;
  }
  const params = {
    title,
    rowData: { ...rowData },
    isView: title === "查看",
    api: title === "新增" ? addMenu : title === "编辑" ? editMenu : undefined,
    getTableList: proTable.value?.getTableList,
  };
  drawerRef.value?.acceptParams(params);
};

/** 删除按钮操作 */
function deleteData(row: Menu.MenuOptions) {
  ElMessageBox.confirm("确定要删除当前行信息吗？", "删除", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  })
    .then(() => {
      try {
        delMenu({ id: row.meta?.id || "" });
        ElMessage({
          type: "success",
          message: "删除成功",
        });
        proTable.value?.getTableList();
      } catch (error) {}
    })
    .catch(() => {
      ElMessage({
        type: "info",
        message: "取消删除",
      });
    });
}
</script>
<style lang="scss" scoped>
.table-box {
  width: 100%;
  height: 100%;
}
</style>

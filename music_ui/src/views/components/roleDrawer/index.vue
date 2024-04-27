<template>
  <el-drawer
    v-model="drawerVisible"
    :title="drawerProps.title"
    direction="rtl"
    size="45%"
  >
    <el-form
      ref="ruleFormRef"
      :model="drawerProps.rowData"
      :rules="rules"
      :disabled="drawerProps.isView"
      label-width="120px"
      :size="formSize"
      status-icon
    >
      <el-row>
        <el-col :span="12">
          <el-form-item label="角色名称" prop="name">
            <el-input v-model="drawerProps.rowData.name" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="角色标识" prop="code">
            <el-input v-model="drawerProps.rowData.code" />
          </el-form-item>
        </el-col>

        <el-col :span="12">
          <el-form-item label="显示顺序">
            <el-input-number v-model="drawerProps.rowData.sort" :min="1" />
          </el-form-item>
        </el-col>

        <el-col :span="12">
          <el-form-item label="数据权限">
            <el-select
              v-model="drawerProps.rowData.dataScope"
              class="m-2"
              placeholder="请选择"
              size="large"
            >
              <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
        </el-col>

        <el-col :span="12">
          <el-form-item label="角色状态">
            <el-switch
              v-model="drawerProps.rowData.status"
              inline-prompt
              :active-value="0"
              :inactive-value="1"
              active-text="是"
              inactive-text="否"
            />
          </el-form-item>
        </el-col>

        <el-col :span="24">
          <el-form-item label="备注">
            <el-input
              v-model="drawerProps.rowData.remark"
              :rows="2"
              type="textarea"
              placeholder="请输入备注信息"
            />
          </el-form-item>
        </el-col>

        <el-col :span="12">
          <el-form-item label="菜单权限">
            <div class="border">
              <el-tree
                :data="menuOptions"
                show-checkbox
                node-key="value"
                default-expand-all
                @check="changeMenuList"
              >
              </el-tree>
            </div>
          </el-form-item>
        </el-col>

        <el-col :span="12" v-show="drawerProps.rowData.dataScope == 4">
          <el-form-item label="数据权限">
            <div class="border">
              <el-tree
                :props="{ label: 'name', children: 'children' }"
                :data="departmentOptions"
                show-checkbox
                node-key="id"
                default-expand-all
                @check="changeDepartmentList"
              >
              </el-tree>
            </div>
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
    <template #footer>
      <div class="drawer_footer">
        <el-button type="primary" @click="submit(ruleFormRef)">确认</el-button>
        <el-button type="danger" @click="cancel(ruleFormRef)">取消</el-button>
      </div>
    </template>
  </el-drawer>
</template>
<script setup lang="ts" name="roleDrawer">
import { ref, reactive, onMounted, getCurrentInstance } from "vue";
import { Role } from "@/type/role/index";
import type { ElForm } from "element-plus";
import { addMenuTree } from "@/api/system/menu/index";
import { departmentTree } from "@/api/system/department/index";

const message =
  getCurrentInstance()?.appContext.config.globalProperties.$message;

type FormInstance = InstanceType<typeof ElForm>;
type FormRules = InstanceType<typeof ElForm>;

interface DrawerProps {
  title: string; //标题
  isView: boolean; //是否是查看
  rowData: Partial<Role>; //数据
  api?: (params: any) => Promise<any>; //请求的api
  getTableList?: () => void; //重新获取数据接口
}

//控制抽屉的显示和隐藏
const drawerVisible = ref<boolean>(false);

//接收抽屉数据
let drawerProps = ref<DrawerProps>({
  isView: false,
  title: "",
  rowData: {},
});

const options = [
  {
    value: 1,
    label: "全部数据",
  },
  {
    value: 2,
    label: "本部门数据",
  },
  {
    value: 3,
    label: "本部门及以下数据",
  },
  {
    value: 4,
    label: "自定义数据",
  },
];

//接收父组件传过来的参数
const acceptParams = (params: DrawerProps) => {
  drawerVisible.value = true;
  drawerProps.value = params;
};

//表单内组件大小
const formSize = ref("large");
const ruleFormRef = ref<FormInstance>();

//表单校验规则
const rules = reactive<FormRules>({
  name: [{ required: true, message: "请填写角色名称", trigger: "blur" }],
  code: [{ required: true, message: "请填写角色标识", trigger: "blur" }],
});

//菜单下拉选项
const menuOptions = ref<Array<{ [key: string]: any }>>([]);
//部门下来树
const departmentOptions = ref<Array<{ [key: string]: any }>>([]);
/** 查询菜单下拉树结构 */
const getTreeSelect = async () => {
  menuOptions.value = [];
  try {
    const data = await addMenuTree();
    menuOptions.value = data as { [Key: string]: any }[];
  } catch (error) {}
};
//查询部门树
const getDepartmentTree = async () => {
  departmentOptions.value = [];
  try {
    const data = await departmentTree();
    departmentOptions.value = data as { [Key: string]: any }[];
  } catch (error) {}
};
//菜单权限复选框被选中时触发
const changeMenuList = (item: any, status: any) => {
  drawerProps.value.rowData.menuList = status.checkedKeys;
};

//资源权限复选框被选中时触发
const changeDepartmentList = (item: any, status: any) => {
  drawerProps.value.rowData.dataScopeDept = status.checkedKeys;
};

//提交请求
const submit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid: boolean, fields: any) => {
    if (valid) {
      try {
        await drawerProps.value.api!(drawerProps.value.rowData);
        //刷先列表数据
        drawerProps.value.getTableList!();
        message.success(`${drawerProps.value.rowData.name}添加成功`);
        //关闭角色抽屉
        drawerVisible.value = false;
        //清除表单
        formEl.resetFields();
      } catch (error) {}
    } else {
      message.error("请填写带*号的选项。。。");
      console.log("error submit!", fields);
    }
  });
};

//取消数据提交
const cancel = (formEl: FormInstance | undefined) => {
  //关闭角色抽屉
  drawerVisible.value = false;
  if (!formEl) return;
  formEl.resetFields();
};

onMounted(() => {
  getTreeSelect();
  getDepartmentTree();
});

//向父组件暴露方法
defineExpose({
  acceptParams,
});
</script>
<style lang="scss" scoped>
.border {
  box-shadow: 0 0 0 1px var(--el-input-border-color, var(--el-border-color))
    inset;
  border-radius: var(--el-input-border-radius, var(--el-border-radius-base));
  padding: 20px;
}
</style>

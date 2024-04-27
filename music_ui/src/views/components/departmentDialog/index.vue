<template>
  <el-dialog
    v-model="centerDialogVisible"
    :title="dialogProps.title"
    width="40%"
    align-center
  >
    <el-form
      ref="ruleFormRef"
      :model="dialogProps.rowData"
      :rules="rules"
      label-width="auto"
      class="demo-ruleForm"
      :size="formSize"
      status-icon
      :disabled="dialogProps.isView"
    >
      <el-row>
        <el-col :span="12">
          <el-form-item label="上级部门">
            <el-tree-select
              v-model="dialogProps.rowData.parentId"
              :data="menuOptions"
              check-strictly
              :props="{ label: 'name', value: 'id', children: 'children' }"
              placeholder="请选择上级部门"
            >
            </el-tree-select>
          </el-form-item>
        </el-col>

        <el-col :span="12">
          <el-form-item label="部门名称" prop="name">
            <el-input v-model="dialogProps.rowData.name" />
          </el-form-item>
        </el-col>

        <el-col :span="24">
          <el-form-item label="部门联系人">
            <el-input v-model="dialogProps.rowData.contacts" />
          </el-form-item>
        </el-col>

        <el-col :span="24">
          <el-form-item label="电话">
            <el-input v-model="dialogProps.rowData.phone" />
          </el-form-item>
        </el-col>

        <el-col :span="24">
          <el-form-item label="邮箱">
            <el-input v-model="dialogProps.rowData.email" />
          </el-form-item>
        </el-col>

        <el-col :span="12">
          <el-form-item label="排序序号">
            <el-input-number v-model="dialogProps.rowData.orderNum" />
          </el-form-item>
        </el-col>

        <el-col :span="12">
          <el-form-item label="状态">
            <el-switch
              v-model="dialogProps.rowData.status"
              inline-prompt
              active-text="正常"
              inactive-text="停用"
              :active-value="0"
              :inactive-value="1"
            />
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="cancel(ruleFormRef)">取消</el-button>
        <el-button type="primary" @click="makeSure(ruleFormRef)">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup name="departmentDialog">
import { reactive, ref } from "vue";
import { Department } from "@/type/department/index";
import { ElMessage } from "element-plus";
import { departmentTree } from "@/api/system/department/index";
import type { ElForm } from "element-plus";

type FormInstance = InstanceType<typeof ElForm>;
type FormRules = InstanceType<typeof ElForm>;

//定义对话框类型
interface DialogProps {
  title: string;
  isView: boolean;
  rowData: Partial<Department>;
  //添加和编辑接口
  api?: (params: any) => Promise<any>;
  //重新获取数据接口
  getTableList?: () => void;
}

//显示参数
const dialogProps = ref<DialogProps>({
  isView: false,
  title: "",
  rowData: {},
});

// 接收父组件传过来的参数
const acceptParams = (params: DialogProps) => {
  dialogProps.value = params;
  centerDialogVisible.value = true;
};

// 控制对话框的显示和隐藏
const centerDialogVisible = ref<boolean>(false);
// 表单内组件的尺寸大小
const formSize = ref("large");

const menuOptions = ref<Department[]>([]);
//获取部门结构树
const getDepartmentTree = async () => {
  menuOptions.value = [];
  try {
    const data = await departmentTree();
    menuOptions.value = data as Department[];
  } catch (error) {}
  const department = { id: "", name: "根部门" };
  menuOptions.value.unshift(department);
};
getDepartmentTree();

const ruleFormRef = ref<FormInstance>();

//表单校验规则
const rules = reactive<FormRules>({
  name: [
    { required: true, message: "请输入部门的名称", trigger: "blur" },
    { min: 3, max: 5, message: "部门名称必须为3-10字符", trigger: "blur" },
  ],
});
//清空表单以及校验规则
const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};
//取消表单提交
const cancel = (formEl: FormInstance | undefined) => {
  centerDialogVisible.value = false;
  resetForm(formEl);
};

//表单提交信息
const makeSure = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid: boolean, fields: any) => {
    if (valid) {
      try {
        dialogProps.value.api?.(dialogProps.value.rowData);
        ElMessage.success({
          message: `${dialogProps.value.rowData.name}操作成功！`,
        });
        //跟新列表数据，并关闭dialog窗口
        dialogProps.value.getTableList?.();
        centerDialogVisible.value = false;
      } catch (error) {}
    } else {
      ElMessage.error("请填写带*号的选项。。。");
      console.log("表单校验失败", fields);
    }
  });
};

//将需要暴露出去的数据与方法都可以暴露出去
defineExpose({
  acceptParams,
});
</script>

<template>
  <div class="container">
    <div class="search card">
      <ProForm :formOption="formColumn">
        <el-button type="success">查询</el-button>
        <el-button type="primary">重置</el-button>
      </ProForm>
    </div>
    <div class="header">
      <el-button type="primary" @click="importTable">导入</el-button>
    </div>
    <ProDialog :dialogConfig="{ visible: true, alignCenter: true }">
      <div>DialogVisible</div>
      <template #footer>
        <el-button type="primary">确定</el-button>
        <el-button type="primary">关闭</el-button>
      </template>
    </ProDialog>
  </div>
</template>
<script setup lang="tsx">
import { ref, onMounted } from "vue";
import ProForm from "@/components/ProForm/index.tsx";
import ProDialog from "@/components/ProDialog/index.tsx";
import { FormOption } from "@/components/ProForm/interface/index";
import { useForm } from "@/components/ProForm/useForm/index";
import { getAllTableList } from "@/api/system/codeGen/index";

const formColumn = ref<FormOption>({
  formItems: [
    {
      type: "input",
      label: "表名",
      prop: "name",
      placeholder: "请输入表名",
    },
    {
      type: "datepicker",
      label: "表名",
      prop: "name",
      dateTimeType: "monthrange",
      placeholder: "请输入表名",
    },
    {
      type: "timepicker",
      label: "表名",
      prop: "name",
      placeholder: "请输入表名",
    },
    {
      type: "select",
      label: "表名",
      prop: "name",
      placeholder: "请输入表名",
      options: [
        {
          label: "开启",
          value: 1,
        },
        {
          label: "关闭",
          value: 0,
        },
      ],
    },
    {
      type: "checkbox",
      label: "表名",
      prop: "name",
      placeholder: "请输入表名",
      options: [
        {
          label: "开启",
          value: 1,
        },
        {
          label: "关闭",
          value: 0,
        },
      ],
    },
    {
      type: "input",
      label: "表名",
      prop: "name",
      placeholder: "请输入表名",
    },
    {
      type: "input",
      label: "表名",
      prop: "name",
      placeholder: "请输入表名",
    },
  ],
  inline: true,
  labelPosition: "right",
  labelWidth: "auto",
  disabled: false,
});

const { getTableList, tableData } = useForm(getAllTableList);
onMounted(() => {
  getTableList();
});

//导入数据表
const importTable = () => {
  console.log("导入数据表");
};
</script>
<style scoped lang="scss">
.container {
  width: 100%;
  height: 100%;
}
</style>

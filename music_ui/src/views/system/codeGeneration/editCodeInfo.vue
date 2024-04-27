<template>
  <div class="card">
    <el-tabs v-model="activeName" class="tabs">
      <el-tab-pane label="基本信息" name="BasicInformation">
        <ProForm :formOption="basicColumn"> </ProForm>
      </el-tab-pane>
      <el-tab-pane label="字段信息" name="FieldInformation">
        <ProTable
          ref="proTable"
          title="部门列表"
          rowKey="id"
          :indent="30"
          :columns="fieldColumn"
          :operate="false"
          :data="data"
          :pagination="false"
          :searchCol="{ xs: 1, sm: 2, md: 2, lg: 3, xl: 4 }"
        >
        </ProTable>
      </el-tab-pane>
      <el-tab-pane label="生成信息" name="GenerateInformation">
        <ProForm :formOption="basicColumn"> </ProForm>
      </el-tab-pane>
    </el-tabs>
    <div class="operate">
      <el-button type="success">提交</el-button>
      <el-button type="primary" @click="goBack">返回</el-button>
    </div>
  </div>
</template>
<script lang="tsx" setup name="editCodeInfo">
import { ref, onMounted } from "vue";
import ProForm from "@/components/ProForm/index.tsx";
import ProTable from "@/components/ProTable/index.vue";
import { FormOption } from "@/components/ProForm/interface/index";
import { GenTableColumn } from "@/type/codeGeneration/index.ts";
import { ColumnProps } from "@/components/ProTable/interface";
import { getColumnsByName } from "@/api/system/codeGen/index";
import { getDictDataByType } from "@/api/system/dict/index";
import { getDictTypeList } from "@/api/system/dict/index";

const queryMethodData = ref<any[]>([]);
const goTypeData = ref<any[]>([]);
const textTypeData = ref<any[]>([]);
const dictTypeData = ref<any[]>([]);

const data = ref<any[]>([]);

onMounted(async () => {
  try {
    getColumnsInfo();
    let data = await getDictDataByType({ dictType: "query_method" });
    let goType = await getDictDataByType({ dictType: "golang_type" });
    let textType = await getDictDataByType({ dictType: "text_type" });
    let dictType = await getDictTypeList();
    queryMethodData.value.push(...data);
    goTypeData.value.push(...goType);
    textTypeData.value.push(...textType);
    dictTypeData.value.push(...dictType);
  } catch (error) {}
});

const basicColumn = ref<FormOption>({
  formItems: [
    {
      type: "input",
      label: "表名名称",
      prop: "name",
      placeholder: "请输入表名",
    },
    {
      type: "input",
      label: "表描述",
      prop: "name",
      placeholder: "请输入表名",
    },
    {
      type: "input",
      label: "结构体名称",
      prop: "name",
      placeholder: "请输入表名",
    },
    {
      type: "input",
      label: "作者",
      prop: "name",
      placeholder: "请输入表名",
    },
    {
      type: "textarea",
      label: "备注",
      prop: "name",
      placeholder: "请输入表名",
    },
  ],
  inline: false,
  labelPosition: "right",
  labelWidth: "auto",
  disabled: false,
});

// 表格配置项
const fieldColumn: ColumnProps<GenTableColumn>[] = [
  { type: "index", label: "序号", width: "60" },
  { prop: "columnName", label: "字段列名" },
  {
    prop: "columnComment",
    label: "字段描述",
    render: (scope) => {
      return (
        <el-input
          v-model={scope.row.columnComment}
          placeholder="请输入"
        ></el-input>
      );
    },
  },
  { prop: "columnType", label: "字段类型" },
  {
    prop: "goType",
    label: "golang类型",
    render: (scope) => {
      return (
        <el-select v-model={scope.row.goType} placeholder="请选择">
          {goTypeData.value.map((item: any) => {
            return (
              <el-option
                key={item.id}
                label={item.label}
                value={item.value}
              ></el-option>
            );
          })}
        </el-select>
      );
    },
  },
  {
    prop: "goProperty",
    label: "golang属性",
    render: (scope) => {
      return (
        <el-input
          v-model={scope.row.goProperty}
          placeholder="请输入"
        ></el-input>
      );
    },
  },
  {
    prop: "queryMethod",
    label: "查询方式",
    render: (scope) => {
      return (
        <el-select v-model={scope.row.queryMethod} placeholder="请选择">
          {queryMethodData.value.map((item: any) => {
            return (
              <el-option
                key={item.id}
                label={item.label}
                value={item.value}
              ></el-option>
            );
          })}
        </el-select>
      );
    },
  },
  {
    prop: "isRequired",
    label: "是否必填",
    render: (scope) => {
      return (
        <el-switch
          v-model={scope.row.isRequired}
          inline-prompt
          active-text="是"
          inactive-text="否"
        />
      );
    },
  },
  {
    prop: "htmlType",
    label: "文本显示类型",
    render: (scope) => {
      return (
        <el-select v-model={scope.row.htmlType} placeholder="请选择">
          {textTypeData.value.map((item: any) => {
            return (
              <el-option
                key={item.id}
                label={item.label}
                value={item.value}
              ></el-option>
            );
          })}
        </el-select>
      );
    },
  },
  {
    prop: "dictType",
    label: "字典类型",
    render: (scope) => {
      return (
        <el-select v-model={scope.row.dictType} placeholder="请选择">
          {dictTypeData.value.map((item: any) => {
            return (
              <el-option key={item.id} label={item.dictName} value={item.dictType}>
                <span style="float: left">{item.dictName}</span>
                <span style="float: right;color: var(--el-text-color-secondary);font-size: 13px;">
                  {item.dictType}
                </span>
              </el-option>
            );
          })}
        </el-select>
      );
    },
  },
];

const activeName = ref("BasicInformation");

//获取字段列信息
const getColumnsInfo = async () => {
  try {
    let param = history.state.params;
    let result = await getColumnsByName(param.tableName);
    data.value = result;
  } catch (error) {
    console.log("error===>", error);
  }
};
//返回上一个页面
const goBack = () => {
  history.go(-1);
};
</script>
<style lang="scss" scoped>
.card {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  .tabs {
    flex: 1;
    height: 100%;
    display: flex;
    flex-direction: column;
    ::v-deep(.el-tabs__content) {
      height: 100%;
    }
    ::v-deep(.el-tab-pane) {
      height: 100%;
    }
  }
  .operate {
    margin-top: 8px;
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
    align-items: center;
  }
}
</style>

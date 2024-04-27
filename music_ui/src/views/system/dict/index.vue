<template>
  <div class="app-container">
    <div class="left border-style">
      <div class="top">
        <div class="top-title">
          <el-icon :size="20">
            <Notebook />
          </el-icon>
          <span>字典类型</span>
        </div>
        <span class="top-icon">
          <el-icon :size="20" class="icon-refresh" @click="dictTypeRefresh">
            <Refresh />
          </el-icon>
          <el-icon :size="20" @click="openDictTypeDialog">
            <Plus />
          </el-icon>
        </span>
      </div>
      <div class="bottom">
        <div class="bottom-scrollbar">
          <el-scrollbar class="left-bottom-scrollbar">
            <el-tree
              default-expand-all
              :data="dictTree.filter((item) => item.id != '0')"
              @node-click="clickDictTree"
              :props="{ label: 'dictName', value: 'id', children: 'children' }"
            />
          </el-scrollbar>
        </div>
      </div>
    </div>
    <div class="right">
      <div class="top border-style">
        <div class="right-top-list">
          <el-auto-resizer>
            <template #default="{ height, width }">
              <el-table-v2
                :columns="columns"
                :data="dictTypeInfo"
                :width="width"
                :height="height"
              />
            </template>
          </el-auto-resizer>
        </div>
      </div>
      <div class="bottom border-style">
        <div class="right-bottom-header">
          <div>
            <el-button
              type="primary"
              plain
              icon="Plus"
              @click="openDictValueDialog"
              >新增</el-button
            >
            <el-button
              type="danger"
              plain
              icon="Delete"
              @click="deleteDictValues"
              >删除</el-button
            >
          </div>
          <div class="dict-type-title">字典值列表（{{ dictTypeTitle }}）</div>
        </div>
        <div class="right-bottom-main">
          <div class="right-bottom-list">
            <el-auto-resizer>
              <template #default="{ height, width }">
                <el-table-v2
                  :columns="dictValueColumns"
                  :data="dictValueList"
                  :width="width"
                  :height="height"
                />
              </template>
            </el-auto-resizer>
          </div>
        </div>
      </div>
    </div>
    <el-dialog
      v-model="dictTypeDialogVisible"
      title="新增字典类型"
      width="30%"
      @close="resetDictTypeForm(dictTypeFormRef)"
    >
      <el-form
        ref="dictTypeFormRef"
        :rules="dictTypeRules"
        :model="dictTypeForm"
        label-width="auto"
        label-position="left"
      >
        <el-form-item label="上级字典">
          <el-tree-select
            clearable
            v-model="dictTypeForm.parentId"
            :data="dictTree"
            check-strictly
            :props="{ label: 'dictName', value: 'id', children: 'children' }"
            placeholder="请选择上级字典类型"
          >
          </el-tree-select>
        </el-form-item>
        <el-form-item label="名称" prop="dictName">
          <el-input v-model="dictTypeForm.dictName" placeholder="请输入名称" />
        </el-form-item>
        <el-form-item label="标识" prop="dictType">
          <el-input v-model="dictTypeForm.dictType" placeholder="请输入标识" />
        </el-form-item>
        <el-form-item label="状态">
          <el-switch
            v-model="dictTypeForm.status"
            inline-prompt
            active-text="正常"
            inactive-text="停用"
            :active-value="0"
            :inactive-value="1"
          />
        </el-form-item>
        <el-form-item label="备注">
          <el-input
            v-model="dictTypeForm.remark"
            :rows="3"
            type="textarea"
            placeholder="请输入备注"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetDictTypeForm(dictTypeFormRef)"
            >取消</el-button
          >
          <el-button
            type="primary"
            @click="submitDictTypeForm(dictTypeFormRef)"
          >
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 新增字典值对话框 -->
    <el-dialog
      v-model="dictValueDialogVisible"
      title="新增字典值"
      width="40%"
      @close="resetDictValueForm(dictValueFormRef)"
    >
      <el-form
        ref="dictValueFormRef"
        :rules="dictValueRules"
        :model="dictValueForm"
        label-width="auto"
        label-position="left"
      >
        <el-form-item label="字典值名称" prop="label">
          <el-input
            v-model="dictValueForm.label"
            placeholder="请输入字典值名称"
          />
        </el-form-item>
        <el-form-item label="字典值" prop="value">
          <el-input
            :min="0"
            v-model="dictValueForm.value"
            placeholder="请输入字典值"
          />
        </el-form-item>
        <el-row>
          <el-col :span="12">
            <el-form-item label="状态">
              <el-switch
                v-model="dictValueForm.status"
                inline-prompt
                active-text="正常"
                inactive-text="停用"
                active-value="0"
                inactive-value="1"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="排序">
              <el-input-number
                v-model="dictValueForm.sort"
                placeholder="请输入排序"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="备注">
          <el-input
            v-model="dictValueForm.remark"
            :rows="3"
            type="textarea"
            placeholder="请输入备注"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="resetDictValueForm(dictValueFormRef)"
            >取消</el-button
          >
          <el-button
            type="primary"
            @click="submitDictValueForm(dictValueFormRef)"
          >
            确定
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>
<script lang="tsx" name="dict" setup>
import { onMounted, ref, nextTick, reactive, unref } from "vue";
import type { FunctionalComponent } from "vue";
import { Dict, DictValue, CellRenderProps } from "@/type/dict/index";
import {
  getDictTypeTree,
  getDictDataByType,
  addDictType,
  dropDictTypeByIds,
  updateDictType,
  updateDictValueData,
  addDictValueData,
  dropDictValueByIds,
} from "@/api/system/dict/index";
import moment from "moment";
import { ElInput, ElCheckbox } from "element-plus";
import type { Column, FormInstance, CheckboxValueType } from "element-plus";
import { ShowOrEdit, DropDownEdit } from "./index";
import { useHandleData } from "@/hooks/useHandleData";

onMounted(() => {
  dictTypeTree();
  calculateWidth();
  calculateDictTypeWidth();
});

//新增字典类型对话框
const dictTypeDialogVisible = ref(false);
//字典类型标题
const dictTypeTitle = ref<string>("");
const dictTypeFormRef = ref<FormInstance>();
const dictTypeForm = ref<Dict>({
  parentId: "",
  dictName: "",
  dictType: "",
  status: 0,
  remark: "",
});
const dictTypeRules = reactive({
  dictName: [
    { required: true, message: "请输入字典类型名称", trigger: "blur" },
  ],
  dictType: [
    { required: true, message: "请输入字典类型标识符", trigger: "blur" },
  ],
});
//新增字典类型数据提交
const submitDictTypeForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      try {
        addDictType(dictTypeForm.value);
        dictTypeDialogVisible.value = false;
        resetDictTypeForm(formEl);
        //从新获取数据
        dictTypeTree();
      } catch (error) {}
    } else {
      console.log("error submit!", fields);
    }
  });
};

//清除字典类型表单
const resetDictTypeForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
  dictTypeForm.value = {
    parentId: "",
    dictName: "",
    dictType: "",
    status: 0,
    remark: "",
  };
  dictTypeDialogVisible.value = false;
};

//重新刷新字典类型
const dictTypeRefresh = () => {
  dictTypeTree();
};

//字典类型树
const dictTree = ref<Dict[]>([]);
//字典类型详细信息
const dictTypeInfo = ref<Dict[]>([]);
//字典值列表
const dictValueList = ref<DictValue[]>([]);
//设置字典类型列的宽度
const dictTypeColumnWidth = ref(200);
const calculateDictTypeWidth = async () => {
  await nextTick(() => {
    let dom = document.documentElement.getElementsByClassName(
      "right-top-list"
    )[0] as HTMLElement;
    let width = dom.offsetWidth / 5;
    dictTypeColumnWidth.value = width;
  });
};
//字典类型表格配置列
const columns: Column<any>[] = [
  {
    key: "dictName",
    dataKey: "dictName",
    title: "名称",
    align: "center",
    width: dictTypeColumnWidth.value,
    cellRenderer: (props) => {
      return (
        <ShowOrEdit
          v-model:value={props}
          onChangeDate={(data) => saveDictTypeData(data)}
        ></ShowOrEdit>
      );
    },
  },
  {
    key: "dictType",
    dataKey: "dictType",
    title: "标识",
    align: "center",
    width: dictTypeColumnWidth.value,
    cellRenderer: (props) => {
      return (
        <ShowOrEdit
          v-model:value={props}
          onChangeDate={(data) => saveDictTypeData(data)}
        ></ShowOrEdit>
      );
    },
  },
  {
    key: "status",
    dataKey: "status",
    title: "状态",
    align: "center",
    width: dictTypeColumnWidth.value,
    cellRenderer: (props) => {
      return (
        <el-switch
          v-model={props.cellData}
          inline-prompt
          active-text="正常"
          inactive-text="停用"
          active-value={0}
          inactive-value={1}
          onClick={() => handleDictTypeStatus(props.rowData)}
        />
      );
    },
  },
  {
    key: "parentId",
    dataKey: "parentId",
    width: dictTypeColumnWidth.value,
    title: "上级目录",
    align: "center",
    cellRenderer: (props) => {
      return (
        <DropDownEdit
          onChangeDate={(data, newData) => saveDictTypeTreeData(data, newData)}
          v-model:value={props}
          v-model:options={dictTree.value}
        ></DropDownEdit>
      );
    },
  },
  {
    key: "remark",
    dataKey: "remark",
    title: "备注",
    align: "center",
    width: dictTypeColumnWidth.value,
    cellRenderer: (props) => {
      return (
        <ShowOrEdit
          v-model:value={props}
          onChangeDate={(data) => saveDictTypeData(data)}
        ></ShowOrEdit>
      );
    },
  },
  {
    key: "operations",
    title: "操作",
    cellRenderer: (props) => (
      <>
        <el-button
          type="danger"
          icon="Delete"
          onClick={() => dropDictType(props.rowData)}
        />
      </>
    ),
    width: dictTypeColumnWidth.value,
    align: "center",
  },
];
//保存字典类型数据
const saveDictTypeData = (data: CellRenderProps<any>) => {
  //获取到要更新的行数据
  const rowData = data.rowData;
  try {
    updateDictType(rowData);
    //刷新数据
    clickDictTree(dictTypeInfo.value[0]);
  } catch (error) {}
};
//保存字典上级目录数据
const saveDictTypeTreeData = async (
  data: CellRenderProps<any>,
  newData: any
) => {
  //获取到要更新的行数据
  const rowData = data.rowData;
  const dataName = rowData.dictName;
  const newDataName = newData.dictName;
  try {
    if (dataName != "根目录") {
      await updateDictType(rowData);
    }
    if (newDataName != undefined && newDataName != "根目录") {
      await updateDictType(newData);
    }
    //重新获取数据
    await dictTypeRefresh();
    //刷新数据
    await clickDictTree(dictTypeInfo.value[0]);
  } catch (error) {}
};
//删除字典类型数据
const dropDictType = async (rowData: Dict) => {
  let ids: Array<string> = [];
  recursiveSearchAllId(rowData, ids);
  // 判断要删除的字典类型下面是否还有子类型
  if (ids.length > 1) {
    await useHandleData(
      dropDictTypeByIds,
      ids,
      `删除操作将${rowData.dictName}的子项也一并删除`
    );
  } else {
    await useHandleData(dropDictTypeByIds, ids, `删除${rowData.dictName}`);
  }
  //删除成功后刷新数据
  await dictTypeTree();
};

//递归寻找出某个部门或公司下面的所以子公司或部门的id
const recursiveSearchAllId = (row: Dict, ids: Array<string>) => {
  row.id && ids.push(row.id);
  if (row.children && row.children.length) {
    row.children.map((item) => {
      recursiveSearchAllId(item, ids);
    });
  }
};

//修改字典类型状态
const handleDictTypeStatus = (row: Dict) => {
  row.status = row.status == 0 ? 1 : 0;
  try {
    updateDictType(row);
    //刷新数据
    clickDictTree(dictTypeInfo.value[0]);
  } catch (error) {}
};

//设置列的宽度
const columnWidth = ref(200);
const calculateWidth = async () => {
  await nextTick(() => {
    let dom = document.documentElement.getElementsByClassName(
      "right-bottom-list"
    )[0] as HTMLElement;
    let width = dom.offsetWidth / 7;
    columnWidth.value = width;
  });
};

type SelectionCellProps = {
  value: boolean;
  intermediate?: boolean;
  onChange: (value: CheckboxValueType) => void;
};

const SelectionCell: FunctionalComponent<SelectionCellProps> = ({
  value,
  intermediate = false,
  onChange,
}) => {
  return (
    <ElCheckbox
      onChange={onChange}
      modelValue={value}
      indeterminate={intermediate}
    />
  );
};
const dictValueSelection = new Set<string>();
//字典值表格配置列
const dictValueColumns: Column<any>[] = [
  {
    key: "selection",
    width: 50,
    cellRenderer: ({ rowData }) => {
      const onChange = (value: CheckboxValueType) => {
        rowData.checked = value;
      };
      return <SelectionCell value={rowData.checked} onChange={onChange} />;
    },
    headerCellRenderer: () => {
      const _data = unref(dictValueList);
      const onChange = (value: CheckboxValueType) => {
        dictValueList.value = _data.map((row) => {
          row.checked = Boolean(value);
          return row;
        });
      };
      const allSelected = _data.every((row) => row.checked);
      const containsChecked = _data.some((row) => row.checked);
      dictValueSelection.clear();
      _data.map((row) => {
        if (row.checked && row.id) {
          dictValueSelection.add(row.id);
        } else if (!row.checked && row.id) {
          dictValueSelection.delete(row.id);
        }
      });

      return (
        <SelectionCell
          value={allSelected}
          intermediate={containsChecked && !allSelected}
          onChange={onChange}
        />
      );
    },
  },
  {
    key: "label",
    dataKey: "label",
    title: "字典名称",
    align: "center",
    width: columnWidth.value,
    cellRenderer: (props) => {
      return (
        <ShowOrEdit
          v-model:value={props}
          onChangeDate={(data) => saveDictValueData(data)}
        ></ShowOrEdit>
      );
    },
  },
  {
    key: "value",
    dataKey: "value",
    title: "字典值",
    align: "center",
    minWidth: 100,
    width: columnWidth.value,
    cellRenderer: (props) => {
      return (
        <ShowOrEdit
          v-model:value={props}
          onChangeDate={(data) => saveDictValueData(data)}
        ></ShowOrEdit>
      );
    },
  },
  {
    key: "dictType",
    dataKey: "dictType",
    title: "字典类型",
    align: "center",
    minWidth: 150,
    width: columnWidth.value,
    cellRenderer: (props) => {
      return (
        <ShowOrEdit
          v-model:value={props}
          onChangeDate={(data) => saveDictValueData(data)}
        ></ShowOrEdit>
      );
    },
  },
  {
    key: "sort",
    dataKey: "sort",
    title: "字典值排序",
    align: "center",
    minWidth: 100,
    width: columnWidth.value,
    cellRenderer: (props) => {
      return (
        <ShowOrEdit
          v-model:value={props}
          onChangeDate={(data) => saveDictValueData(data)}
        ></ShowOrEdit>
      );
    },
  },
  {
    key: "status",
    dataKey: "status",
    title: "字典值状态",
    align: "center",
    minWidth: 100,
    width: columnWidth.value,
    cellRenderer: (props) => {
      return (
        <el-switch
          v-model={props.cellData}
          inline-prompt
          active-text="正常"
          inactive-text="停用"
          active-value="0"
          inactive-value="1"
          onClick={() => handleDictValueStatus(props.rowData)}
        />
      );
    },
  },
  {
    key: "remark",
    dataKey: "remark",
    title: "备注",
    align: "center",
    width: columnWidth.value,
    cellRenderer: (props) => {
      return (
        <ShowOrEdit
          v-model:value={props}
          onChangeDate={(data) => saveDictValueData(data)}
        ></ShowOrEdit>
      );
    },
  },
  {
    key: "createdAt",
    dataKey: "createdAt",
    title: "创建时间",
    align: "center",
    width: columnWidth.value + 100,
    cellRenderer: (props) => {
      return (
        <span>{moment(props.cellData).format("YYYY-MM-DD HH:mm:ss")}</span>
      );
    },
  },
  {
    key: "operations",
    title: "操作",
    cellRenderer: (props) => (
      <>
        <el-button
          type="danger"
          icon="Delete"
          onClick={() => handleDictValueDelete(props.rowData)}
        />
      </>
    ),
    width: columnWidth.value,
    align: "center",
  },
];

//刷新字典值列表
const refreshDictValue = async () => {
  if (dictTypeInfo.value[0].dictType) {
    const data = await getDictDataByType({
      dictType: dictTypeInfo.value[0].dictType,
    });
    dictValueList.value = [];
    dictValueList.value.push(...(data as unknown as DictValue[]));
  }
};

//修改字典值列表
const saveDictValueData = async (data: CellRenderProps<any>) => {
  //拿到需要更新的行数据
  const rowData = data.rowData;
  rowData.sort = Number(rowData.sort);
  try {
    await updateDictValueData(rowData);
    //修改成功刷新数据
    await refreshDictValue()
  } catch (error) {}
};

//获取字典类型树
const dictTypeTree = async () => {
  try {
    const data = await getDictTypeTree();
    dictTree.value = data as unknown as Dict[];
    let rootTree = {
      id: "0",
      parentId: "",
      dictName: "根目录",
      dictType: "root",
      remark: "根目录",
      status: 0,
      children: [],
    };
    dictTree.value.unshift(rootTree);
    await clickDictTree(dictTree.value[1]);
  } catch (error) {}
};

//点击字典类型树
const clickDictTree = async (data: Dict) => {
  dictTypeTitle.value = data.dictName ?? "";
  dictTypeInfo.value = [];
  dictTypeInfo.value.push(data);
  //根据字典类型获取字典的值
  try {
    if (data.dictType) {
      const datas = await getDictDataByType({ dictType: data.dictType });
      dictValueList.value = [];
      dictValueList.value.push(...(datas as unknown as DictValue[]));
    }
  } catch (error) {}
};

//修改字典值状态
const handleDictValueStatus = async (row: DictValue) => {
  row.status = row.status == "0" ? "1" : "0";
  try {
    await updateDictValueData(row);
    //修改成功刷新数据
    await refreshDictValue()
  } catch (error) {}
};

//删除字典值数据
const handleDictValueDelete = async (row: DictValue) => {
  const ids = [];
  if (row.id) {
    ids.push(row.id);
  }
  await useHandleData(dropDictValueByIds, ids, `删除${row.label}`);
  //刷新数据
  await refreshDictValue()
};
//打开字典类型新增对话框
const openDictTypeDialog = () => {
  dictTypeDialogVisible.value = true;
};

const dictValueDialogVisible = ref(false);
const dictValueFormRef = ref<FormInstance>();

//打开字典值新增对话框
const openDictValueDialog = () => {
  dictValueDialogVisible.value = true;
};
//重置字典值对话框
const resetDictValueForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
  dictValueDialogVisible.value = false;
  dictValueForm.value = {
    label: "",
    value: "",
    status: "0",
    remark: "",
    dictType: "",
    sort: 0,
  };
};

const dictValueForm = ref<DictValue>({
  label: "",
  value: "",
  status: "0",
  remark: "",
  dictType: "",
  sort: 0,
});
const dictValueRules = reactive({
  label: [{ required: true, message: "请输入字典值名称", trigger: "blur" }],
  value: [{ required: true, message: "请输入字典值", trigger: "blur" }],
});

//字典值新增对话框提交
const submitDictValueForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      dictValueForm.value.dictType = dictTypeInfo.value[0].dictType;
      try {
        await addDictValueData(dictValueForm.value);
        //关闭新增对话框
        dictValueDialogVisible.value = false;
        //重置表单
        await resetDictValueForm(formEl);
        await refreshDictValue()
      } catch (error) {}
    } else {
      console.log("error submit!", fields);
    }
  });
};

//批量删除字典值
const deleteDictValues = async () => {
  const ids = [...dictValueSelection];
  await useHandleData(dropDictValueByIds, ids, "删除字典值");
  //刷新数据
  if (dictTypeInfo.value[0].dictType) {
    try {
      const data = getDictDataByType({
        dictType: dictTypeInfo.value[0].dictType,
      });
      dictValueList.value = [];
      dictValueList.value.push(...(data as unknown as DictValue[]));
    } catch (error) {}
  }
};
</script>
<style lang="scss" scoped>
.app-container {
  display: flex;

  .left {
    padding: 10px;
    height: 100%;
    width: 17%;
    margin-right: 10px;
    display: flex;
    flex-direction: column;

    .top {
      display: flex;
      align-items: center;
      justify-content: space-between;

      .top-title {
        display: flex;
        align-items: center;
      }

      .top-icon {
        display: flex;
        justify-content: space-around;
        align-items: center;

        .icon-refresh {
          margin-right: 8px;
        }
      }
    }

    .bottom {
      flex: 1;
      margin-top: 10px;
      position: relative;

      .bottom-scrollbar {
        width: 100%;
        height: 100%;
        position: absolute;

        .left-bottom-scrollbar {
          height: 100%;
        }
      }
    }
  }

  .right {
    flex: 1;
    display: flex;
    flex-direction: column;

    .top {
      margin-bottom: 10px;
      height: 19%;
      position: relative;

      .right-top-list {
        box-sizing: border-box;
        padding: 5px;
        position: absolute;
        width: 100%;
        height: 100%;

        ::v-deep(.el-table-v2__row) {
          border-bottom: none;
        }
      }
    }

    .bottom {
      flex: 1;
      padding: 5px;
      display: flex;
      box-sizing: border-box;
      flex-direction: column;

      .right-bottom-header {
        display: flex;
        flex-direction: row;

        .dict-type-title {
          flex: 1;
          display: flex;
          flex-direction: row;
          justify-content: center;
          align-items: center;
          font-weight: bold;
          font-size: 18px;
        }
      }

      .right-bottom-main {
        flex: 1;
        position: relative;

        .right-bottom-list {
          position: absolute;
          width: 100%;
          height: 100%;
        }
      }
    }
  }
}
</style>

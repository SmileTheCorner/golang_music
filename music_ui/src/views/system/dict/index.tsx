import { h, defineComponent, ref, nextTick, reactive } from "vue";
import { ElInput, ElTreeSelect } from "element-plus";
import type { CellRenderProps } from "@/type/dict/index";
import { getDictTreeName } from "@/utils/dict";
import { swapProperty, isNodeInDomTree } from "@/utils/index";

export const ShowOrEdit = defineComponent({
  name: "editCell",
  //接收父组件传过来的值
  props: {
    value: { type: Object as () => CellRenderProps<any> },
  },
  emits: {
    changeDate: (val: CellRenderProps<any>) => val,
  },
  setup(props, ctx) {
    //将接收到的数据变成响应式数据
    const data = reactive<CellRenderProps<any>>(props.value!);
    //是否可以编辑
    const isEdit = ref(false);
    //输入框聚焦
    const inputRef = ref<HTMLInputElement>(null!);

    //进入编辑模式
    function handleOnClick() {
      isEdit.value = true;
      nextTick(() => {
        inputRef.value?.focus();
      });
    }

    //退出编辑模式
    function handleExitClick() {
      ctx.emit("changeDate", data);
      isEdit.value = false;
    }

    //输入框发生改变
    function handleChange(e: any) {
      data.rowData[data.column.dataKey!] = e;
    }

    return () =>
      h(
        "div",
        {
          onClick: handleOnClick,
        },
        isEdit.value
          ? h(ElInput, {
              ref: inputRef,
              modelValue: data.rowData[data.column.dataKey!],
              onInput: handleChange,
              onBlur: handleExitClick,
            })
          : data.rowData[data.column.dataKey!]
      );
  },
});

//下拉编辑
export const DropDownEdit = defineComponent({
  name: "selectEdit",
  //接收父组件传过来的值
  props: {
    value: { type: Object as () => CellRenderProps<any> },
    options: { type: Array as () => any[] },
  },
  emits: {
    changeDate: (val: CellRenderProps<any>, newVal: any) => val,
  },
  setup(props, ctx) {
    let conf = { label: "dictName", value: "id", children: "children" };
    //将接收到的数据变成响应式数据
    const data = reactive<CellRenderProps<any>>(props.value!);
    let newData = ref({});
    //是否可以编辑
    const isEdit = ref(false);
    //输入框聚焦
    const inputRef = ref<HTMLInputElement>(null!);
    //获取当前值
    let currentData = data.rowData[data.column.dataKey!];
    //点击进入编辑模式
    const handleClick = () => {
      isEdit.value = true;
      nextTick(() => {
        inputRef.value?.focus();
      });
    };
    //选中当前节点对象
    const currentChange = async (e: any) => {
      let id = e.id;
      let rowParentId = data.rowData[data.column.dataKey!];
      //判断当前修改的节点和选中的节点是否在一颗树内
      let inNode = await isNodeInDomTree(e, data.rowData, props.options);
      //如果在一颗树中交换各自的父亲ID,否则修改parentId
      if (inNode && id !== rowParentId) {
        if (e.parentId == rowParentId || e.parentId == "0") {
          data.rowData[data.column.dataKey!] = e.id;
        } else {
          await swapProperty(e, data.rowData, "parentId");
          newData.value = e;
        }
      } else {
        data.rowData[data.column.dataKey!] = e.id;
      }
    };
    //退出编辑模式
    const handleExitEdit = () => {
      ctx.emit("changeDate", data, newData.value);
      isEdit.value = false;
    };
    return () => {
      if (isEdit.value) {
        return (
          <div onClick={handleClick}>
            <ElTreeSelect
              onCurrentChange={currentChange}
              onBlur={handleExitEdit}
              modelValue={currentData}
              data={props.options}
              clearable
              check-strictly
              placeholder="请选择上级字典类型"
              props={conf}
            ></ElTreeSelect>
          </div>
        );
      } else {
        return (
          <div onClick={handleClick}>
            {getDictTreeName(props.options || [], currentData)}
          </div>
        );
      }
    };
  },
});

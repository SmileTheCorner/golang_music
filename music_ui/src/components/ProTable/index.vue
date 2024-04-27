<template>
  <div class="table-container">
    <!-- 树形选择 -->
    <div class="tree-filter card" v-show="treeFilter">
      <TreeFilter
        :id="treeFilterRowKey"
        :label="treeFilterLabel"
        :title="treeFilterTitle"
        :propsConfig="propsConfig"
        :data="treeFilterData"
        :request-api="treeFilterApi"
        :multiple="treeFilterMultiple"
        :default-value="initParam.id"
      ></TreeFilter>
    </div>
    <!-- 对表格进行封装 -->
    <div class="pro-table">
      <!-- 查询表单 card -->
      <SearchForm
        v-show="isShowSearch && showSearchIsEmpty"
        :search="search"
        :reset="reset"
        :columns="searchColumns"
        :search-param="searchParam"
        :search-col="searchCol"
      ></SearchForm>
      <!-- 表格内容 card -->
      <div id="table">
        <!-- 表格头部 操作按钮 -->
        <div id="table-header">
          <!-- 表格左侧操作按钮 -->
          <div class="table-header-button-lf">
            <slot name="tableHeader" :selected-list-ids="selectedListIds" :selected-list="selectedList" :is-selected="isSelected"></slot>
          </div>
          <!-- 表格右侧操作按钮 -->
          <div class="table-header-button-lr" v-if="operate">
            <slot name="toolButton">
              <el-button icon="Refresh" circle @click="getTableList" />
              <el-button
                v-if="columns.length"
                icon="Printer"
                circle
                @click="print"
              />
              <el-button
                v-if="columns.length"
                icon="Operation"
                circle
                @click="openColSetting"
              />
              <el-button
                v-if="searchColumns.length"
                icon="Search"
                circle
                @click="isShowSearch = !isShowSearch"
              />
            </slot>
          </div>
        </div>
        <div id="table-main">
          <!-- 表格主体 -->
          <el-table
            ref="tableRef"
            :border="border"
            :data="data ?? tableData"
            :row-key="rowKey"
            v-bind="$attrs"
            default-expand-all
            @selection-change="selectionChange"
          >
            <!-- 循环表头的字段 -->
            <template v-for="item in tableColumns" :key="item">
              <!-- selection || index || expand 用来控制多选、序号、数据展开的表头 -->
              <el-table-column
                v-if="
                  item.type &&
                  ['selection', 'index', 'expand'].includes(item.type)
                "
                v-bind="item"
                :align="item.align ?? 'center'"
                :reserve-selection="item.type == 'selection'"
              >
                <template v-if="item.type == 'expand'" #default="scope">
                  <component
                    :is="item.render"
                    v-bind="scope"
                    v-if="item.render"
                  >
                  </component>
                  <slot v-else :name="item.type" v-bind="scope"></slot>
                </template>
              </el-table-column>

              <!-- 其他正常的表格字段 -->
              <TableColumn
                v-if="!item.type && item.prop && item.isShow"
                :column="item"
              >
                <template v-for="slot in Object.keys($slots)" #[slot]="scope">
                  <slot :name="slot" v-bind="scope"></slot>
                </template>
              </TableColumn>
            </template>

            <!-- 插入表格最后一行的内容 -->
            <template #append>
              <slot name="append"> </slot>
            </template>

            <!-- 无数据 -->
            <template #empty>
              <div class="table-empty">
                <slot name="empty">
                  <img src="@/assets/images/notData.png" alt="notData" />
                  <div>暂无数据</div>
                </slot>
              </div>
            </template>
          </el-table>
        </div>
        <div id="pagination">
          <slot name="pagination">
            <Pagination
              v-if="pagination"
              :page-able="pageAble"
              :handle-size-change="handleSizeChange"
              :handle-current-change="handleCurrentChange"
            />
          </slot>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts" name="ProTable">
import { ElTable } from "element-plus";
import { ColumnProps, BreakPoint } from "./interface/index";
import TableColumn from "./module/TableColumn.vue";
import TreeFilter from "../TreeFilter/index.vue";
import SearchForm from "./module/searchForm/index.vue";
import Pagination from "./module/Pagination.vue";
import { useTable } from "@/hooks/useTable/index";
import { useSelection } from "@/hooks/useSelections/index";
import { ref, provide, reactive, onMounted } from "vue";

export interface ProTableProps {
  columns: ColumnProps[]; // 列配置项  ==> 必传
  data?: any[]; // 静态 table data 数据，若存在则不会使用 requestApi 返回的 data ==> 非必传
  requestApi?: (params: any) => Promise<any>; // 请求表格数据的 api ==> 非必传
  requestAuto?: boolean; // 是否自动执行请求 api ==> 非必传（默认为true）
  requestError?: (params: any) => void; // 表格 api 请求错误监听 ==> 非必传
  dataCallback?: (data: any) => any; // 返回数据的回调函数，可以对数据进行处理 ==> 非必传
  title?: string; // 表格标题，目前只在打印的时候用到 ==> 非必传
  pagination?: boolean; // 是否需要分页组件 ==> 非必传（默认为true）
  treeFilter?: boolean; //是否需要左侧树型组织树
  operate?: boolean; //是否需要表头的操作按钮
  treeFilterData?: any[]; //左侧树型组织树数据
  treeFilterTitle?: string; //左侧树型组织树标题
  treeFilterMultiple?: boolean; //左侧树是否支持多选
  treeFilterApi?: (params: any) => Promise<any>; //左侧树型组织树请求api
  treeFilterRowKey?: string; //每个树节点用来作为唯一标识的属性，整棵树应该是唯一的
  treeFilterLabel?: string; //侧树型组织树显示的label
  initParam?: any; // 初始化请求参数 ==> 非必传（默认为{}）
  border?: boolean; // 是否带有纵向边框 ==> 非必传（默认为true）
  toolButton?: boolean; // 是否显示表格功能按钮 ==> 非必传（默认为true）
  rowKey?: string; // 行数据的 Key，用来优化 Table 的渲染，当表格数据多选时，所指定的 id ==> 非必传（默认为 id）
  searchCol?: number | Record<BreakPoint, number>; // 表格搜索项 每列占比配置 ==> 非必传 { xs: 1, sm: 2, md: 2, lg: 3, xl: 4 }
}

// 接受父组件参数，配置默认值
const props = withDefaults(defineProps<ProTableProps>(), {
  columns: () => [],
  requestAuto: true,
  pagination: true,
  operate: true,
  initParam: {},
  border: true,
  toolButton: true,
  rowKey: "id",
  treeFilterRowKey: "id",
  treeFilterLabel: "label",
  treeFilterTitle: "树型数据",
  treeFilter: false,
  treeFilterMultiple: false,
  searchCol: () => ({ xs: 1, sm: 2, md: 2, lg: 3, xl: 4 }),
});

// 如果表格需要初始化请求参数，直接定义传给 ProTable(之后每次请求都会自动带上该参数，此参数更改之后也会一直带上，改变此参数会自动刷新表格数据)
const initParam = reactive({ id: "" });

// 表格 DOM 元素
const tableRef = ref<InstanceType<typeof ElTable>>();

// 表格多选 Hooks
const { selectionChange, selectedList, selectedListIds, isSelected } =
  useSelection(props.rowKey);
// 清空选中数据列表
const clearSelection = () => tableRef.value!.clearSelection();

// 表格操作 Hooks
const {
  search,
  reset,
  searchParam,
  getTableList,
  tableData,
  pageAble,
  handleCurrentChange,
  handleSizeChange,
} = useTable(
  props.requestApi,
  props.initParam,
  props.pagination,
  props.dataCallback,
  props.requestError
);

//判断搜索列表是否为空，如果为空则不显示
const showSearchIsEmpty = ref(true);
//初始化获取数据列表
onMounted(() => {
  if (props.requestAuto) {
    getTableList();
  }
  if (searchColumns.length == 0) {
    showSearchIsEmpty.value = false;
  }
});

// 接收 columns 并设置为响应式
const tableColumns = ref<ColumnProps[]>(props.columns);

// 是否显示搜索模块
const isShowSearch = ref(true);

// 定义 enumMap 存储 enum 值（避免异步请求无法格式化单元格内容 || 无法填充搜索下拉选择）
const enumMap = ref(new Map<string, { [key: string]: any }[]>());

provide("enumMap", enumMap);

const setEnumMap = async (col: ColumnProps) => {
  if (!col.enum) return;
  // 如果当前 enum 为后台数据需要请求数据，则调用该请求接口，并存储到 enumMap
  if (typeof col.enum !== "function")
    return enumMap.value.set(col.prop!, col.enum!);
  const { data } = await col.enum();
  enumMap.value.set(col.prop!, data);
};

// 扁平化 columns
const flatColumnsFunc = (
  columns: ColumnProps[],
  flatArr: ColumnProps[] = []
) => {
  columns.forEach(async (col) => {
    if (col._children?.length) flatArr.push(...flatColumnsFunc(col._children));
    flatArr.push(col);

    // 给每一项 column 添加 isShow && isFilterEnum 默认属性
    col.isShow = col.isShow ?? true;
    col.isFilterEnum = col.isFilterEnum ?? true;

    // 设置 enumMap
    setEnumMap(col);
  });
  return flatArr.filter((item) => !item._children?.length);
};

// flatColumns
const flatColumns = ref<ColumnProps[]>();
flatColumns.value = flatColumnsFunc(tableColumns.value);

// 过滤需要搜索的配置项
const searchColumns = flatColumns.value.filter(
  (item) => item.search?.el || item.search?.render
);

//打印表格数据
const print = () => {};
//对表格列显示掩藏设置
const openColSetting = () => {};

//将对应的方法暴露给父组件使用
defineExpose({
  getTableList,
  clearSelection,
  isSelected,
  selectedList,
  selectedListIds
});
</script>
<style scoped lang="scss">
.table-container {
  display: flex;
  height: 100%;
  width: 100%;
  .tree-filter {
    height: 100%;
    margin-right: 10px;
    width: 220px;
  }
  .pro-table {
    flex: 1;
    display: flex;
    flex-direction: column;
    height: 100%;
    #table {
      flex: 1;
      display: flex;
      flex-direction: column;
      box-sizing: border-box;
      padding: 20px;
      background-color: var(--el-bg-color);
      border: 1px solid var(--el-border-color-light);
      box-shadow: 0 0 12px #0000000d;
      border-radius: 10px;
      #table-header {
        box-sizing: border-box;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
      }
      #table-main {
        box-sizing: border-box;
        margin: 20px 0;
        flex: 1;
        position: relative;
        .table-header-style {
          background: "#eef1f6";
          color: "#606266";
        }
        .el-table {
          position: absolute;
          height: 100%;
          :deep(.el-scrollbar__wrap) {
            height: 100%;
          }
          :deep(.el-scrollbar__view) {
            height: 100%;
          }
        }
        .table-empty {
          display: flex;
          justify-content: center;
          align-items: center;
          flex-direction: column;
        }
      }
      #pagination {
        box-sizing: border-box;
        margin: 0 auto;
      }
    }
  }
}
</style>

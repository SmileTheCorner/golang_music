<template>
  <div class="singer-list">
    <ProTable
      ref="proTable"
      title="歌手列表"
      :columns="columns"
      :requestApi="getSingerList"
      :pagination="true"
    >
      <template #tableHeader>
        <el-button type="primary" icon="CirclePlus" @click="openDrawer('新增')">
          新增歌词
        </el-button>
      </template>

      <template #operation="scope">
        <el-button
          type="primary"
          plain
          size="small"
          icon="View"
          @click="openDrawer('查看', scope.row)"
        >
          查看
        </el-button>
        <el-button
          type="success"
          plain
          size="small"
          icon="EditPen"
          @click="openDrawer('编辑', scope.row)"
        >
          编辑
        </el-button>
        <el-button
          type="danger"
          plain
          size="small"
          icon="Delete"
          @click="handelDelete(scope.row)"
        >
          删除
        </el-button>
      </template>
    </ProTable>
    <SingerDrawer ref="singerDrawer"></SingerDrawer>
  </div>
</template>
<script setup lang="tsx" name="SongList">
import { ref, onBeforeMount } from "vue";
import { Singer } from "@/type/singer/index";
import {
  addSinger,
  editSinger,
  getSingerList,
  deleteSingerByIds,
} from "@/api/singer/index";
import { getDictDataByType } from "@/api/system/dict/index";
import {
  ColumnProps,
  ProTableInstance,
} from "@/components/ProTable/interface/index";
import ProTable from "@/components/ProTable/index.vue";
import { useHandleData } from "@/hooks/useHandleData/index";
import SingerDrawer from "@/views/components/singerDrawer/index.vue";
import moment from "moment";
import {getDictName} from "@/utils/dict"

//歌手分类
const singerClassify = ref<any[]>([]);
onBeforeMount(async () => {
  try {
    const data = await getDictDataByType({ dictType: "singer_classify" })
    singerClassify.value.push(...data as unknown as any[]);
  } catch (error) {
    
  }
});

// 获取 ProTable 元素，调用其获取刷新数据方法（还能获取到当前查询参数，方便导出携带参数）
const proTable = ref<ProTableInstance>();

// 表格配置项
const columns: ColumnProps<any>[] = [
  { type: "selection", fixed: "left", width: 60 },
  { type: "index", label: "序号", width: 80 },
  {
    prop: "singerName",
    label: "歌手名称",
    search: { el: "input" },
  },

  {
    prop: "singerEHName",
    label: "歌手英文名称",
  },
  {
    prop: "singerAvatar",
    label: "歌手头像",
    render: (scope) => {
      return (
        <>
          <el-image src={scope.row.singerAvatar} fit="cover" />
        </>
      );
    },
  },
  {
    prop: "singerClassify",
    label: "歌手分类",
    search: {
      el: "select",
    },
    enum: singerClassify.value,
    render: (scope) => {
      return getDictName(singerClassify.value, scope.row.singerClassify);
    },
  },
  {
    prop: "singerDescription",
    label: "歌手描述",
    showOverflowTooltip: true,
  },
  {
    prop: "createdAt",
    label: "创建时间",
    width: "180",
    render: (scope) => {
      return moment(scope.row.createdAt).format("YYYY-MM-DD HH:mm:ss");
    },
  },
  { prop: "operation", label: "操作", fixed: "right", width: 250 },
];

// 打开 drawer(新增、查看、编辑)
const singerDrawer = ref<InstanceType<typeof SingerDrawer> | null>(null);
const openDrawer = (title: string, row?: Partial<Singer>) => {
  let params = {
    title,
    isView: title === "查看",
    row: { ...row },
    api: title == "新增" ? addSinger : title == "编辑" ? editSinger : undefined,
    getTableList: proTable.value?.getTableList,
  };
  //打开抽屉
  singerDrawer.value?.acceptParams(params);
};
//删除数据
const handelDelete = async (row: Singer) => {
  const ids: string[] = [];
  ids.push(row.id);
  await useHandleData(deleteSingerByIds, ids, `删除【${row.songName}】歌曲`);
  //从新获取数据
  proTable.value?.getTableList();
};
</script>
<style scoped lang="scss">
.singer-list {
  width: 100%;
  height: 100%;
}
</style>

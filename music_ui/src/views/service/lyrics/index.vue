<template>
    <div class="lyrics-list">
      <ProTable
        ref="proTable"
        title="歌词列表"
        :columns="columns"
        :requestApi="getLyricsList"
        :pagination="true"
      >
        <template #tableHeader>
          <el-button type="primary" icon="CirclePlus" @click="openDrawer('新增')"> 新增歌词 </el-button>
        </template>
  
        <template #operation="scope">
          <el-button type="primary" plain size="small" icon="View" @click="openDrawer('查看',scope.row)">
            查看
          </el-button>
          <el-button type="success" plain size="small" icon="EditPen" @click="openDrawer('编辑',scope.row)">
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
      <LyricDrawer ref="lyricDrawer"></LyricDrawer>
    </div>
  </template>
  <script setup lang="tsx" name="SongList">
  import {ref } from "vue";
  import { Lyrics } from "@/type/lyrics/index";
  import { getLyricsList,deleteLyricsByIds,addLyrics,editLyrics } from "@/api/lyrics/index";
  import {
    ColumnProps,
    ProTableInstance,
  } from "@/components/ProTable/interface/index";
  import ProTable from "@/components/ProTable/index.vue";
  import { useHandleData } from "@/hooks/useHandleData/index";
  import LyricDrawer from "@/views/components/lyricsDrawer/index.vue"
  import moment from "moment";
  // 获取 ProTable 元素，调用其获取刷新数据方法（还能获取到当前查询参数，方便导出携带参数）
  const proTable = ref<ProTableInstance>();
  
  // 表格配置项
  const columns: ColumnProps<any>[] = [
    { type: "selection", fixed: "left", width: 60 },
    { type: "index", label: "序号", width: 80 },
    {
      prop: "songName",
      label: "歌曲名称",
      search:{el:"input"}
    },
    {
      prop: "lyricUrl",
      label: "歌曲地址",
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
  const lyricDrawer = ref<InstanceType<typeof LyricDrawer> | null>(null);
  const openDrawer = (title:string,row?:Partial<Lyrics>)=>{
       let params = {
        title,
        isView: title === "查看",
        row: {...row},
        api: title == '新增'? addLyrics : title == '编辑' ? editLyrics : undefined,
        getTableList: proTable.value?.getTableList
       }
       //打开抽屉
       lyricDrawer.value?.acceptParams(params);
  }
  //删除数据
  const handelDelete = async (row: Lyrics) => {
    const ids = []
    ids.push(row.id)
    await useHandleData(deleteLyricsByIds, ids, `删除【${row.songName}】歌曲`);
    //从新获取数据
    proTable.value?.getTableList();
  }
  </script>
  <style scoped lang="scss">
  .lyrics-list {
    width: 100%;
    height: 100%;
  }
  </style>
  
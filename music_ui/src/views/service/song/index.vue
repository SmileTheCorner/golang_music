<template>
  <div class="song-list">
    <ProTable
      ref="proTable"
      title="歌曲"
      :columns="columns"
      :requestApi="getSongList"
    >
      <template #tableHeader="scope">
        <el-button type="primary" icon="CirclePlus" @click="openDrawer('新增')"> 新增歌曲 </el-button>
        <el-button type="primary" icon="Download" plain>
          导出用户数据
        </el-button>
        <el-button type="danger" icon="Delete" plain> 批量删除用户 </el-button>
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
    <SongDrawer ref="songDrawer"></SongDrawer>
  </div>
</template>
<script setup lang="tsx" name="SongList">
import {ref,onMounted} from "vue";
import { Song } from "@/type/song/index";
import { getSongList,dropSongById,addSong,editSong } from "@/api/song/index";
import {
  ColumnProps,
  ProTableInstance,
} from "@/components/ProTable/interface/index";
import ProTable from "@/components/ProTable/index.vue";
import { useHandleData } from "@/hooks/useHandleData/index";
import SongDrawer from "@/views/components/songDrawer/index.vue"
import moment from "moment";
import { getDictDataByType } from "@/api/common/index";
import {getDictName} from "@/utils/dict"

const songType = ref<any[]>([])
const newMusic = ref<any[]>([])
const recommendMusic = ref<any[]>([])
const rankingBoard = ref<any[]>([])
onMounted(async () => {
  try {
    const songTypeData = await getDictDataByType({ dictType: "song_type" });
    songType.value.push(...songTypeData);
    const newMusicData = await getDictDataByType({ dictType: "new_music" });
    newMusic.value.push(...newMusicData);
    const recommendMusicData = await getDictDataByType({ dictType: "recommend_music" });
    recommendMusic.value.push(...recommendMusicData);
    const rankingBoardData = await getDictDataByType({ dictType: "ranking_board" });
    rankingBoard.value.push(...rankingBoardData);
  } catch (error) {
    
  }
  
})


// 获取 ProTable 元素，调用其获取刷新数据方法（还能获取到当前查询参数，方便导出携带参数）
const proTable = ref<ProTableInstance>();
//vue3中没有this对象则通过getCurrentInstance来获取当前实列

// 表格配置项
const columns: ColumnProps<any>[] = [
  { type: "selection", fixed: "left", width: 60 },
  { type: "index", label: "序号", width: 80 },
  {
    prop: "songName",
    label: "歌名",
    width: "180",
    search: {
      el: "input",
    },
  },
  {
    prop: "songUrl",
    label: "地址",
    width: "180",
    showOverflowTooltip: true,
  },
  {
    prop: "songCover",
    label: "封面",
    width: "180",
    showOverflowTooltip: true,
  },
  {
    prop: "songSingerId",
    label: "歌手",
  },
  {
    prop: "lyricsId",
    label: "歌词",
  },
  {
    prop: "songDescription",
    label: "描述",
    width: "180",
    showOverflowTooltip: true,
  },
  {
    prop: "songType",
    label: "类型",
    search:{el:"select"},
    enum: songType.value
  },
  {
    prop: "recommend",
    label: "推荐",
    search:{el:"select"},
    width: "100",
    enum: recommendMusic.value,
    render(scope) {
        return (
          <span>
            {scope.row.recommend ? <el-tag type="warning">推荐</el-tag> : <el-tag type="danger">不推荐</el-tag>}
          </span>
        )
    },
  },
  {
    prop: "isNew",
    label: "新音乐",
    search:{el:"select"},
    enum: newMusic.value,
    render(scope) {
        return (
          <span>
            {scope.row.isNew ? <el-tag>是</el-tag> : <el-tag type="success">否</el-tag>}
          </span>
        )
    },
  },
  {
    prop: "rankingListId",
    label: "排行榜",
    render(scope) {
     let name = getDictName(rankingBoard.value,scope.row.rankingListId)
     return (
        <span>
          {name}
        </span>
      )
    },
  },
  {
    prop: "playCount",
    label: "播放次数",
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
const songDrawer = ref<InstanceType<typeof SongDrawer> | null>(null);
const openDrawer = (title:string,row?:Partial<Song>)=>{
     let params = {
      title,
      isView: title === "查看",
      row: {...row},
      api: title == '新增'? addSong : title == '编辑' ? editSong : undefined,
      getTableList: proTable.value?.getTableList
     }
     //打开抽屉
     songDrawer.value?.acceptParams(params);
}
//删除数据
const handelDelete = async (row: Song) => {
  await useHandleData(dropSongById, row.id, `删除【${row.songName}】歌曲`);
  //从新获取数据
  proTable.value?.getTableList();
}
</script>
<style scoped lang="scss">
.song-list {
  width: 100%;
  height: 100%;
}
</style>

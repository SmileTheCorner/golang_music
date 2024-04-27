<template>
  <div class="song-drawer">
    <el-drawer
      v-model="drawerVisible"
      :destroy-on-close="true"
      :title="`${drawerProps.title}歌曲`"
    >
      <el-form
        ref="SongFormRef"
        :model="drawerProps.row"
        :rules="rules"
        :disabled="drawerProps.isView"
        :hide-required-asterisk="drawerProps.isView"
        label-width="auto"
        label-position="left"
      >
        <el-form-item label="歌曲名称" prop="songName">
          <el-input
            v-model="drawerProps.row!.songName"
            placeholder="请输入歌曲名称"
            clearable
          ></el-input>
        </el-form-item>

        <el-form-item label="歌曲描述" prop="songDescription">
          <el-input
            :rows="4"
            v-model="drawerProps.row!.songDescription"
            maxlength="200"
            placeholder="输入歌曲描述"
            show-word-limit
            type="textarea"
          />
        </el-form-item>

        <el-form-item label="歌曲地址" prop="songUrl">
          <UploadFile
            v-model:fileUrl="drawerProps.row.songUrl"
            :api="uploadSong"
          ></UploadFile>
        </el-form-item>

        <el-form-item label="歌曲封面" prop="songCover">
          <UploadImg
            :disabled="!drawerProps.isView"
            v-model:imageUrl="drawerProps.row!.songCover"
            :api="uploadSongCover"
          ></UploadImg>
        </el-form-item>

        <el-form-item label="是否是新音乐" prop="isNew">
          <el-switch v-model="drawerProps.row!.isNew" />
        </el-form-item>

        <el-form-item label="是否推荐" prop="recommend">
          <el-switch v-model="drawerProps.row!.recommend" />
        </el-form-item>

        <el-form-item label="排行版" prop="rankingListId">
          <el-select
            v-model="drawerProps.row!.rankingListId"
            placeholder="请选择排行榜"
            clearable
          >
            <el-option
              v-for="item in rankingList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="歌曲类型" prop="songType">
          <el-select
            v-model="drawerProps.row!.songType"
            placeholder="请选择歌曲类型"
            filterable
          >
            <el-option
              v-for="item in songType"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="歌手" prop="songSingerId">
          <el-select
            v-model="drawerProps.row!.songSingerId"
            filterable
            placeholder="歌手列表"
          >
            <el-option
              v-for="item in singerList"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button
          type="success"
          plain
          v-show="!drawerProps.isView"
          @click="submit(SongFormRef)"
          >确定</el-button
        >
        <el-button type="danger" plain @click="drawerVisible = false"
          >取消</el-button
        >
      </template>
    </el-drawer>
  </div>
</template>
<script setup lang="ts" name="SongDrawer">
import { getDictDataByType } from "@/api/common/index";
import { getSingerList } from "@/api/singer/index";
import { uploadSongCover, uploadSong } from "@/api/song/index";
import UploadImg from "@/components/upload/uploadImg.vue";
import UploadFile from "@/components/upload/uploadFile.vue";
import { Song, SingerList } from "@/type/song/index";
import { ref, reactive, onMounted } from "vue";
import { FormInstance, ElMessage } from "element-plus";

interface DrawerProps {
  title: string; //标题
  isView: boolean; //是否是查看
  row: Partial<Song>; //数据
  api?: (params: any) => Promise<any>; //请求的api
  getTableList?: () => void; //重新获取数据接口
}

//抽屉是否显示
const drawerVisible = ref(false);

//接收抽屉数据
let drawerProps = ref<DrawerProps>({
  isView: false,
  title: "",
  row: {},
});

//接收父组件传过来的参数
const acceptParams = (params: DrawerProps) => {
  drawerVisible.value = true;
  drawerProps.value = params;
};

const SongFormRef = ref<FormInstance>();

//表单校验规则
const rules = reactive({
  songName: [{ required: true, message: "请输入歌曲名称", trigger: "blur" }],
  songUrl: [{ required: true, message: "请输入上传歌曲文件", trigger: "blur" }],
  songCover: [
    { required: true, message: "请输入上传歌曲封面", trigger: "blur" },
  ],
  rankingListId: [{ required: true, message: "请选择排行榜", trigger: "blur" }],
  songType: [{ required: true, message: "请选择歌曲类型", trigger: "blur" }],
  songSingerId: [
    { required: true, message: "请选择歌手信息", trigger: "blur" },
  ],
});

//排行版列表
let rankingList = reactive<any>([]);
//歌手列表
let singerList = reactive<SingerList[]>([]);
//歌曲类型
let songType = reactive<any>([]);

onMounted(async () => {
  try {
    //获取排行榜字典值
    const rankingData = await getDictDataByType({ dictType: "ranking_board" });
    rankingList.push(...rankingData);
    //获取歌曲类型字典值
    const songTypeData = await getDictDataByType({ dictType: "song_type" });
    songType.push(...songTypeData);
    //初始化歌手列表
    const singerData = await getSingerList({});
    for (let key in singerData) {
      let items = singerData[key];
      singerList.push(...items);
    }
  } catch (error) {}
});

//提交数据
const submit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  //等待校验
  await formEl.validate(async (valid: boolean, fields: Record<string, any>) => {
    if (valid) {
      await drawerProps.value.api!(drawerProps.value.row);
      ElMessage.success({
        message: `添加${drawerProps.value.row.songName}歌曲成功！`,
      });
      //从新获取列表数据
      drawerProps.value.getTableList?.();
      //将侧边的抽屉关闭
      drawerVisible.value = false;
    } else {
      console.log("error submit!", fields);
    }
  });
};

//向父组件暴露方法
defineExpose({
  acceptParams,
});
</script>

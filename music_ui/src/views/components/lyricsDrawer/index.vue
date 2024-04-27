<template>
  <div class="lyrics-drawer">
    <el-drawer
      v-model="drawerVisible"
      :destroy-on-close="true"
      :title="`${drawerProps.title}歌词`"
    >
      <el-form
        ref="lyricsFormRef"
        :model="drawerProps.row"
        :rules="rules"
        :disabled="drawerProps.isView"
        :hide-required-asterisk="drawerProps.isView"
        label-width="auto"
        label-position="left"
      >
        <el-form-item label="歌词名称" prop="songName">
          <el-input
            v-model="drawerProps.row!.songName"
            placeholder="请输入歌词名称"
            clearable
          ></el-input>
        </el-form-item>

        <el-form-item label="歌词地址" prop="lyricUrl">
          <UploadFile
            v-model:fileUrl="drawerProps.row.lyricUrl"
            :api="uploadLyrics"
          ></UploadFile>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button
          type="success"
          plain
          v-show="!drawerProps.isView"
          @click="submit(lyricsFormRef)"
          >确定</el-button
        >
        <el-button type="danger" plain @click="drawerVisible = false"
          >取消</el-button
        >
      </template>
    </el-drawer>
  </div>
</template>
<script setup lang="ts" name="LyricsDrawer">
import { Lyrics } from "@/type/lyrics/index";
import { uploadLyrics } from "@/api/lyrics/index";
import UploadFile from "@/components/upload/uploadFile.vue";
import { ref, reactive } from "vue";
import { FormInstance, ElMessage } from "element-plus";

interface DrawerProps {
  title: string; //标题
  isView: boolean; //是否是查看
  row: Partial<Lyrics>; //数据
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

const lyricsFormRef = ref<FormInstance>();

//表单校验规则
const rules = reactive({
  songName: [{ required: true, message: "请输入歌词名称", trigger: "blur" }],
  lyricUrl: [
    { required: true, message: "请输入上传歌词文件", trigger: "blur" },
  ],
});

//提交数据
const submit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  //等待校验
  await formEl.validate(async (valid: boolean, fields: any) => {
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

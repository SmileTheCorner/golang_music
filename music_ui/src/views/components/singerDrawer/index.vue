<template>
  <div class="singer-drawer">
    <el-drawer
      v-model="drawerVisible"
      :destroy-on-close="true"
      :title="`${drawerProps.title}歌手`"
    >
      <el-form
        ref="singerFormRef"
        :model="drawerProps.row"
        :rules="rules"
        :disabled="drawerProps.isView"
        :hide-required-asterisk="drawerProps.isView"
        label-width="auto"
        label-position="left"
      >
        <el-form-item label="歌手名称" prop="singerName">
          <el-input
            v-model="drawerProps.row!.singerName"
            placeholder="请输入歌手名称"
            clearable
          ></el-input>
        </el-form-item>

        <el-form-item label="歌手英文名称">
          <el-input
            v-model="drawerProps.row!.singerEHName"
            placeholder="请输入歌手英文名称"
            clearable
          ></el-input>
        </el-form-item>

        <el-form-item label="歌手头像" prop="singerAvatar">
          <UploadImg
            :disabled="!drawerProps.isView"
            v-model:imageUrl="drawerProps.row!.singerAvatar"
            :api="uploadSongCover"
          ></UploadImg>
        </el-form-item>

        <el-form-item label="歌手分类" prop="singerClassify">
          <el-select
            v-model="drawerProps.row!.singerClassify"
            placeholder="请选择排行榜"
            clearable
          >
            <el-option
              v-for="item in singerClassifyList"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            >
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="歌手描述">
          <el-input
            type="textarea"
            v-model="drawerProps.row!.singerDescription"
            placeholder="请输入歌手描述"
            clearable
          ></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button
          type="success"
          plain
          v-show="!drawerProps.isView"
          @click="submit(singerFormRef)"
          >确定</el-button
        >
        <el-button type="danger" plain @click="drawerVisible = false"
          >取消</el-button
        >
      </template>
    </el-drawer>
  </div>
</template>
<script setup lang="ts" name="SingerDrawer">
import { getDictDataByType } from "@/api/common/index";
import { uploadSongCover, uploadSong } from "@/api/song/index";
import UploadImg from "@/components/upload/uploadImg.vue";
import { Singer } from "@/type/singer/index";
import { ref, reactive, onMounted } from "vue";
import { FormInstance, ElMessage } from "element-plus";
import { async, sync } from "fast-glob";

interface DrawerProps {
  title: string; //标题
  isView: boolean; //是否是查看
  row: Partial<Singer>; //数据
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

const singerFormRef = ref<FormInstance>();

//表单校验规则
const rules = reactive({
  singerName: [{ required: true, message: "请输入歌手名称", trigger: "blur" }],
  singerAvatar: [
    { required: true, message: "请输入上传歌手头像", trigger: "blur" },
  ],
  singerClassify: [
    { required: true, message: "请选择歌曲分类", trigger: "blur" },
  ],
});

//排行版列表
let singerClassifyList = reactive<any>([]);

onMounted(async () => {
  //获取排行榜字典值
  try {
    const data = await getDictDataByType({ dictType: "singer_classify" });
    singerClassifyList.push(...data);
  } catch (error) {}
});

//提交数据
const submit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  //等待校验
  await formEl.validate(async (valid: boolean, fields: any) => {
    if (valid) {
      await drawerProps.value.api!(drawerProps.value.row);
      ElMessage.success({
        message: `添加${drawerProps.value.row.songName}歌手成功！`,
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

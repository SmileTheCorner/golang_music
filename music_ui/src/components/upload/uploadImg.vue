<template>
  <div class="upload-container">
    <el-upload
      :id="uuid"
      action="#"
      class="upload"
      :show-file-list="false"
      :http-request="handleHttpUpload"
      :disabled="!disabled"
    >
      <el-progress
        class="progress"
        type="circle"
        :percentage="progressPercent"
        v-show="isShowProgress"
      />
      <template v-if="imageUrl">
        <img :src="imageUrl" class="upload-image" />
        <div class="img-handle" @click.stop>
          <div class="handle-icon" @click="editImg" v-if="disabled">
            <el-icon><Edit /></el-icon>
            <span>编辑</span>
          </div>
          <div class="handle-icon" @click="imgViewVisible = true">
            <el-icon><ZoomIn /></el-icon>
            <span>查看</span>
          </div>
          <div class="handle-icon" @click="dropImg" v-if="disabled">
            <el-icon><Delete /></el-icon>
            <span>删除</span>
          </div>
        </div>
      </template>
      <el-icon v-if="!imageUrl && !isShowProgress" class="uploader-icon"
        ><Plus
      /></el-icon>
    </el-upload>
    <el-image-viewer
      v-if="imgViewVisible"
      :url-list="[imageUrl]"
      @close="imgViewVisible = false"
    />
  </div>
</template>

<script setup lang="ts" name="UploadImg">
import { AxiosProgressEvent } from "axios";
import { dropFileByFileName } from "@/api/common/index";

import { ref, getCurrentInstance } from "vue";

const currentInstance = getCurrentInstance();

interface UploadFileProps {
  imageUrl: string; // 图片地址 ==> 必传
  api?: (
    params: any,
    progressCallback: (progressEvent: AxiosProgressEvent) => void
  ) => Promise<any>; // 上传图片的 api 方法，一般项目上传都是同一个 api 方法，在组件里直接引入即可 ==> 非必传
  drag?: boolean; // 是否支持拖拽上传 ==> 非必传（默认为 true）
  disabled?: boolean; // 是否禁用上传组件 ==> 非必传（默认为 false）
  fileSize?: number; // 图片大小限制 ==> 非必传（默认为 5M）
  fileType?: File.ImageMimeType[]; // 图片类型限制 ==> 非必传（默认为 ["image/jpeg", "image/png", "image/gif"]）
  height?: string; // 组件高度 ==> 非必传（默认为 150px）
  width?: string; // 组件宽度 ==> 非必传（默认为 150px）
  borderRadius?: string; // 组件边框圆角 ==> 非必传（默认为 8px）
}
// 接受父组件参数
const props = withDefaults(defineProps<UploadFileProps>(), {
  imageUrl: "",
  drag: true,
  disabled: false,
  fileSize: 5,
  fileType: () => ["image/jpeg", "image/png", "image/gif", "image/jpg"],
  height: "150px",
  width: "150px",
  borderRadius: "8px",
});

// 生成组件唯一id
const uuid = ref("id-" + crypto.randomUUID());

const imgViewVisible = ref(false);

//进度条默认为0
const progressPercent = ref<number>(0);

//是否显示进度条
const isShowProgress = ref<boolean>(false);

/**
 * @description 图片上传
 * @param options upload 所有配置项
 * */
interface UploadEmits {
  (e: "update:imageUrl", value: string): void;
}
const emit = defineEmits<UploadEmits>();

//上传文件
const handleHttpUpload = async (params: any) => {
  //上传之前将进度条恢复成默认值
  progressPercent.value = 0;
  //显示进度条
  isShowProgress.value = true;

  //获取到上传的文件
  let file = params.file;
  let formData = new FormData();
  formData.append("file", file);

  //进度条
  const uploadProgressEvent = (progressEvent: AxiosProgressEvent) => {
    if (progressEvent.total) {
      progressPercent.value = Math.floor(
        (progressEvent.loaded / progressEvent.total) * 100
      );
    }
  };

  try {
    if (props.api) {
      const { data } = await props.api(formData, uploadProgressEvent);
      //上传成功向父组件提交上传的图片地址
      emit("update:imageUrl", data);
      isShowProgress.value = false;
    }
  } catch (error) {
    console.log("error", error);
  }
};

//编辑图片
const editImg = () => {
  const dom = document.querySelector(`#${uuid.value} .el-upload__input`);
  dom && dom.dispatchEvent(new MouseEvent("click"));
};

//删除图片
const dropImg = () => {
  //删除服务器上的照片
  let arr = props.imageUrl.split("/");
  let relativePath = `/${arr[arr.length - 2]}/${arr[arr.length - 1]}}`;
  try {
    dropFileByFileName({ fileName: relativePath });
    emit("update:imageUrl", "");
    currentInstance?.appContext.config.globalProperties.$message.success(
      "删除成功"
    );
  } catch (error) {
    console.log("error", error);
  }
};
</script>

<style lang="scss" scoped>
.upload-container {
  .upload {
    border: 2px dashed var(--el-border-color-darker);
    width: v-bind(width);
    height: v-bind(height);
    border-radius: v-bind(borderRadius);
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    &:hover {
      border-color: var(--el-color-primary);
      .uploader-icon {
        color: var(--el-color-primary);
      }
      .img-handle {
        opacity: 1;
      }
    }
    .progress {
      position: absolute;
      left: 50%;
      top: 50%;
      transform: translate(-50%, -50%);
      z-index: 10;
    }
    .uploader-icon {
      color: var(--el-border-color-darker);
      font-size: 1.5em;
    }
    .upload-image {
      width: 100%;
      height: 100%;
      object-fit: contain;
    }
    .img-handle {
      height: 100%;
      width: 100%;
      position: absolute;
      display: flex;
      justify-content: space-around;
      align-items: center;
      box-sizing: border-box;
      background: rgb(0 0 0 / 60%);
      opacity: 0;
      cursor: pointer;
      transition: var(--el-transition-duration-fast);
      .handle-icon {
        display: flex;
        justify-content: center;
        align-items: center;
        color: #fff;
      }
    }
  }
}
</style>

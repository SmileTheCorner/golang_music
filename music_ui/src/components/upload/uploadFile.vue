<template>
  <div class="upload-container">
    <el-upload
      class="upload"
      :drag="drag"
      action="#"
      :show-file-list="false"
      :http-request="handleHttpUpload"
    >
      <div class="info">
        <div class="icon">
          <slot name="icon">
            <el-icon><UploadFilled /></el-icon>
          </slot>
        </div>
        <div class="el-upload__text">点击或<em>拖拽</em></div>
      </div>
      <div class="progress" v-if="isShowProgress">
        <el-progress :percentage="progressPercent"/>
      </div>
    </el-upload>
    <div class="file-name">{{ fileName }}</div>
  </div>
</template>

<script setup lang="ts">
import {ref,computed} from "vue"
import { AxiosProgressEvent } from "axios";
interface UploadFileProps {
  fileUrl: string; // 图片地址 ==> 必传
  api?: (
    params: any,
    progressCallback: (progressEvent: AxiosProgressEvent) => void
  ) => Promise<any>; // 上传文件的 api 方法，一般项目上传都是同一个 api 方法，在组件里直接引入即可 ==> 非必传
  drag?: boolean; // 是否支持拖拽上传 ==> 非必传（默认为 true）
  disabled?: boolean; // 是否禁用上传组件 ==> 非必传（默认为 false）
  fileSize?: number; // 文件大小限制 ==> 非必传（默认为 5M）
  fileType?: string; // 文件类型限制 ==> 非必传（默认为）
  height?: string; // 组件高度 ==> 非必传（默认为 150px）
  width?: string; // 组件宽度 ==> 非必传（默认为 150px）
  borderRadius?: string; // 组件边框圆角 ==> 非必传（默认为 8px）
}

/**
 * @description 图片上传
 * @param options upload 所有配置项
 * */
 interface UploadEmits {
  (e: "update:fileUrl", value: string): void;
}
const emit = defineEmits<UploadEmits>();

//接受父组件参数
const props = withDefaults(defineProps<UploadFileProps>(), {
  fileUrl: "",
  drag: true,
  disabled: false,
  fileSize: 5,
  height: "150px",
  width: "150px",
  borderRadius: "8px",
});

//进度百分比
const progressPercent = ref<number>(40)
//是否显示进度百分比
const isShowProgress = ref<boolean>(false)
//上传成功后回显的文件名称
//const fileName = ref<string>("")
const fileName = computed(()=>{
    const filename = props.fileUrl.replace(/^.*[\\\/]/, '');
    return filename
})

//自定义上传的接口
const handleHttpUpload = async (param:any)=>{
  //获取到要上传的文件
  let file = param.file
  //构建 formData对象
  let formData = new FormData()
  //将要上传的文件添加进去
  formData.append("file",file)

  //上传文件之前先讲进度条重置为 0
  progressPercent.value = 0
  //同时将进度条进行显示
  isShowProgress.value = true

  //计算上传的进度
  const uploadProgress = (progressEvent:AxiosProgressEvent)=>{
      //如果有文件上传则进行计算
      if(progressEvent.total){
         progressPercent.value = Math.floor(progressEvent.loaded/progressEvent.total)*100
      }
  }

  //调用后台真正的上传接口
  try {
    if(props.api){
      const { data } = await props.api(formData, uploadProgress);
      //获取上传成功后的文件名
      let name  = param.file.name
      fileName.value = name
      //上传成功向父组件提交上传的图片地址
       emit("update:fileUrl", data);
    }
  } catch (error) {
    console.log("error",error)
  }
}
</script>

<style lang="scss" scoped>
.upload-container {
  width: v-bind(width);
  .upload {
    box-sizing: border-box;
    border: 2px dashed var(--el-border-color-darker);
    border-radius: v-bind(borderRadius);
    width: v-bind(width);
    height: v-bind(height);
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    position: relative;
    &:hover {
      border-color: var(--el-color-primary);
      .icon {
        color: var(--el-color-primary);
      }
    }
    :deep(.el-upload) {
      width: 100%;
      height: 100%;
    }
    :deep(.el-upload-dragger) {
      padding: 0;
      height: 100%;
      display: flex;
      flex-direction: column;
      align-items: center;
    }
    .info {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      flex: 1;
      .icon {
        font-size: 3rem;
        color: var(--el-border-color-darker);
      }
    }

    .progress {
      width: 90%;
      height: 15%;
      :deep(.el-progress__text) {
        font-size: 10px !important;
        min-width: auto !important;
      }
    }
  }
  .file-name {
    color: var(--el-border-color-darker);
    width: 100%;
    overflow: hidden; //超出的隐藏
    text-overflow: ellipsis; //显示省略符号来代表被修剪的文本。
    white-space: nowrap; //不换行
  }
}
</style>

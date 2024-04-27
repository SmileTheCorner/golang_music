import axios from "@/request/index";

//根据字典类型获取字典值
export const getDictDataByType=(params:{dictType:string})=>{
   return axios.request({
    url:"/api/v1/dictType/getByDictType",
    method:"get",
    params
   })
}

//根据文件名删除文件
export const dropFileByFileName = (params:{fileName:string})=>{
   return axios.request({
      url:"/api/v1/song/dropUploadedFile",
      method:"get",
      params
   })
}
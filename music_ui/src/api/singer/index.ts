import axios from "@/request/index";
import {Singer} from "@/type/singer/index"

//歌手列表
export const getSingerList = (data:any)=>{
    return axios.request({
        url:"/api/v1/singer/queryList",
        method:"post",
        data
    })
}
//新增歌手信息
export const addSinger = (params: Singer)=>{
    return axios.request({
        url:"/api/v1/singer/add",
        method:"post",
        data:params
    })
}
//编辑歌手
export const editSinger = (params: Singer)=>{
    return axios.request({
        url:"/api/v1/singer/edit",
        method:"post",
        data:params
    })
}
//根据id数组删除歌手
export const deleteSingerByIds = (ids: string[])=>{
    return axios.request({
        url:"/api/v1/singer/dropByIds",
        method:"delete",
        data:ids
    })
}
//根据歌手名查询歌手信息
export const querySingerByName = (name: string)=>{
    return axios.request({
        url:`/api/v1/singer/querySingerByName/${name}`,
        method:"get",
    })
}
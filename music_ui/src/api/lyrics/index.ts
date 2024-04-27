import axios from "@/request/index";

//分页获取歌词列表
export const getLyricsList = (params: any)=>{
    return axios.request({
        url:"/api/v1/lyric/list",
        method:"post",
        data: params,
    })
}

//添加歌词
export const addLyrics = (data:any)=>{
    return axios.request({
        url:"/api/v1/lyric/add",
        method:"post",
        data
    })
}

//上传歌词到对象存储中
export const uploadLyrics = (data:any)=>{
    return axios.request({
        url:"/api/v1/lyric/upload",
        method:"post",
        data
    })
}

//根据id数组删除歌词
export const deleteLyricsByIds = (ids:string[])=>{
    return axios.request({
        url:"/api/v1/lyric/drop",
        method:"delete",
        data:ids
    })
}

//编辑歌词
export const editLyrics = (data:any)=>{
    return axios.request({
        url:"/api/v1/lyric/edit",
        method:"put",
        data
    })
}



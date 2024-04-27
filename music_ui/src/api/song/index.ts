import axios from "@/request/index";
import { Song } from "@/type/song/index";
import {AxiosProgressEvent} from "axios"

//获取音乐列表数据
export const getSongList = (params: any) => {
  return axios.request({
    url: "/api/v1/song/pageList",
    method: "post",
    data: params,
  });
};

//根据Id删除歌曲
export const dropSongById = (songId: string) => {
  return axios.request({
    url: "/api/v1/song/dropSongById",
    method: "get",
    params: { songId },
  });
};

//新增歌曲
export const addSong = (row: Song) => {
  return axios.request({
    url: "/api/v1/song/addSong ",
    method: "post",
    data: row,
  });
};

//编辑歌曲
export const editSong = (row: Song) => {
  return axios.request({
    url: "/api/v1/song/update",
    method: "post",
    data: row,
  });
};

//上传歌曲封面
export const uploadSongCover = (params: any, onUploadProgress: (progressEvent: AxiosProgressEvent) => void) => {
  return axios.request({
    url: "/api/v1/song/uploadSongCover",
    method: "post",
    onUploadProgress: onUploadProgress,
    data: params,
  });
};

//上传歌曲
export const uploadSong = (params: any,onUploadProgress: (progressEvent: AxiosProgressEvent) => void) => {
  return axios.request({
    url: "/api/v1/song/uploadSong",
    method: "post",
    onUploadProgress: onUploadProgress,
    data: params,
  });
};

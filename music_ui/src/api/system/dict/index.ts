import axios from "@/request/index";
import { Dict, DictValue } from "@/type/dict";
//获取所有的字典类型
export const getDictTypeList = () => {
    return axios.request({
        url: "/api/v1/dictType/getAllType",
        method: "get",
    })
}
//根据字典类型获取字典值
export const getDictDataByType=(params:{dictType:string})=>{
    return axios.request({
     url:"/api/v1/dictType/getByDictType",
     method:"get",
     params
    })
 }

//获取字典类型树
export const getDictTypeTree = () => {
    return axios.request({
        url: "/api/v1/dictType/dictTypeTree",
        method: "get",
    })
}

//根据id获取字典类型信息
export const getDictTypeById = (id:string) => {
    return axios.request({
        url: "/api/v1/dictType/getById",
        method: "get",
        params:{id}
    })
}

//新增字典类型
export const addDictType = (params:Dict) => {
    return axios.request({
        url: "/api/v1/dictType/add",
        method: "post",
        data:params
    })
}

//修改字典类型
export const updateDictType = (params:Dict) => {
    return axios.request({
        url: "/api/v1/dictType/update",
        method: "put",
        data:params
    })
}

//批量删除字典类型
export const dropDictTypeByIds = (ids:string[]) => {
    return axios.request({
        url: "/api/v1/dictType/deleteByIds",
        method: "delete",
        data:ids
    })
}

//修改字典值
export const updateDictValueData = (params:DictValue) => {
    return axios.request({
        url: "/api/v1/dictData/update",
        method: "put",
        data:params
    })
}

//新增字典值
export const addDictValueData = (params:DictValue) => {
    return axios.request({
        url: "/api/v1/dictData/add",
        method: "post",
        data:params
    })
}

//根据 ids 数组删除字典值
export const dropDictValueByIds = (ids:string[]) => {
    return axios.request({
        url: "/api/v1/dictData/dropByIds",
        method: "delete",
        data:ids
    })
}
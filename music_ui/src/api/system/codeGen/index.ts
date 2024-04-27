import axios from "@/request/index";
import { CodeInfo } from "@/type/codeGeneration/index";


//获取所有表信息
export const getAllTableList = (param: any) => {
  return axios.request({
    url: "/api/v1/codeGen/allTables",
    method: "post",
    data: param
  });
};

// 导入数据表到生成的业务表中
export const importGenTable = (param: any) => {
  return axios.request({
    url: "/api/v1/codeGen/importGenTable",
    method: "post",
    data: param
  });
};

//分页查询代码生成业务表
export const getGenTableList = (param: any) => {
  return axios.request({
    url: "/api/v1/codeGen/genTableList",
    method: "post",
    data: param
  });
};

//根据表名查询表信息
export const getColumnsByName = (tableName: string) => {
  return axios.request({
    url: "/api/v1/codeGen/allColumnsByTableName",
    method: "get",
    params: { tableName }
  });
};

//一键生成代码
export const oneKeyGenerateCode = (param: CodeInfo) => {
  return axios.request({
    url: "/api/v1/codeGen/oneKeyGenerateCode",
    method: "post",
    data: param
  });
};

//根据ID切片删除代码生成业务表
export const dropGenTableById = (id: string[]) => {
  return axios.request({
    url: "/api/v1/codeGen/dropGenTableByIds",
    method: "delete",
    data: id
  });
};

//下载代码
export const downloadCode = (id: string[]) => {
  return axios.download("/api/v1/codeGen/downloadCode", id);
};

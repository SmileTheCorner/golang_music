import axios from "@/request/index";
import { Department } from "@/type/department/index";

//获取部门列表
export const getDepartmentList = (params: Department) => {
  return axios.request({
    url: "/api/v1/department/list",
    method: "post",
    data: params,
  });
};
//添加部门信息
export const addDepartment = (params: Department) => {
  return axios.request({
    url: "/api/v1/department/add",
    method: "post",
    data: params,
  });
};
//删除部门信息
export const dropDepartment = (params: Department) => {
  return axios.request({
    url: "/api/v1/department/drop",
    method: "post",
    data: params,
  });
};
//修改部门信息
export const updateDepartment = (params: Department) => {
  return axios.request({
    url: "/api/v1/department/update",
    method: "post",
    data: params,
  });
};
//获取部门树
export const departmentTree = () => {
  return axios.request({
    url: "/api/v1/department/tree",
    method: "post",
  });
};

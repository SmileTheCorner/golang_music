import axios from "@/request/index";
import { Role } from "@/type/role/index"

//获取角色列表
export const getRoleList = (params: Role) => {
  return axios.request({
    url: "/api/v1/role/list",
    method: "post",
    data: params,
  });
};
//添加角色信息
export const addRole = (params: Role) => {
  return axios.request({
    url: "/api/v1/role/add",
    method: "post",
    data: params,
  });
};
//删除角色信息
export const dropRole = (id: string) => {
  return axios.request({
    url: "/api/v1/role/drop",
    method: "post",
    data: id,
  });
};
//修改角色信息
export const updateRole = (params: Role) => {
  return axios.request({
    url: "/api/v1/role/update",
    method: "post",
    data: params,
  });
};
//根据id获取角色信息
export const getRoleById = (id: string) => {
  return axios.request({
    url: "/api/v1/role/infoById",
    method: "post",
    data: { id }
  });
};
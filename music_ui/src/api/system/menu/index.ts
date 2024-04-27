import axios from "@/request/index";

// * 获取按钮权限
export const getAuthButtonListApi = () => {
	let AuthButtons ={
        "code": 200,
        "data": {
            "useProTable": ["add", "batchAdd", "export", "batchDelete", "status"],
            "authButton": ["add", "edit", "delete", "import", "export"]
        },
        "msg": "成功"
    }
	return AuthButtons;
};

// * 获取菜单列表
export const getAuthMenuListApi = () => {
    return axios.request({
     url:"/api/v1/menu/list",
     method:"post"
    })
};

//菜单列表
export const listMenu = (params:any) => {
    return axios.request({
        url:"/api/v1/menu/query",
        method:"post",
        data:params
       })
};

//菜单列表
export const addMenuTree = () => {
    return axios.request({
        url:"/api/v1/menu/add/tree",
        method:"get",
       })
};

//新增菜单
export const addMenu = (params:Menu.MenuOptions) => {
    return axios.request({
        url:"/api/v1/menu/add",
        method:"post",
        data:params
       })
};

//删除菜单
export const delMenu = (params:{id:string}) => {
    return axios.request({
        url:"/api/v1/menu/drop",
        method:"get",
        params
       })
};

//修改菜单
export const editMenu = (params:Menu.MenuOptions) => {
    return axios.request({
        url:"/api/v1/menu/update",
        method:"post",
        data:params
       })
};


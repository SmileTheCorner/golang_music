import axios from "@/request/index";
import { Login } from "@/type/user";

export function doLogin(params: Login.ReqLoginForm) {
    return axios.request({
        url: '/api/v1/public/user/login',
        method: 'post',
        data: params
    })
}

export function list(params: any) {
    return axios.request({
        url: '/api/v1/user/list',
        method: 'post',
        data: params
    })
}

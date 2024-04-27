import axios, { AxiosInstance, AxiosRequestConfig } from "axios"
import { ResultData } from "@/type/common/index";
import { AxiosResponse } from "axios";
import { useRouter } from "vue-router";
import { ElMessage } from 'element-plus'

const config = {
    //默认请求路径
    baseUrl: import.meta.env.VITE_API_URL,
    //请求超时时间
    timeOut: 10000,
    //跨域时是否携带凭证
    withCredentials: true
}

//定义请求的类
class RequestHttp {
    //定义axios的实列对象属性
    instance: AxiosInstance

    //初始化构造器
    public constructor(config: AxiosRequestConfig) {
        // 实例化axios
        this.instance = axios.create(config)

        /**
        * @description 请求拦截器
        * 客户端发送请求 -> [请求拦截器] -> 服务器
        * token校验(JWT) : 接受服务器返回的token,存储到pinia/本地储存当中
        */
        this.instance.interceptors.request.use(config => {
            const token = JSON.parse(localStorage.getItem('GlobalState') || '{}').token
            config.headers.Authorization = "Bearer " + token;
            return config
        }, error => {
            return Promise.reject(error)
        })

        /**
         * @description 响应拦截器
         *  服务器换返回信息 -> [拦截统一处理] -> 客户端JS获取到信息
         */
        this.instance.interceptors.response.use((response: AxiosResponse) => {
            const { data, code, msg } = response.data as ResultData
            if (code === 200) {
                return data
            } else if (code === 403) {
                //token过期
                const router = useRouter();
                //过期,清除token
                localStorage.removeItem('GlobalState')
                //跳转到登录页面
                router.push('/login')
                console.log("token过期", msg)
            } else {
                //提示信息
                ElMessage.error(msg)
            }
        }, error => {
            ElMessage.error(error)
            return Promise.reject(error)
        })

    }

    //定义请求方法
    request(axiosConfig: AxiosRequestConfig): Promise<any> {
        return this.instance.request(axiosConfig)
    }
    //定义下载请求方法
    download(url: string, data: any): Promise<any> {
        return this.instance.post(url, data, { responseType: 'blob' }).then(res => {
            console.log("下载的数据===>", res)
            // 创建一个Blob对象，并用它创建一个URL
            const blob = new Blob([res]);
            const url = window.URL.createObjectURL(blob);

            // 创建一个a标签，设置其href为Blob URL，触发点击事件进行下载
            const link = document.createElement("a");
            link.href = url;
            link.download = "downloadedFile.zip";
            link.click();

            // 释放Blob URL资源
            window.URL.revokeObjectURL(url);
        })
    }
}

export default new RequestHttp(config)


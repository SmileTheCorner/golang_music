// * 请求响应参数(不包含data)
export interface Result {
	code: number;
	msg: string;
}

// * 请求响应参数(包含data)
export interface ResultData<T = any> extends Result {
	data: T;
}

/* UserState */
export interface UserState {
	token: string;
	userInfo: { name: string };
  }

 
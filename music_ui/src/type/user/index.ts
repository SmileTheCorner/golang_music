// * 登录模块
export namespace Login {
	export interface ReqLoginForm {
		username: string;
		password: string;
	}
	export interface ResLogin {
		access_token: string;
	}
}

//用户信息
export interface User {
	ID?:number,
	CreatedAt?:string,
	UpdatedAt?:string,
	name?:string,
	realName?:string,
	avatar?:string,
	mobile?:string,
	email?:string,
	password?:string,
	gender?:number,
	idCard?:string,
	address?:string,
	status?:number,
	age?:number,
}
export interface Role {
    id?:string,
    name?:string,
    code?:string,
    sort?:number,
    menuList?:string[],
    dataScope?:number,
    dataScopeDept?:string[],
    status?:number,
    remark?:string,
    createdAt?:string,
    updatedAt?:string,
}
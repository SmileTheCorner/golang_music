export interface Department {
    id?:string,
    createdAt?:string,
    parentId?:string,
    name?:string,
    contacts?:string,
    phone?:string,
    email?:string,
    orderNum?:number,
    status?:number,
    children?:Department[],
}
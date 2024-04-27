import type { Column } from 'element-plus'

export interface Dict {
    id?: string,
    createBy?: string,
    createdAt?: string,
    updateBy?: string,
    updatedAt?: string,
    deletedAt?: string,
    parentId?: string,
    dictName?: string,
    dictType?: string,
    status?: number,
    remark?: string,
    children?: Dict[]
}

export interface DictValue {
    id?: string
    createdAt?: string
    sort?: number
    label?: string
    value?: string
    dictType?: string
    status?: string
    remark?: string,
    checked?: boolean;
}

export type CellRenderProps<T> = {
    cellData: T
    column: Column<T>
    columns: Column<T>[]
    columnIndex: number
    rowData: any
    rowIndex: number
}
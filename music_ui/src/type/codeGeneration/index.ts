export interface CodeGen {
    tableSchema?:string,
    tableName?:string,
    createTime?:string,
    updateTime?:string,
    tableComment?:string,
}

export interface CodeInfo {
    moduleTitle?:string,
    author?:string,
    date?:string,
    moduleName?:string,
    fields?:ColumnInfo[],   
}

export interface ColumnInfo {
    tableSchema?:string,
    tableName?:string,
    columnName?:string,
    columnComment?:string,
    dataType?:string,
    columnType?:string,
    characterMaximumLength?:string,
    characterOctetLength?:string,
    columnKey?:string,
}

export interface GenTable{
    id?:string,
    tableName?:string,
    tableComment?:string,
    subTableName?:string,
    subTableFkName?:string,
    structName?:string,
    templateUse?:string,
    packageName?:string,
    authorName?:string,
    genType?:string,
    genPath?:string,
    options?:string,
    remark?:string,
    createBy?:string,
    createdAt?:string,
    updateBy?:string,
    updatedAt?:string,
    deletedAt?:string,
}

export interface GenTableColumn{
    id?:string,
    tableId?:string,
    columnName?:string,
    columnComment?:string,
    columnType?:string,
    goType?:string,
    goProperty?:string,
    isPk?:string,
    isIncrement?:string,
    isRequired?:string,
    isQuery?:string,
    queryMethod?:string,
    htmlType?:string,
    dictType?:string,
    sort?:number,
}
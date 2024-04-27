//表单类型
type FormType = 
"input" 
| "select" 
| "checkbox" 
| "radio" 
| "textarea" 
| "password"
| "datepicker"
| "timepicker"
| "switch"

type DateType = 
'year' 
| 'month' 
| 'date' 
| 'dates' 
| 'datetime' 
| 'week' 
| 'datetimerange' 
| 'daterange' 
| 'monthrange'

type DateTimeType =
'year' 
| 'month'
| 'date'
| 'datetime'
| 'week'
| 'datetimerange'
| 'daterange'

//下拉选择框类型
interface ItemOption {
    label: string
    value: string | number | any
}

//表单项配置
export interface FormItem {
    type: FormType //输入框类型
    label: string //输入框标题
    prop: string //字段名
    value?: any //值
    labelWidth?: string | number //标签的宽度
    placeholder?: string //输入框默认显示内容
    options?: ItemOption[] //选择器的可选子选项 select
    dateTimeType?: DateType | DateTimeType //日期选择器的类型
    isRange?: boolean //是否是时间范围选择器
}

export interface FormOption {
    formItems: FormItem[]
    labelWidth?: string | number //标签的长度
    labelPosition?: 'left' | 'right' | 'top' //标签位置
    size?: 'large' | 'small' | 'default' //表单内组件大小
    inline?:boolean //是否是行内表单
    disabled?:boolean //是否禁用表单
    callBackFunc?: (data: any) => void //回调函数
}
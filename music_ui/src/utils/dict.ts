/**
 * @description：传入字典对象和字典值，返回字典值对应的字典名称
 */
export function getDictName(dict: any[], value: any) {
    for (var i = 0; i < dict.length; i++) {
        if (dict[i].value == value) {
            return dict[i].label;
        }
    }
    return "";
}

export function getDictTreeName(dict: any[], value: any) {
    if (value == null || value === "") {
        return "根目录"
    }
    for (let i = 0; i < dict.length; i++) {
        if (dict[i].id === value) {
            return dict[i].dictName;
        }
        const result: any = getDictTreeName(dict[i].children, value);
        if (result) {
            return result;
        }
    }
    return null;
}

//下拉树选择时自己不能作为自己的父亲
export function notSelectSelfAsParent(dict: any[], value: any) {
    for (let i = 0; i < dict.length; i++) {
        if (dict[i].id === value) {
            dict[i].disabled = true
        } else {
            dict[i].disabled = false
        }
        if (dict[i].children && dict[i].children.length > 0) {
            notSelectSelfAsParent(dict[i].children, value);
        }
    }
}


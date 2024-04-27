import { FieldNamesProps } from "@/components/ProTable/interface/index"


/**
 * @description 交换对象中的属性值
 * @param obj1 {Object}
 * @param obj2 {Object}
 * @param property string 属性名
 */
export function swapProperty(obj1: any, obj2: any, property: string) {
  const temp = obj1[property];
  obj1[property] = obj2[property];
  obj2[property] = temp;
}
/**
 * @description 判断修改节点和选中节点是否在树中
 *  @param {Object} editNode 选中节点
 * @param {Object} selectedNode 选中节点
 * @param {Array} tree 树
 */
export function isNodeInDomTree(editNode: any, selectedNode: any, tree: any) {
  for (let i = 0; i < tree.length; i++) {
    let node = tree[i];
    let editIsInTree = isNodeInTree(node,editNode)
    let selectedIsInTree = isNodeInTree(node,selectedNode)
    if (editIsInTree && selectedIsInTree) {
      return true;
    }
  }
  return false;
}
/**
 * @description 判断节点是否在一颗树中
 * @param {Array} tree 树 
 * @param {Object} targetNode 目标节点
 */
function isNodeInTree(tree: any, targetNode: any) {
  if (tree === targetNode) {
    return true;
  }
  if (tree.children) {
    for (let i = 0; i < tree.children.length; i++) {
      if (isNodeInTree(tree.children[i], targetNode)) {
        return true;
      }
    }
  }
  return false;
}
/**
 * @description 获取localStorage
 * @param {String} key Storage名称
 * @return string
 */
export function localGet(key: string) {
  const value = window.localStorage.getItem(key);
  try {
    return JSON.parse(window.localStorage.getItem(key) as string);
  } catch (error) {
    return value;
  }
}

/**
 * @description 存储localStorage
 * @param {String} key Storage名称
 * @param {Any} value Storage值
 * @return void
 */
export function localSet(key: string, value: any) {
  window.localStorage.setItem(key, JSON.stringify(value));
}

/**
 * @description 清除localStorage
 * @param {String} key Storage名称
 * @return void
 */
export function localRemove(key: string) {
  window.localStorage.removeItem(key);
}

/**
 * @description 清除所有localStorage
 * @return void
 */
export function localClear() {
  window.localStorage.clear();
}

/**
 * @description 获取当前时间对应的提示语
 * @return string
 */
export function getTimeState() {
  // 获取当前时间
  let timeNow = new Date();
  // 获取当前小时
  let hours = timeNow.getHours();
  // 判断当前时间段
  if (hours >= 6 && hours <= 10) return `早上好 ⛅`;
  if (hours >= 10 && hours <= 14) return `中午好 🌞`;
  if (hours >= 14 && hours <= 18) return `下午好 🌞`;
  if (hours >= 18 && hours <= 24) return `晚上好 🌛`;
  if (hours >= 0 && hours <= 6) return `凌晨好 🌛`;
}

/**
 * @description 递归查询当前路由所对应的路由
 * @param {Array} menuList 所有菜单列表
 * @param {String} path 当前访问地址
 * @return array
 */
export function filterCurrentRoute(menuList: Menu.MenuOptions[], path: string) {
  let result = {};
  for (let item of menuList) {
    if (item.path === path) return item;
    if (item.children) {
      const res = filterCurrentRoute(item.children, path);
      if (Object.keys(res).length) result = res;
    }
  }
  return result;
}

/**
 * @description 扁平化数组对象(主要用来处理路由菜单)
 * @param {Array} menuList 所有菜单列表
 * @return array
 */
export function getFlatArr(menuList: Menu.MenuOptions[]) {
  let newMenuList: Menu.MenuOptions[] = JSON.parse(JSON.stringify(menuList));
  return newMenuList.reduce(
    (pre: Menu.MenuOptions[], current: Menu.MenuOptions) => {
      let flatArr = [...pre, current];
      if (current.children)
        flatArr = [...flatArr, ...getFlatArr(current.children)];
      return flatArr;
    },
    []
  );
}

/**
 * @description 使用递归，过滤需要缓存的路由（暂时没有使用）
 * @param {Array} menuList 所有菜单列表
 * @param {Array} cacheArr 缓存的路由菜单 name ['**','**']
 * @return array
 * */
export function getKeepAliveRouterName(
  menuList: Menu.MenuOptions[],
  keepAliveArr: string[] = []
) {
  menuList.forEach((item) => {
    item.meta.isCache && item.name && keepAliveArr.push(item.name);
    item.children?.length &&
      getKeepAliveRouterName(item.children, keepAliveArr);
  });
  return keepAliveArr;
}

/**
 * @description 使用递归，过滤出需要渲染在左侧菜单的列表（剔除 visible == true 的菜单）
 * @param {Array} menuList 所有菜单列表
 * @return array
 * */
export function getShowMenuList(menuList: Menu.MenuOptions[]) {
  let newMenuList: Menu.MenuOptions[] = JSON.parse(JSON.stringify(menuList));
  return newMenuList.filter((item) => {
    item.children?.length && (item.children = getShowMenuList(item.children));
    return !item.meta?.visible;
  });
}

/**
 * @description 使用递归处理路由菜单 path，生成一维数组(第一版本地路由鉴权会用到)
 * @param {Array} menuList 所有菜单列表
 * @param {Array} menuPathArr 菜单地址的一维数组 ['**','**']
 * @return array
 */
export function getMenuListPath(
  menuList: Menu.MenuOptions[],
  menuPathArr: string[] = []
) {
  menuList.forEach((item: Menu.MenuOptions) => {
    typeof item === "object" && item.path && menuPathArr.push(item.path);
    item.children?.length && getMenuListPath(item.children, menuPathArr);
  });
  return menuPathArr;
}

/**
 * @description 递归找出所有面包屑存储到 pinia 中
 * @param {Array} menuList 所有菜单列表
 * @param {Object} result 输出的结果
 * @param {Array} parent 父级菜单
 * @returns object
 */
export const getAllBreadcrumbList = (
  menuList: Menu.MenuOptions[],
  result: { [key: string]: any } = {},
  parent = []
) => {
  for (const item of menuList) {
    result[item.path] = [...parent, item];
    if (item.children)
      getAllBreadcrumbList(item.children, result, result[item.path]);
  }
  return result;
};

/**
 * @description 处理 prop，当 prop 为多级嵌套时 ==> 返回最后一级 prop
 * @param {String} prop 当前 prop
 * @returns {String}
 * */
export function handleProp(prop: string) {
  const propArr = prop.split(".");
  if (propArr.length == 1) return prop;
  return propArr[propArr.length - 1];
}

/**
 * @description 处理 prop 为多级嵌套的情况，返回的数据 (列如: prop: user.name)
 * @param {Object} row 当前行数据
 * @param {String} prop 当前 prop
 * @returns {*}
 * */
export function handleRowAccordingToProp(row: { [key: string]: any }, prop: string) {
  if (!prop.includes(".")) return row[prop] ?? "--";
  prop.split(".").forEach(item => (row = row[item] ?? "--"));
  return row;
}

/**
 * @description 根据枚举列表查询当需要的数据（如果指定了 label 和 value 的 key值，会自动识别格式化）
 * @param {String} callValue 当前单元格值
 * @param {Array} enumData 字典列表
 * @param {Array} fieldNames label && value && children 的 key 值
 * @param {String} type 过滤类型（目前只有 tag）
 * @returns {String}
 * */
export function filterEnum(
  callValue: any,
  enumData?: any,
  fieldNames?: FieldNamesProps,
  type?: "tag"
) {
  const value = fieldNames?.value ?? "value";
  const label = fieldNames?.label ?? "label";
  const children = fieldNames?.children ?? "children";
  let filterData: { [key: string]: any } = {};
  // 判断 enumData 是否为数组
  if (Array.isArray(enumData))
    filterData = findItemNested(enumData, callValue, value, children);
  // 判断是否输出的结果为 tag 类型
  if (type == "tag") {
    return filterData?.tagType ? filterData.tagType : "";
  } else {
    return filterData ? filterData[label] : "--";
  }
}

/**
 * @description 递归查找 callValue 对应的 enum 值
 * */
export function findItemNested(enumData: any, callValue: any, value: string, children: string) {
  return enumData.reduce((accumulator: any, current: any) => {
    if (accumulator) return accumulator;
    if (current[value] === callValue) return current;
    if (current[children]) return findItemNested(current[children], callValue, value, children);
  }, null);
}

/**
* @description 处理值无数据情况
* @param {*} callValue 需要处理的值
* @returns {String}
* */
export function formatValue(callValue: any) {
  // 如果当前值为数组，使用 / 拼接（根据需求自定义）
  if (isArray(callValue)) return callValue.length ? callValue.join(" / ") : "--";
  return callValue ?? "--";
}

/**
* @description:  是否为数组
*/
export function isArray(val: any): val is Array<any> {
  return val && Array.isArray(val);
}


/**
* @description 防抖函数
* @param {fn()} 对那个函数进行防抖
* @param {delay} 多长时间才发送请求
* */
export function debounce(fn: () => void, delay: number) {
  let time: number;
  return function () {
    if (time) {
      clearTimeout(time);
    }
    time = setTimeout(() => {
      fn();
    }, delay);
  };
};

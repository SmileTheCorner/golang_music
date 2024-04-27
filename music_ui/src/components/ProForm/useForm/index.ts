import { reactive,toRefs} from "vue"
interface PageAble {
    pageNum: number;
    pageSize: number;
    total: number;
}
interface StateProps {
    tableData: any[];
    pageAble: PageAble;
    searchParam: {
        [key: string]: any;
    };
    totalParam: {
        [key: string]: any;
    };
}
/**
 * 封装表单hook方法
 */

export const useForm = (
    //搜索接口
    api : (params: any) => Promise<any>,
    //初始化参数
    searchInitParam ?: object,
    //是否分页
    isPageAble : boolean=false,
    //请求成功回调
    dataCallBack?: (data: any) => any,
    //请求错误回调
    requestError?: (error: any) => void
) => {
    const state = reactive<StateProps>({
        // 表格数据
        tableData: [],
        // 分页数据
        pageAble: {
            // 当前页数
            pageNum: 1,
            // 每页显示条数
            pageSize: 10,
            // 总条数
            total: 0,
        },
        // 查询参数(只包括查询)
        searchParam: {},
        // 总参数(包含分页和查询参数)
        totalParam: {},
    })

    //表格查询数据
    const search = () => {
        //将要查询的页数设置为1
        state.pageAble.pageNum = 1
        //更新要查询的总参数列表
        updatedTotalParam();
        //重新获取表格数据
        getTableList();
    }

    //重置表单
    const reset = () => {
        //将查询的页数重置为1
        state.pageAble.pageNum = 1;
        searchInitParam = {};
        //跟新总的查询参数
        updatedTotalParam();
        //重新获取表格数据
        getTableList();
      };

    //更新参数
    const updatedTotalParam = () => {
        //将总的查询参数列表置为空
        state.totalParam = {};
        state.searchParam = {};
        state.searchParam = {...searchInitParam}
        //将查询参数和分页参数合并
        let param = {
            "keyWord": { ...state.searchParam },
        }
        state.totalParam = { ...param, ...state.pageAble }
    }

    //获取表格数据
    const getTableList = async () => {
        if (!api) return
        try {
            //获取数据
            let data = await api(state.totalParam);
            //如果有回调函数则将数据返回
            dataCallBack && dataCallBack(data)
            //给state中的数据列表赋值 如果是有分页则去data.list 否则取data
            state.tableData = isPageAble ? data.list : data;
            // 解构后台返回的分页数据 (如果有分页更新分页信息)
            if (isPageAble) {
                const { pageNum, pageSize, total } = data;
                //更新表格参数
                updatePageAble({ pageNum, pageSize, total });
            }
        } catch (error) {
            requestError?.(error)
        }
    }

    //更新分页数据
    const updatePageAble = (pageAble: PageAble) => {
        state.pageAble = pageAble;
    }

    /**
   * @description 每页条数改变
   * @param {Number} val 当前条数
   * @return void
   * */
  const handleSizeChange = (val: number) => {
    state.pageAble.pageNum = 1;
    state.pageAble.pageSize = val;
    getTableList();
  };

  /**
   * @description 当前页改变
   * @param {Number} val 当前页
   * @return void
   * */
  const handleCurrentChange = (val: number) => {
    state.pageAble.pageNum = val;
    getTableList();
  };
  
  //将函数暴露出去
  return {
    ...toRefs(state),
    handleSizeChange,
    handleCurrentChange,
    search,
    reset,
    getTableList,
  };

}
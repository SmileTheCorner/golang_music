<template>
    <div class="pagination">
        <el-pagination v-model:current-page="currentPage  " v-model:page-size="pageSize" :page-sizes="[100, 200, 300, 400]"
            :small="small" :disabled="disabled" :background="background" :layout="layout"
            :total="total" @size-change="handleSizeChange" @current-change="handleCurrentChange" />
    </div>
</template>
 
<script setup>
import { ref, reactive, onMounted } from "vue"
const props = defineProps({
    total: {
        type: Number,
        default: 0, //总页数
    },
    pagerCount: {
        type: Number,
        default: 5, //如果页数很多大概展示的页码
    },
    layout: {
        type: String,
        default: "total,sizes, prev, pager, next, jumper, ->, slot", //分页组件会展示的功能项
    },
    pageSize: {
        type: Array,
        default: () => {
            return [10, 20, 50, 100] //指定分页展示条数
        },
    },
    currentPage: {
        type: Number,
        default: 1, //指定跳转到多少页
    },
    pageSize: {
        type: Number,
        default: 10, //每页展示的条数，根据自己实际，且会带入请求
    },
    pageNum: {
        type: Number,
        default: 1, //第几页数据，根据自己实际，且会带入请求
    },
})
</script>
<template>
  <!-- 分页组件 -->
  <el-pagination
    :background="true"
    :current-page="pageAble.pageNum"
    :page-size="pageAble.pageSize"
    :page-sizes="[10, 20, 50, 100]"
    :total="pageAble.total"
    layout="total, sizes, prev, pager, next, jumper"
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
  ></el-pagination>
</template>

<script setup lang="ts" name="Pagination">
interface PageAble {
  pageNum: number;
  pageSize: number;
  total: number;
}

interface PaginationProps {
  pageAble: PageAble;
  handleSizeChange: (size: number) => void;
  handleCurrentChange: (currentPage: number) => void;
}

defineProps<PaginationProps>();
</script>

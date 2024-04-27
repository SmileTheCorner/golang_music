<template>
  <RenderTableColumn v-bind="column" />
</template>

<script setup lang="tsx" name="TableColumn">
import {
  ColumnProps,
  RenderScope,
} from "@/components/ProTable/interface/index";
import { handleProp } from "@/utils/index";
import {
  filterEnum,
  formatValue,
  handleRowAccordingToProp,
} from "@/utils/index";
import { useSlots, inject, ref } from "vue";

defineProps<{ column: ColumnProps }>();
const slots = useSlots();
const enumMap = inject("enumMap", ref(new Map()));

// 获取 tag 类型
const getTagType = (item: ColumnProps, scope: RenderScope<any>) => {
  return filterEnum(
    handleRowAccordingToProp(scope.row, item.prop!),
    enumMap.value.get(item.prop),
    item.fieldNames,
    "tag"
  );
};

// 渲染表格数据
const renderCellData = (item: ColumnProps, scope: RenderScope<any>) => {
  return enumMap.value.get(item.prop) && item.isFilterEnum
    ? filterEnum(
        handleRowAccordingToProp(scope.row, item.prop!),
        enumMap.value.get(item.prop)!,
        item.fieldNames
      )
    : formatValue(handleRowAccordingToProp(scope.row, item.prop!));
};

const RenderTableColumn = (item: ColumnProps) => {
  return (
    <>
      {item.isShow && (
        <el-table-column {...item} align={item.align ?? "center"}>
          {{
            default: (scope: RenderScope<any>) => {
              if (item._children)
                return item._children.map((child) => RenderTableColumn(child));
              if (item.render) return item.render(scope);
              if (slots[handleProp(item.prop!)])
                return slots[handleProp(item.prop!)]!(scope);
              if (item.tag)
                return (
                  <el-tag type={getTagType(item, scope)}>
                    {renderCellData(item, scope)}
                  </el-tag>
                );
              return renderCellData(item, scope);
            },
          }}
        </el-table-column>
      )}
    </>
  );
};
</script>

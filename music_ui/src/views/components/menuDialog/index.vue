<template>
  <el-dialog
    v-model="centerDialogVisible"
    :title="dialogProps.title"
    width="50%"
    align-center
    @close="close"
  >
    <el-form
      ref="ruleFormRef"
      :disabled="dialogProps.isView"
      :model="dialogProps.rowData"
      :rules="rules"
      label-width="100px"
    >
      <el-row>
        <!-- 上级菜单 -->
        <el-col :span="24">
          <el-form-item label="上级菜单" class="parent-menu">
            <el-tree-select
              v-model="dialogProps.rowData.meta!.parentId"
              :data="menuOptions"
              placeholder="选择上级菜单"
              check-strictly
              @current-change="currentChange"
            />
          </el-form-item>
        </el-col>

        <!-- 菜单类型 -->
        <el-col :span="24">
          <el-form-item label="菜单类型" prop="meta.menuType">
            <el-radio-group v-model="dialogProps.rowData.meta!.menuType">
              <el-radio label="M">目录</el-radio>
              <el-radio label="C">菜单</el-radio>
              <el-radio label="F">按钮</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>

        <!-- 菜单图标 -->
        <el-col :span="24" v-if="dialogProps.rowData.meta!.menuType != 'F'">
          <el-form-item label="菜单图标">
            <el-popover placement="bottom-start" :width="540" trigger="click">
              <template #reference>
                <el-input
                  v-model="dialogProps.rowData.meta!.icon"
                  placeholder="点击选择图标"
                  readonly
                >
                  <template #prefix>
                    <svg-icon
                      v-if="dialogProps.rowData.meta!.icon"
                      :icon-class="dialogProps.rowData.meta!.icon"
                      class="el-input__icon"
                      style="height: 32px; width: 16px"
                    />
                    <el-icon v-else style="height: 32px; width: 16px">
                      <search />
                    </el-icon>
                  </template>
                </el-input>
              </template>
              <icon-select
                ref="iconSelectRef"
                @selected="selectedIcon"
                :active-icon="dialogProps.rowData.meta!.icon"
              />
            </el-popover>
          </el-form-item>
        </el-col>

        <!-- 菜单标题 -->
        <el-col :span="12">
          <el-form-item prop="title">
            <template #label>
              <span>
                <el-tooltip
                  placement="top"
                  content="填写相应的中文名称如:系统管理"
                >
                  <el-icon><question-filled /></el-icon>
                </el-tooltip>
                标题
              </span>
            </template>
            <el-input
              v-model="dialogProps.rowData.meta.title"
              placeholder="请输入菜单标题"
            />
          </el-form-item>
        </el-col>

        <!-- 菜单名称 -->
        <el-col :span="12">
          <el-form-item prop="name">
            <template #label>
              <span>
                <el-tooltip
                  placement="top"
                  content="填写相应的英文名称如:system"
                >
                  <el-icon><question-filled /></el-icon>
                </el-tooltip>
                名称
              </span>
            </template>
            <el-input
              v-model="dialogProps.rowData.name"
              placeholder="请输入菜单名称"
            />
          </el-form-item>
        </el-col>

        <!-- 显示排序 -->
        <el-col :span="12">
          <el-form-item label="显示排序" prop="meta.orderNum">
            <el-input-number
              v-model="dialogProps.rowData.meta!.orderNum"
              controls-position="right"
              :min="0"
            />
          </el-form-item>
        </el-col>

        <!-- 是否外链 -->
        <el-col :span="12" v-if="dialogProps.rowData.meta!.menuType != 'F'">
          <el-form-item>
            <template #label>
              <span>
                <el-tooltip
                  content="选择是外链则路由地址需要以`http(s)://`开头"
                  placement="top"
                >
                  <el-icon><question-filled /></el-icon>
                </el-tooltip>
                是否外链
              </span>
            </template>
            <el-radio-group v-model="dialogProps.rowData.meta!.isFrame">
              <el-radio :label="0">是</el-radio>
              <el-radio :label="1">否</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>

        <!-- 是否全屏 -->
        <el-col :span="12" v-if="dialogProps.rowData.meta!.menuType != 'F'">
          <el-form-item>
            <template #label>
              <span>
                <el-tooltip
                  placement="top"
                  content="在浏览器中展示的时候是否以全屏的方式展示"
                >
                  <el-icon><question-filled /></el-icon> </el-tooltip
                >是否全屏
              </span>
            </template>
            <el-radio-group v-model="dialogProps.rowData.meta!.isFull">
              <el-radio :label="1">是</el-radio>
              <el-radio :label="0">否</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>

        <!-- 路由地址 -->
        <el-col :span="12" v-if="dialogProps.rowData.meta!.menuType != 'F'">
          <el-form-item prop="path">
            <template #label>
              <span>
                <el-tooltip
                  placement="top"
                  content="访问的路由地址，如:`/user`，根据该路由地址则可以访问组件"
                >
                  <el-icon><question-filled /></el-icon>
                </el-tooltip>
                路由地址
              </span>
            </template>
            <el-input
              v-model="dialogProps.rowData!.path"
              placeholder="请输入路由地址"
            />
          </el-form-item>
        </el-col>

        <!-- 组件路径 -->
        <el-col :span="12" v-if="dialogProps.rowData.meta!.menuType == 'C'">
          <el-form-item prop="component">
            <template #label>
              <span>
                <el-tooltip
                  content="访问的组件路径，如：`system/user/index`，默认在`views`目录下"
                  placement="top"
                >
                  <el-icon><question-filled /></el-icon>
                </el-tooltip>
                组件路径
              </span>
            </template>
            <el-input
              v-model="dialogProps.rowData!.component"
              placeholder="请输入组件路径"
            />
          </el-form-item>
        </el-col>

        <!-- 权限字符 -->
        <el-col :span="12" v-if="dialogProps.rowData.meta!.menuType != 'M'">
          <el-form-item>
            <el-input
              v-model="dialogProps.rowData.meta!.perms"
              placeholder="请输入权限标识"
              maxlength="100"
            />
            <template #label>
              <span>
                <el-tooltip
                  content="控制器中定义的权限字符，如：@PreAuthorize(`@ss.hasPermi('system:user:list')`)"
                  placement="top"
                >
                  <el-icon><question-filled /></el-icon>
                </el-tooltip>
                权限字符
              </span>
            </template>
          </el-form-item>
        </el-col>

        <!-- 路由参数 -->
        <el-col :span="12" v-if="dialogProps.rowData.meta!.menuType == 'C'">
          <el-form-item>
            <el-input
              v-model="dialogProps.rowData.meta!.query"
              placeholder="请输入路由参数"
              maxlength="255"
            />
            <template #label>
              <span>
                <el-tooltip
                  content='访问路由的默认传递参数，如：`{"id": 1, "name": "ry"}`'
                  placement="top"
                >
                  <el-icon><question-filled /></el-icon>
                </el-tooltip>
                路由参数
              </span>
            </template>
          </el-form-item>
        </el-col>

        <!-- 是否缓存 -->
        <el-col :span="12" v-if="dialogProps.rowData.meta!.menuType == 'C'">
          <el-form-item>
            <template #label>
              <span>
                <el-tooltip
                  content="选择是则会被`keep-alive`缓存，需要匹配组件的`name`和地址保持一致"
                  placement="top"
                >
                  <el-icon><question-filled /></el-icon>
                </el-tooltip>
                是否缓存
              </span>
            </template>
            <el-radio-group v-model="dialogProps.rowData.meta!.isCache">
              <el-radio :label="0">缓存</el-radio>
              <el-radio :label="1">不缓存</el-radio>
            </el-radio-group>
          </el-form-item>
        </el-col>

        <!-- 显示状态 -->
        <el-col :span="12" v-if="dialogProps.rowData.meta!.menuType != 'F'">
          <el-form-item>
            <template #label>
              <span>
                <el-tooltip
                  content="选择隐藏则路由将不会出现在侧边栏，但仍然可以访问"
                  placement="top"
                >
                  <el-icon><question-filled /></el-icon>
                </el-tooltip>
                显示状态
              </span>
            </template>
            <el-radio-group v-model="dialogProps.rowData.meta!.visible">
              <el-radio
                v-for="dict in displayStatus"
                :key="dict.value"
                :label="parseInt(dict.value)"
                >{{ dict.label }}</el-radio
              >
            </el-radio-group>
          </el-form-item>
        </el-col>

        <!-- 菜单状态 -->
        <el-col :span="12" v-if="dialogProps.rowData.meta!.menuType != 'F'">
          <el-form-item>
            <template #label>
              <span>
                <el-tooltip
                  content="选择停用则路由将不会出现在侧边栏，也不能被访问"
                  placement="top"
                >
                  <el-icon><question-filled /></el-icon>
                </el-tooltip>
                菜单状态
              </span>
            </template>
            <el-radio-group v-model="dialogProps.rowData.meta!.status">
              <el-radio
                v-for="dict in menuStatus"
                :key="dict.value"
                :label="parseInt(dict.value)"
                >{{ dict.label }}</el-radio
              >
            </el-radio-group>
          </el-form-item>
        </el-col>

        <!-- 备注 -->
        <el-col :span="24">
          <el-form-item label="备注">
            <el-input
              v-model="dialogProps.rowData.meta!.remark"
              placeholder="请输入备注信息"
            />
          </el-form-item>
        </el-col>
      </el-row>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="cancel">取消</el-button>
        <el-button type="primary" @click="makeSure(ruleFormRef)">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script setup lang="ts" name="menuDialog">
import IconSelect from "@/components/IconSelect/index.vue";
import SvgIcon from "@/components/SvgIcon/index.vue";
import { addMenuTree } from "@/api/system/menu/index";
import { ref, reactive, onMounted } from "vue";
import type { ElForm } from "element-plus";
import { ElMessage } from "element-plus";
import { getDictDataByType } from "@/api/system/dict/index";

type FormInstance = InstanceType<typeof ElForm>;
type FormRules = InstanceType<typeof ElForm>;

const menuStatus = ref<any[]>([]);
const displayStatus = ref<any[]>([]);
onMounted(async () => {
  try {
    const menuStatusData = await getDictDataByType({ dictType: "menu_status" });
    menuStatus.value.push(...(menuStatusData as unknown as any[]));
    const displayStatusData = await getDictDataByType({ dictType: "menu_display" });
    displayStatus.value.push(...(displayStatusData as unknown as any[]));
  } catch (error) {}
});

const rules = reactive<FormRules>({
  name: [{ required: true, message: "请输入菜单名称", trigger: "blur" }],
  "meta.menuType": [
    { required: true, message: "请选择菜单类型", trigger: "blur" },
  ],
  "meta.orderNum": [
    { required: true, message: "请输入排序数值", trigger: "blur" },
  ],
  path: [{ required: true, message: "请输入路由地址", trigger: "blur" }],
  component: [{ required: true, message: "请输入组件路径", trigger: "blur" }],
});

//定义对话框类型
interface DialogProps {
  title: string;
  isView: boolean;
  rowData: Partial<Menu.MenuOptions>;
  api?: (params: any) => Promise<any>;
  getTableList?: () => void; //重新获取数据接口
}

//显示参数
const dialogProps = ref<DialogProps>({
  isView: false,
  title: "",
  rowData: {
    meta: {}, // 提供一个默认值
  },
});

//菜单下拉选项
const menuOptions = ref<Array<{ [key: string]: any }>>([]);
//控制对话框是否显示
const centerDialogVisible = ref(false);
// 接收父组件传过来的参数
const acceptParams = (params: DialogProps) => {
  dialogProps.value = params;
  centerDialogVisible.value = true;
};
/** 查询菜单下拉树结构 */
async function getTreeSelect() {
  menuOptions.value = [];
  try {
    const data = await addMenuTree();
    menuOptions.value = data;
    const menu = { value: "", label: "主类目", children: [] };
    menuOptions.value.unshift(menu);
  } catch (error) {}
}
getTreeSelect();

/** 选择图标 */
function selectedIcon(name: string) {
  dialogProps.value.rowData.meta!.icon = name;
}
/**表单校验ref */
const ruleFormRef = ref<FormInstance>();
/**取消按钮 */
const cancel = () => {
  resetForm(ruleFormRef.value);
  centerDialogVisible.value = false;
};
/**确定 */
const makeSure = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  //表单校验
  await formEl.validate(async (valid: boolean) => {
    if (valid) {
      let metaInfo = dialogProps.value.rowData.meta;
      let name = dialogProps.value.rowData.name;
      let path = dialogProps.value.rowData.path;
      let component = dialogProps.value.rowData.component;
      let params = {
        name,
        path,
        component,
        ...metaInfo,
      };
      await dialogProps.value.api!(params)
        .then((res) => {
          ElMessage.success({
            message: `${dialogProps.value.rowData.name}操作成功！`,
          });
          //重新获取表单数据
          dialogProps.value.getTableList?.();
          resetForm(ruleFormRef.value);
          centerDialogVisible.value = false;
        })
        .catch((err) => {
          ElMessage.error(
            `${dialogProps.value.rowData.name}操作失败！错误原因${err.message}`
          );
        });
    } else {
      ElMessage.error("请填写带*号的选项。。。");
    }
  });
};
/**清空表单 */
const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};
/**关闭对话框 */
const close = () => {
  resetForm(ruleFormRef.value);
};

/** 当父节点发生改变时触发 */
const currentChange = (e: { value: string }) => {
  if (dialogProps.value.rowData.meta) {
    dialogProps.value.rowData.meta.parentId = e.value;
  }
};

//将需要暴露出去的数据与方法都可以暴露出去
defineExpose({
  acceptParams,
});
</script>
<style scoped>
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>

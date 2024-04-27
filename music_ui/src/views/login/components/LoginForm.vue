<template>
  <el-form
    ref="loginFormRef"
    :model="loginForm"
    :rules="loginRules"
    size="large"
  >
    <el-form-item prop="username">
      <el-input v-model="loginForm.username" placeholder="用户名">
        <template #prefix>
          <el-icon class="el-input__icon">
            <user />
          </el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item prop="password">
      <el-input
        type="password"
        v-model="loginForm.password"
        placeholder="密码"
        show-password
        autocomplete="new-password"
      >
        <template #prefix>
          <el-icon class="el-input__icon">
            <lock />
          </el-icon>
        </template>
      </el-input>
    </el-form-item>
  </el-form>
  <div class="login-btn">
    <el-button
      icon="CircleClose"
      round
      @click="resetForm(loginFormRef)"
      size="large"
      >重置</el-button
    >
    <el-button
      icon="UserFilled"
      round
      @click="login(loginFormRef)"
      size="large"
      type="primary"
    >
      登录
    </el-button>
  </div>
</template>

<script setup lang="ts">
import { Login } from "../../../type/login/index";
import { ElForm } from "element-plus";
import { reactive, ref } from "vue";
import { doLogin } from "@/api/system/user/index";
import GlobalStore from "@/store/module/global";
import { initDynamicRouter } from "@/router/module/dynamicRouter";
import  {useRouter}  from "vue-router";
import { ElNotification } from "element-plus";
import { HOME_URL } from "@/config/index";
import { getTimeState } from "@/utils/index";

const globalStore = GlobalStore();
const router = useRouter();
type FormInstance = InstanceType<typeof ElForm>;
const loginFormRef = ref<FormInstance>();
const loginRules = reactive({
  username: [
    {
      required: true,
      message: "请输入用户名",
      trigger: "blur",
    },
  ],
  password: [
    {
      required: true,
      message: "请输入密码",
      trigger: "blur",
    },
  ],
});
const loginForm = reactive<Login.ReqLoginForm>({
  username: "转角的微笑",
  password: "123456",
});
const login = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.validate(async (valid: any) => {
    if (!valid) return;
    try {
      let data = await doLogin(loginForm);
      //设置token
      globalStore.setToken(data.access_token);
        //登录成功添加动态路由
        await initDynamicRouter();
        //登录成功跳到首页
        router.push("/layout")
        ElNotification({
          title: getTimeState(),
          message: "欢迎登录",
          type: "success",
          duration: 3000,
        });
    }catch (err) {
      //登录失败清空token
      globalStore.setToken("");
    }
  });
};

// resetForm
const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};
</script>

<style scoped lang="scss">
@import "../index.scss";
</style>
../../../type/user/index@/api/system/user/index
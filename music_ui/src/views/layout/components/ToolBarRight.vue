<template>
  <div class="right">
    <AssemblySize class="operation-icon"></AssemblySize>
    <SearchMenu class="operation-icon"></SearchMenu>
    <span
      class="iconfont icon-clothes-full operation-icon"
      @click="changLayout"
    ></span>
    <Notify class="operation-icon"></Notify>
    <span
      class="operation-icon"
      :class="[
        full ? 'iconfont icon-cancel-full-screen' : 'iconfont icon-quanping',
      ]"
      @click="changFull"
    ></span>
    <span class="operation-icon">admin</span>

    <el-dropdown>
      <img src="@/assets/images/giphy.gif" alt="avatar" />
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>个人信息</el-dropdown-item>
          <el-dropdown-item>修改密码</el-dropdown-item>
          <el-dropdown-item>退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
    <el-drawer v-model="drawerLayout" title="布局设置" size="25%">
      <div class="theme-color">
        <el-divider>
          <el-icon size="25"><Histogram /></el-icon>
          <span class="title">布局样式</span>
        </el-divider>
        <div class="layout-items">
          <!-- 经典 classic-->
          <div class="card box-shadow-card" @click="setLayout('classic')">
            <div class="container-horizontal">
              <div class="top"></div>
              <div class="bottom">
                <div class="left"></div>
                <div class="right">
                  <el-icon class="icon"><CircleCheck /></el-icon>
                </div>
              </div>
            </div>
          </div>

          <!-- 纵向 vertical-->
          <div class="card box-shadow-card" @click="setLayout('vertical')">
            <div class="container">
              <div class="left"></div>
              <div class="right">
                <div class="top"></div>
                <div class="bottom">
                  <el-icon class="icon"><CircleCheck /></el-icon>
                </div>
              </div>
            </div>
          </div>

          <!-- 横向 crosswise-->
          <div class="card box-shadow-card" @click="setLayout('crosswise')">
            <div class="container-vertical">
              <div class="top"></div>
              <div class="bottom">
                <el-icon class="icon"><CircleCheck /></el-icon>
              </div>
            </div>
          </div>
          <!-- 分栏 columns-->
          <div class="card box-shadow-card" @click="setLayout('columns')">
            <div class="container-endwise">
              <div class="left"></div>
              <div class="middle"></div>
              <div class="right">
                <el-icon class="icon"><CircleCheck /></el-icon>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="theme-color">
        <el-divider>
          <el-icon size="25"><Cloudy /></el-icon>
          <span class="title">全局主题</span>
        </el-divider>
        <div class="items">
          <div class="item">
            <span>主题颜色</span>
            <span>
              <el-color-picker
                v-model="globalSettings.color"
                show-alpha
                :predefine="predefineColors"
                @change="changePrimary"
              />
            </span>
          </div>
          <div class="item">
            <span>暗黑模式</span>
            <span>
              <el-switch
                v-model="globalSettings.dark"
                style="margin-left: 24px"
                inline-prompt
                :active-icon="Sunny"
                :inactive-icon="Moon"
                @change="changeDark"
              />
            </span>
          </div>
          <div class="item">
            <span>灰色模式</span>
            <span>
              <el-switch
                v-model="globalSettings.grey"
                style="margin-left: 24px"
                inline-prompt
                @change="changeGreyOrWeak('grey', !!$event)"
              />
            </span>
          </div>
          <div class="item">
            <span>弱色模式</span>
            <span>
              <el-switch
                v-model="globalSettings.weak"
                style="margin-left: 24px"
                inline-prompt
                @change="changeGreyOrWeak('weak', !!$event)"
              />
            </span>
          </div>
        </div>
      </div>
      <div class="theme-color">
        <el-divider>
          <el-icon size="25"><Setting /></el-icon>
          <span class="title">界面设置</span>
        </el-divider>
        <div class="items">
          <div class="item">
            <span>页脚显示</span>
            <span>
              <el-switch
                v-model="globalSettings.footer"
                style="margin-left: 24px"
                inline-prompt
                active-text="是"
                inactive-text="否"
                @change="setFooter(!!$event)"
              />
            </span>
          </div>
          <div class="item">
            <span>菜单折叠</span>
            <span>
              <el-switch
                v-model="globalSettings.collapse"
                style="margin-left: 24px"
                inline-prompt
                active-text="是"
                inactive-text="否"
                @change="setCollapse(!!$event)"
              />
            </span>
          </div>
        </div>
      </div>
    </el-drawer>
  </div>
</template>
<script setup lang="ts">
import { ref, reactive } from "vue";
import { storeToRefs } from "pinia";
import { Sunny, Moon } from "@element-plus/icons-vue";
import {
  setThemeColor,
  setThemeDark,
  setGreyOrWeak,
  setFooter,
  setCollapse,
} from "@/utils/globalSettings";
import GlobalStore from "@/store/module/global";
import { LayoutType } from "@/store/interface/index";
import SearchMenu from "./SearchMenu.vue";
import AssemblySize from "./AssemblySize.vue"
import Notify from "./Notify.vue"

const globalStore = GlobalStore();
const full = ref<boolean>(false);
const drawerLayout = ref<boolean>(false);

const setting = storeToRefs(globalStore);
const globalSettings = reactive({
  color: setting.themeConfig.value.primary,
  dark: setting.themeConfig.value.isDark,
  grey: setting.themeConfig.value.isGrey,
  weak: setting.themeConfig.value.isWeak,
  footer: setting.themeConfig.value.footer,
  collapse: setting.themeConfig.value.isCollapse,
});

const boxShadowColor = setting.themeConfig.value.isDark
  ? "rgba(255, 255, 255,0.9)"
  : "rgba(0, 0, 0, 0.5)";

const predefineColors = ref([
  "#ff4500",
  "#ff8c00",
  "#ffd700",
  "#90ee90",
  "#00ced1",
  "#1e90ff",
  "#c71585",
  "rgba(255, 69, 0, 0.68)",
  "rgb(255, 120, 0)",
  "hsv(51, 100, 98)",
  "hsva(120, 40, 94, 0.5)",
  "hsl(181, 100%, 37%)",
  "hsla(209, 100%, 56%, 0.73)",
  "#c7158577",
]);

//全屏
const changFull = () => {
  let isFull = document.fullscreenElement;
  full.value = isFull === null ? true : false;
  if (!isFull) {
    //全屏则退出全屏
    document.documentElement.requestFullscreen();
  } else {
    //不是全屏则展开全屏
    document.exitFullscreen();
  }
};

//改变布局
const changLayout = () => {
  drawerLayout.value = !drawerLayout.value;
};

//改变暗黑模式
const changeDark = (e: boolean) => {
  globalSettings.dark = e;
  globalStore.setThemeDark(e);
  setThemeDark(globalSettings.dark);
  setShadow(e);
};

//设置布局移入和移出边框阴影
const setShadow = (e: boolean) => {
  let card = document.getElementsByClassName("box-shadow-card");
  Array.from(card).forEach((item: any) => {
    item.onmouseover = function () {
      item.style.boxShadow = e
        ? "0px 2px 5px rgba(255, 255, 255,0.9)"
        : "0px 2px 5px rgba(0, 0, 0, 0.5)";
    };
    item.onmouseout = function () {
      item.style.boxShadow = "none";
    };
  });
};

//设置选中的模式
const setLayout = (model: LayoutType) => {
  globalStore.setLayoutModel(model);
  console.log(model);
};

//修改主题颜色
const changePrimary = (color: string | null) => {
  globalStore.setThemeColor(color);
  //设置主题颜色
  setThemeColor(color, globalSettings.dark);
};

//改变灰色模式和色弱模式
const changeGreyOrWeak = (type: "grey" | "weak", value: boolean) => {
  value && type === "grey"
    ? (globalSettings.weak = false)
    : (globalSettings.grey = false);
  setGreyOrWeak(type, value);
};
</script>
<style scoped lang="scss">
:deep(.el-divider__text) {
  display: flex;
  justify-content: center;
  align-items: center;
}
.right {
  display: flex;
  justify-content: center;
  align-items: center;
  .operation-icon {
    margin: 0 6px;
  }
  img {
    width: 35px;
    height: 35px;
    border-radius: 50%;
  }
  .theme-color {
    .title {
      font-size: 18px;
      font-weight: bold;
    }
    .layout-items {
      display: grid;
      grid-template-columns: repeat(2, 1fr);
      grid-gap: 10px;
      .card {
        padding: 10px;
        .container {
          width: 100%;
          height: 100px;
          display: flex;
          flex-direction: row;
          align-items: center;
          .left {
            height: 100%;
            width: 14%;
            border-radius: 4px;
            background-color: var(--el-color-primary);
          }
          .right {
            flex: 1;
            width: 100%;
            height: 100%;
            margin-left: 10px;
            display: flex;
            flex-direction: column;
            .top {
              width: 100%;
              height: 13%;
              border-radius: 4px;
              background-color: var(--el-color-primary-light-5);
            }
            .bottom {
              position: relative;
              box-sizing: border-box;
              margin-top: 10px;
              flex: 1;
              width: 100%;
              height: 100%;
              border-radius: 4px;
              border: 1px dashed var(--el-color-primary);
              background-color: var(--el-color-primary-light-6);
              .icon {
                position: absolute;
                left: 50%;
                top: 50%;
                transform: translate(-50%, -50%);
              }
            }
          }
        }

        .container-horizontal {
          width: 100%;
          height: 100px;
          display: flex;
          flex-direction: column;
          .top {
            width: 100%;
            height: 13%;
            border-radius: 4px;
            background-color: var(--el-color-primary);
          }
          .bottom {
            flex: 1;
            width: 100%;
            height: 100%;
            margin-top: 10px;
            display: flex;
            flex-direction: row;
            .left {
              width: 14%;
              height: 100%;
              border-radius: 4px;
              background-color: var(--el-color-primary-light-5);
            }
            .right {
              position: relative;
              box-sizing: border-box;
              margin-left: 10px;
              flex: 1;
              width: 100%;
              height: 100%;
              border-radius: 4px;
              border: 1px dashed var(--el-color-primary);
              background-color: var(--el-color-primary-light-6);
              .icon {
                position: absolute;
                left: 50%;
                top: 50%;
                transform: translate(-50%, -50%);
              }
            }
          }
        }
        .container-vertical {
          width: 100%;
          height: 100px;
          display: flex;
          flex-direction: column;
          .top {
            width: 100%;
            height: 13%;
            border-radius: 4px;
            background-color: var(--el-color-primary);
          }
          .bottom {
            position: relative;
            flex: 1;
            width: 100%;
            height: 100%;
            margin-top: 10px;
            border-radius: 4px;
            box-sizing: border-box;
            border: 1px dashed var(--el-color-primary);
            background-color: var(--el-color-primary-light-6);
            .icon {
              position: absolute;
              left: 50%;
              top: 50%;
              transform: translate(-50%, -50%);
            }
          }
        }
        .container-endwise {
          width: 100%;
          height: 100px;
          display: flex;
          flex-direction: row;
          .left {
            height: 100%;
            width: 14%;
            border-radius: 4px;
            background-color: var(--el-color-primary);
          }
          .middle {
            width: 13%;
            height: 100%;
            border-radius: 4px;
            background-color: var(--el-color-primary-light-5);
            margin-left: 10px;
          }
          .right {
            position: relative;
            margin-left: 10px;
            box-sizing: border-box;
            flex: 1;
            width: 100%;
            height: 100%;
            border-radius: 4px;
            border: 1px dashed var(--el-color-primary);
            background-color: var(--el-color-primary-light-6);
            .icon {
              position: absolute;
              left: 50%;
              top: 50%;
              transform: translate(-50%, -50%);
            }
          }
        }
      }
      .card:hover {
        box-shadow: 0px 2px 5px v-bind(boxShadowColor);
      }
    }
    .items {
      width: 80%;
      margin: 0 auto;
      .item {
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin-bottom: 10px;
      }
    }
  }
}
</style>

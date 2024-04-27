import { getDarkColor, getLightColor } from "@/config/color";
import { DEFAULT_PRIMARY } from "@/config/index"
import GlobalStore from "@/store/module/global";



//设置主题颜色
export function setThemeColor(color: string | null, dark: boolean) {
    if (!color) {
        color = DEFAULT_PRIMARY;
    }
    // 计算主题颜色变化
    document.documentElement.style.setProperty("--el-color-primary", color);
    document.documentElement.style.setProperty(
        "--el-color-primary-dark-2",
        dark
            ? `${getLightColor(color, 0.2)}`
            : `${getDarkColor(color, 0.3)}`
    );
    for (let i = 1; i <= 9; i++) {
        const primaryColor = dark
            ? `${getDarkColor(color, i / 10)}`
            : `${getLightColor(color, i / 10)}`;
        document.documentElement.style.setProperty(
            `--el-color-primary-light-${i}`,
            primaryColor
        );
    }
}

//设置暗黑模式
export function setThemeDark(dark: boolean) {
    const html = document.documentElement as HTMLElement;
    if (dark) {
        html.setAttribute("class", "dark");
        document.getElementsByTagName("body")[0].style.setProperty("--backgroundPrimaryColor", "#141414");
        document.getElementsByTagName("body")[0].style.setProperty("--borderColor", "#414243");
    } else {
        html.setAttribute("class", "");
        document.getElementsByTagName("body")[0].style.setProperty("--backgroundPrimaryColor", "#f2f3f5");
        document.getElementsByTagName("body")[0].style.setProperty("--borderColor", "#e4e7ed");
    }
}

//灰色和弱色模式切换
export function setGreyOrWeak(type: "grey" | "weak", value: boolean) {
    const globalStore = GlobalStore();
    const body = document.body as HTMLElement
    const propName = type === "grey" ? "isGrey" : "isWeak"
    //如果不是灰色和弱色模式，则移除样式
    if (!value) {
        body.removeAttribute("style")
        globalStore.setThemeGreyOrWeak(propName, value)
    } else {
        //定义灰色和弱色模式的样式
        const styles: Record<"grey" | "weak", string> = {
            grey: "filter: grayscale(1);overflow:hidden",
            weak: "filter: invert(70%);overflow:hidden"
        }
        //添加样式
        body.setAttribute("style", styles[type]);
        //修改gloable状态
        globalStore.setThemeGreyOrWeak(propName, value)
    }
}

//设置底部的显示和掩藏
export function setFooter(value: boolean) {
    const globalStore = GlobalStore();
    globalStore.setFooter(value);
}

//设置菜单栏的折叠和展开
export function setCollapse(value: boolean) {
    const globalStore = GlobalStore();
    globalStore.setMenuCollapse(value);
}
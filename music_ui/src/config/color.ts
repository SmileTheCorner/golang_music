import { ElMessage } from "element-plus";


/**
 * @description rgb颜色转Hex颜色
 * @param {*} r 代表红色
 * @param {*} g 代表绿色
 * @param {*} b 代表蓝色
 * @param {*} a 代表透明度
 * @returns {String} 返回处理后的颜色值
 */
export function rgbaToHex(r: number, g: number, b: number, a?: number): string {
    function componentToHex(c: number): string {
        const hex = c.toString(16);
        return hex.length === 1 ? '0' + hex : hex;
    }

    const redHex = componentToHex(r);
    const greenHex = componentToHex(g);
    const blueHex = componentToHex(b);

    let alphaHex = '';
    if (typeof a !== 'undefined') {
        alphaHex = Math.round(a * 255).toString(16);
        alphaHex = alphaHex.length === 1 ? '0' + alphaHex : alphaHex;
    }

    return '#' + redHex + greenHex + blueHex + alphaHex;
}



/**
 * @description 加深颜色值
 * @param {String} color 颜色值字符串
 * @param {Number} level 加深的程度，限0-1之间
 * @returns {String} 返回处理后的颜色值
 */
export function getDarkColor(color: string, level: number) {
    let rgbList = color.match(/rgba?\((\d+),\s*(\d+),\s*(\d+)(?:,\s*([\d.]+))?\)/);
    if (!rgbList) {
        return ElMessage.warning("输入错误的颜色值");
    } else {
        //删除第一个元素
        rgbList.shift();
    }
    for (let i = 0; i < 3; i++) {
        rgbList[i] = Math.round(20.5 * level + Number(rgbList[i]) * (1 - level)).toString();
    }
    return rgbaToHex(parseInt(rgbList[0]), parseInt(rgbList[1]), parseInt(rgbList[2]), parseInt(rgbList[3]));
}

/**
 * @description 变浅颜色值
 * @param {String} color 颜色值字符串
 * @param {Number} level 加深的程度，限0-1之间
 * @returns {String} 返回处理后的颜色值
 */
export function getLightColor(color: string, level: number) {
    let rgbList = color.match(/rgba?\((\d+),\s*(\d+),\s*(\d+)(?:,\s*([\d.]+))?\)/);
    
    if (!rgbList) {
        return ElMessage.warning("输入错误的颜色值");
    } else {
        //删除第一个元素
        rgbList.shift();
    }
    for (let i = 0; i < 3; i++) {
        rgbList[i] = Math.round(255 * level + Number(rgbList[i]) * (1 - level)).toString();
    }
    return rgbaToHex(parseInt(rgbList[0]), parseInt(rgbList[1]), parseInt(rgbList[2]), parseInt(rgbList[3]));
}


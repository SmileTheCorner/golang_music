import { PersistedStateOptions } from "pinia-plugin-persistedstate"

/**
 * @description pinia持久化参数配置
 * @param {String} key 存储到持久化的 name
 * @param {Array} paths 需要持久化的 state name
 * @return persist
 **/

const piniaPersistConfig = (key: string, paths?: string[]) => {
    const persist: PersistedStateOptions = {
        key,
        storage: localStorage,
        paths,
        serializer: { // 指定参数序列化器
            serialize: JSON.stringify,
            deserialize: JSON.parse,
        }
    }
    return persist
}

export default piniaPersistConfig
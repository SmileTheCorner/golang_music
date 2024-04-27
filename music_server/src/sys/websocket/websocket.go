package websocket

import (
	"fmt"
	"github.com/golang-module/carbon"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

// 基础的连接
type WebsocketClientBase struct {
	//连接的唯一标识
	ID string
	//连接
	Conn *websocket.Conn
	//最后心跳时间
	LastHeartbeat int64
	//绑定的UID
	BindUid string
	//连接的组
	JoinGroup []string
}

type WebSocketUserBase struct {
	//用户的唯一标识
	Uid string
	//连接的ID
	ClientID []string
}

// 连接的用户组
type WebSocketGroupBase struct {
	//组的唯一标识
	ClientID []string
}

var GatewayClients, GatewayUser, GatewayGroup sync.Map

/**
 * @Description: 客户端心跳检测，超时将断开连接，减少服务器的开销
 */
func clientHeartbeatCheck(clientID string) {
	for {
		time.Sleep(time.Second * 10)
		clientInterface, exist := GatewayClients.Load(clientID)
		//如果不存在，说明已经断开连接，退出
		if !exist {
			break
		}
		client, _ := clientInterface.(*WebsocketClientBase)
		//如果最后心跳时间小于当前时间10秒，说明已经超时，断开连接
		if carbon.Now().Timestamp()-client.LastHeartbeat > int64(HeartbeatTime) {
			fmt.Println("Client", clientID, "heartbeat timeout")
			//断开连接
			client.Conn.Close()
			//删除连接
			GatewayClients.Delete(clientID)
			break
		}
	}
}

/**
 * @description: 客户端断线时自动踢出Uid绑定列表
 * @param {string} clientID
 * @param {string} uid
 */
func clientUnBindUid(clientID string, uid string) {

	value, ok := GatewayUser.Load(uid)

	if ok {

		users := value.(*WebSocketUserBase)

		for k, v := range users.ClientID {

			if v == clientID {
				users.ClientID = append(users.ClientID[:k], users.ClientID[k+1:]...)
			}
		}

		if len(users.ClientID) == 0 {
			GatewayUser.Delete(uid)
		}

	}
}

/**
 * @description: 客户端断线时自动踢出已加入的群组
 * @param {string} clientID
 */
func clientLeaveGroup(clientID string) {
	// 使用 Load 方法获取值
	value, ok := GatewayClients.Load(clientID)
	if !ok {
		// 如果没有找到对应的值，处理相应的逻辑
		return
	}

	client := value.(*WebsocketClientBase)

	// 遍历 JoinGroup
	for _, v := range client.JoinGroup {
		// 使用 Load 方法获取值
		groupValue, groupOK := GatewayGroup.Load(v)
		if !groupOK {
			// 如果没有找到对应的值，处理相应的逻辑
			continue
		}

		group := groupValue.(*WebSocketGroupBase)

		// 在群组中找到对应的 clientID，并删除
		for j, id := range group.ClientID {
			if id == clientID {
				copy(group.ClientID[j:], group.ClientID[j+1:])
				group.ClientID = group.ClientID[:len(group.ClientID)-1]
				// 如果群组中没有成员了，删除群组
				if len(group.ClientID) == 0 {
					GatewayGroup.Delete(v)
				}
				break
			}
		}
	}
}

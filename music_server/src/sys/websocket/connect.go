package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"runtime/debug"
	"time"
)

var (
	upGrader = websocket.Upgrader{
		// 设置消息接收缓冲区大小（byte），如果这个值设置得太小，可能会导致服务端在读取客户端发送的大型消息时遇到问题
		ReadBufferSize: 2048,
		// 设置消息发送缓冲区大小（byte），如果这个值设置得太小，可能会导致服务端在发送大型消息时遇到问题
		WriteBufferSize: 2048,
		// 消息包启用压缩
		EnableCompression: true,
		// ws握手超时时间
		HandshakeTimeout: 10 * time.Second,
		// ws握手过程中允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 设置心跳检测间隔时长（秒）
	HeartbeatTime = 10
)

/**
 * @description: 初始化客户端连接
 * @param {*websocket.Conn} conn
 */
func handleClientInit(conn *websocket.Conn) string {
	clientID := uuid.New().String()
	// 初始化客户端连接
	client := &WebsocketClientBase{
		ID:            clientID,
		Conn:          conn,
		LastHeartbeat: carbon.Now().Timestamp(),
	}
	// 使用 Store 方法存储值
	GatewayClients.Store(clientID, client)
	if err := conn.WriteMessage(2048, []byte(clientID)); err != nil {
		handleClientDisconnect(clientID)
		return ""
	}

	return clientID
}

/**
 * @description: 主动关闭客户端连接
 * @param {string} clientID
 */
func handleClientDisconnect(clientID string) {

	// 使用 Load 和 Delete 方法，不需要额外的锁定操作
	v, ok := GatewayClients.Load(clientID)
	if ok {

		client := v.(*WebsocketClientBase)

		if client.BindUid != "" {
			clientUnBindUid(clientID, client.BindUid)
		}

		if len(client.JoinGroup) > 0 {
			clientLeaveGroup(clientID)
		}
		GatewayClients.Delete(clientID)
	}
}

/**
 * @description: 向客户端回复心跳消息
 * @param {*websocket.Conn} conn
 * @param {string} clientID
 * @param {int} messageType
 * @param {[]byte} message
 */
func handleClientMessage(conn *websocket.Conn, clientID string, messageType int, message []byte) {

	// 使用 Load 方法获取值
	v, ok := GatewayClients.Load(clientID)
	if !ok {
		// 如果没有找到对应的值，处理相应的逻辑
		handleClientDisconnect(clientID)
		return
	}

	client := v.(*WebsocketClientBase)

	if messageType == 2048 && string(message) == "ping" {

		if err := conn.WriteMessage(2048, []byte("pong")); err != nil {
			handleClientDisconnect(clientID)
			return
		}

		GatewayClients.Store(clientID, &WebsocketClientBase{
			ID:            clientID,
			Conn:          conn,
			LastHeartbeat: carbon.Now().Timestamp(),
			BindUid:       client.BindUid,
			JoinGroup:     client.JoinGroup,
		})
	}
}

func WsServer(c *gin.Context) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("WsServer panic: %v\n", err)
			debug.PrintStack()
		}
	}()

	// 将 HTTP 连接升级为 WebSocket 连接
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}

	defer conn.Close()

	// 客户端唯一身份标识
	clientID := handleClientInit(conn)

	// 发送客户端唯一标识 ID
	if clientID == "" {
		return
	}

	go clientHeartbeatCheck(clientID)

	for {
		// 读取客户端发送过来的消息
		messageType, message, err := conn.ReadMessage()

		// 当收到err时则标识客户端连接出现异常，如断线
		if err != nil {
			handleClientDisconnect(clientID)
		} else {
			handleClientMessage(conn, clientID, messageType, message)
		}
	}

}

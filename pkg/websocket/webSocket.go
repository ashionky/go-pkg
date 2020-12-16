/**
 * @Author pibing
 * @create 2020/8/4 3:34 PM
 */

package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/unknwon/com"
	"net/http"
	"sync"
	"time"
)

func init() {
	go broadcaster()
}

type Client struct {
	conn    *websocket.Conn // 用户websocket连接
	user_id int             // 用于区分用户的socket
}

type Message struct {
	Content string `json:"content"` //消息内容，可以是对象
}

var clients = sync.Map{}       // 用户组映射  user -client
var MesChannel = make(chan Message, 500) // 消息通道

var join = make(chan Client, 100)  // 用户加入通道
var leave = make(chan Client, 100) // 用户退出通道

//websocket连接方法 u_id代表用户标识
func WS(c *gin.Context) {

	user_id := com.StrTo(c.Query("u_id")).MustInt()
	if user_id == 0 {
		c.Redirect(302, "/")
		return
	}
	// 检验http头中upgrader属性，若为websocket，则将http协议升级为websocket协议
	//upGrader := websocket.Upgrader{
	//	// cross origin domain
	//	CheckOrigin: func(r *http.Request) bool {
	//		return true
	//	},
	//	// 处理 Sec-WebSocket-Protocol Header
	//	Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	//}
	//conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Writer, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		fmt.Println("Cannot setup WebSocket connection:", err)
		return
	}

	 client := Client{
		 conn:    conn,
		 user_id: user_id,
	 }

	// 如果用户列表中没有该用户
	if _, ok := clients.Load(user_id); !ok {
		clients.Store(user_id,client) // 将用户加入map
		join <- client
	}

	// 当函数返回时，将该用户加入退出通道，并断开用户连接
	defer func() {
		    clients.Delete(client.user_id)
			leave <- client //离开
			client.conn.Close()
	}()

	for {
		// 读取消息。如果连接断开，则会返回错误
		_, msgStr, err := client.conn.ReadMessage()

		// 如果返回错误，就退出循环
		if err != nil {
			break
		}

		if err := client.conn.WriteMessage(websocket.TextMessage, msgStr); err != nil {
			fmt.Println("heartcheck userid:", client.user_id)
			break
		}

		// 如果没有错误，则把用户发送的信息放入message通道中,
		msg := Message{
			Content: string(msgStr),
		}
		MesChannel <- msg
	}
}

/*前端send的消息格式,可自行修改：
{
   content:"new"
}
*/
func broadcaster() {
	for {

		select {
		// 消息通道中有消息则执行，否则堵塞
		case msg := <-MesChannel:
			content := msg.Content
			//广播给在线的用户
			clients.Range(func(key, value interface{}) bool {
				if client, ok := clients.Load(key); ok {
					if c,ok:=client.(Client);ok{
						if err := c.conn.WriteMessage(websocket.TextMessage, []byte(content)); err != nil {
							fmt.Println("push fail userid:", key)
						}else {
							fmt.Println("push success userid:", key)
						}
					}
				}
				return true
			})

			// 有用户加入
		case client := <-join:
			fmt.Println("userid:", client.user_id, "---join，time：", time.Now().Format("2006-01-02 15:04:05"))
			// 有用户退出
		case client := <-leave:
			fmt.Println("userid:", client.user_id, "---leave，time：", time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

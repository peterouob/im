package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
	"net"
	"net/http"
	"strconv"
	"sync"
)

//消息
type Message struct {
	gorm.Model
	FromId   int64  //發送者
	TargetId int64  //接收者
	Type     int    //消息類席 群聊 私聊 廣播
	Media    string //消息類型 文字 圖片 音頻
	Content  string //消息內容
	Pic      string
	Url      string
	Desc     string //描述
	Amount   int    //其他數字統計
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte   //信息管道
	GroupSets set.Interface //集合
}

//用戶集合
var ClientMap map[int64]*Node = make(map[int64]*Node, 0)

//讀寫鎖防安全性問題
var rwLock sync.RWMutex

//發送者ID，接受者ID，消息類型，發送的內容，發送類型
func Chat(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	//1. 獲取參數並校驗token等合法性
	Id := query.Get("userId")
	userId, _ := strconv.ParseInt(Id, 10, 64)
	//targetId := query.Get("targetId")
	//msgType := query.Get("msgType")
	//context := query.Get("context")
	isValid := true // checkToken(){} 待處理...
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isValid
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket error", err)
	}

	//2. 獲取Conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//3. 用戶關係
	//4. 綁定userId和node 並加鎖
	rwLock.Lock()
	ClientMap[userId] = node
	rwLock.Unlock()
	//5. 完成發送邏輯
	go sendProc(node)
	//6. 完成接受蘿雞
	go recvProc(node)
	sendMsg(userId, []byte("親，歡迎進入聊天室"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadcast(data)
		fmt.Println("[ws] <<<<<< ", data)
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadcast(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpSendProc()
	go udpRecvProc()
}

//完成udp數據發送協程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(con)

	for {
		select {
		case data := <-udpsendChan:
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

//完成udp數據接受協程
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero, //全部都可以連結
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(con)

	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n])
	}
}

//後端調度邏輯
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: //私信
		sendMsg(msg.TargetId, data) //發送給誰？所以使用target而不是formID
	case 2: //群聊
		//sendGroupMsg()
	case 3: //廣播
		//sendAllMsg()
	case 4:
		//
	}
}

func sendMsg(userId int64, msg []byte) {
	rwLock.RLock()
	node, ok := ClientMap[userId]
	rwLock.RUnlock() //Rlock對應RUnlock否則發生iowait問題
	if ok {
		node.DataQueue <- msg
	}
}

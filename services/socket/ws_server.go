package socket

import (
	"RPJ_Overseas_Exim/go_mod_home/db/models"
	"RPJ_Overseas_Exim/go_mod_home/utils"
	"bytes"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

const (

    writeWait = 10 *time.Second

    pongWait = 60 * time.Second

    pingPeriod = pongWait *9/10

    maxMessageSize = 512
)

var (
    newLine = []byte{'\n'}
    space = []byte{' '}
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

type ChatParticipant struct {
    hub *Hub
    conn *websocket.Conn
    messages chan []byte
    userId string
    chatId string
    role string
}

func (c *ChatParticipant) readPump() {
    defer func(){
        c.hub.unregister <- c
        c.conn.Close()
    }()

    c.conn.SetReadLimit(maxMessageSize)
    c.conn.SetReadDeadline(time.Now().Add(pongWait))
    c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil})
    for {
        _, message, err := c.conn.ReadMessage()

        if err != nil{
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure){
                log.Printf("error: %v", err)
            }
            log.Println("Error while reading message")
            break
        }

        message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))

        notification := NewNotification("reload", string(message), c.chatId)
        c.hub.notification <- notification
        c.hub.broadcast <- models.NewMessage(c.chatId, c.userId, message)
    }
    log.Println("read pump dead")
}

func (c *ChatParticipant) writePump(){
    ticker := time.NewTicker(pingPeriod)
    defer func(){
        ticker.Stop()
        c.conn.Close()
    }()
    for {
        select {
        case message, ok := <-c.messages:
            c.conn.SetWriteDeadline(time.Now().Add(writeWait))
            if !ok {
                c.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }

            w, err := c.conn.NextWriter(websocket.TextMessage)
            if err != nil{
                return
            }

            w.Write(message)

            n:= len(c.messages)
            for i:=0;i<n;i++{
                w.Write(newLine)
                w.Write(<-c.messages)
            }

            if err := w.Close();err != nil{
                return
            }

        case <- ticker.C:
            c.conn.SetWriteDeadline(time.Now().Add(writeWait))
            if err := c.conn.WriteMessage(websocket.PingMessage, nil); err !=nil{
                return
            }
        }
    }
}

func ServeWs(chatId string, userId string, hub *Hub, c echo.Context){
    conn, err := upgrader.Upgrade(c.Response() , c.Request(), nil)

    if err != nil{
        log.Println(err)
        return
    }

    cp := &ChatParticipant{chatId:chatId, userId:userId, hub:hub, conn:conn, messages: make(chan []byte, 256), role:"user"}
    // log.Println("Participant chat id: ", cp.chatId)

    cp.hub.register <- cp

    go cp.writePump()
    go cp.readPump()
}
 
func ServeAdminWs(userId string, hub *Hub, c echo.Context) (*ChatParticipant, error){
    conn, err := upgrader.Upgrade(c.Response() , c.Request(), nil)

    var cp *ChatParticipant
    if err != nil{
        log.Println(err)
        return cp, &utils.HTTPException{Message:"Could not upgrade socket connection"}
    }

    cp = &ChatParticipant{chatId: "adminTemp", userId: userId, hub: hub, conn: conn, messages: make(chan []byte, 256), role: "admin"}

    cp.hub.register <- cp

    go cp.writePump()
    go cp.readPump()
    return cp,nil
}

func SwitchChats(cp *ChatParticipant, chatId string, hub *Hub){
    chat, ok := hub.chats[chatId]

    if ok {
        adminFound := false
        for k := range chat.cp{
            if k.role == "admin"{
                k.chatId = chatId
                adminFound = true
            }
        }

        if !adminFound{
            cp.chatId = chatId
            chat.cp[cp] = true
        }

    }else{
        cp.chatId = chatId
        hub.chats[cp.chatId] = NewChat(cp)
    }
}

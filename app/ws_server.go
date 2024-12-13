package main

import (
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
    send chan []byte
    userId string
    chatId string
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
            break
        }

        message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))

        c.hub.broadcast <- message
    }
}

func (c *ChatParticipant) writePump(){
    ticker := time.NewTicker(pingPeriod)
    defer func(){
        ticker.Stop()
        c.conn.Close()
    }()
    for {
        select {
        case message, ok := <-c.send:
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

            n:= len(c.send)
            for i:=0;i<n;i++{
                w.Write(newLine)
                w.Write(<-c.send)
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

func serveWs(chatId string, hub *Hub, c echo.Context){
    conn, err := upgrader.Upgrade(c.Response() , c.Request(), nil)
    if err != nil{
        log.Println(err)
        return
    }

    client := &ChatParticipant{chatId:chatId, hub:hub, conn:conn, send: make(chan []byte, 256)}
    client.hub.register <- client
    go client.writePump()
    go client.readPump()
}


package main

type Message struct {
    chatId string
    msg []byte
}

type Hub struct {
    broadcast chan Message
    register chan *ChatParticipant
    unregister chan *ChatParticipant
    chats map[string] *Chat
}

type Chat struct{
    cp map[*ChatParticipant]bool
}

func newChat(cp *ChatParticipant) *Chat{
    return &Chat{
        cp: map[*ChatParticipant]bool{cp:true},
    }
}

func newHub() *Hub {
    return &Hub{
        broadcast: make(chan Message),
        register: make(chan *ChatParticipant),
        unregister: make(chan *ChatParticipant),
        chats: make(map[string] *Chat),
    }
}

func (h*Hub) run(){
    for {
        select {
        case cp := <-h.register:
            chat, ok := h.chats[cp.chatId] 
            if !ok {
                chat = newChat(cp)
            }else{
                chat.cp[cp]  =  true
            }

        case cp := <-h.unregister:
            if _,ok := h.chats[cp.chatId].cp[cp];ok{
                delete(h.chats[cp.chatId].cp, cp)
                close(cp.send)
            }
        case message := <-h.broadcast:
            for cp := range h.chats[message.chatId]{
                select {
                case cp.send <- message:
                default:
                    close(cp.send)
                    delete(h.cp, cp)
                }
            }
        }
    }
}

package socket


type Message struct {
    chatId,
    userId string
    msg []byte
}

func NewMessage(chatId , userId string, msg []byte) *Message{
    return &Message{
        chatId,
        userId, 
        msg,
    }
}

type Hub struct {
    broadcast chan *Message
    register chan *ChatParticipant
    unregister chan *ChatParticipant
    chats map[string] *Chat
}

type Chat struct{
    Id string
    cp map[*ChatParticipant]bool
}

func NewChat(cp *ChatParticipant) *Chat{
    return &Chat{
        Id:cp.chatId,
        cp:map[*ChatParticipant]bool{cp:true},
    }
}

func NewHub() *Hub {
    return &Hub{
        broadcast: make(chan *Message),
        register: make(chan *ChatParticipant),
        unregister: make(chan *ChatParticipant),
        chats: make(map[string] *Chat),
    }
}

func (h*Hub) Run(){
    for {
        select {
        case cp := <-h.register:
            // log.Println("registering ", cp.userId)
            _, ok := h.chats[cp.chatId] 
            if !ok {
                h.chats[cp.chatId] = NewChat(cp)
            }else{
                h.chats[cp.chatId].cp[cp]  =  true
            }

        case cp := <-h.unregister:
            // log.Println("unregistering ", cp.userId)
            if _, ok := h.chats[cp.chatId].cp[cp] ; ok{
                _, ok := h.chats[cp.chatId].cp[cp]

                if ok {
                    delete(h.chats[cp.chatId].cp,cp)
                }
                close(cp.messages)
            }

        case message := <-h.broadcast:
            // log.Println("Message: ", message.userId, message.chatId)
            // log.Println("Message ChatId: ", message.chatId)
            participants := h.chats[message.chatId].cp

            // log.Println("participant list")
            for cp := range participants{
                // log.Println(cp.userId)
                select {
                case cp.messages <- message.msg:
                default:
                    close(cp.messages)
                    delete(h.chats[cp.chatId].cp,cp)
                }
            }
            // log.Println("end")
        }
    }
}

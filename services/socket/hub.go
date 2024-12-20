package socket

import (
	"RPJ_Overseas_Exim/go_mod_home/db/models"
	"encoding/json"
	"log"
	"time"

	"gorm.io/gorm"
)

type Hub struct {
	notification         chan *Notification
	broadcast            chan *models.Message
	register, unregister chan *ChatParticipant
	chats                map[string]*Chat
	dbConn               *gorm.DB
}

type Notification struct {
	Event   string `json:"event"`
	Message string `json:"message"`
	ChatId  string
    UserId string
}

func NewNotification(event, message, chatId, userId string) *Notification {
	return &Notification{
		event,
		message,
		chatId,
        userId,
	}
}

type Chat struct {
	Id string
	cp map[*ChatParticipant]bool
}

func NewChat(cp *ChatParticipant) *Chat {
	return &Chat{
		Id: cp.chatId,
		cp: map[*ChatParticipant]bool{cp: true},
	}
}

func NewHub(dbConn *gorm.DB) *Hub {
	return &Hub{
		broadcast:    make(chan *models.Message),
		register:     make(chan *ChatParticipant),
		unregister:   make(chan *ChatParticipant),
		notification: make(chan *Notification),
		chats:        make(map[string]*Chat),
		dbConn:       dbConn,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cp := <-h.register:
			// log.Println("registering ", cp.userId)
			_, ok := h.chats[cp.chatId]

			h.dbConn.Model(&models.SocketUser{}).Where("id = ?", cp.userId).Update("Online", true)

			if !ok {
				h.chats[cp.chatId] = NewChat(cp)
			} else {
				h.chats[cp.chatId].cp[cp] = true
			}

		case cp := <-h.unregister:
			// log.Println("unregistering ", cp.userId)
			h.dbConn.Model(&models.SocketUser{}).Where("id = ?", cp.userId).Update("Online", false)
			h.dbConn.Model(&models.Participant{}).Where("socket_user_id = ?", cp.userId).Update("LastSeen", time.Now())

            if cp.role=="admin"{
                delete(h.chats["adminTemp"].cp, cp)
            }
			if _, ok := h.chats[cp.chatId].cp[cp]; ok {
				_, ok := h.chats[cp.chatId].cp[cp]

				if ok {
					delete(h.chats[cp.chatId].cp, cp)
				}

				close(cp.messages)
			}

		case notif := <-h.notification:
			log.Println("Notification: ", notif)
			participants := h.chats[notif.ChatId].cp
			adminMap := h.chats["adminTemp"].cp
			adminFound := false
			for cp := range participants {
				// log.Println(cp.userId)
				if cp.role == "admin" {
					adminFound = true
				}
                if cp.userId==notif.UserId{
                    continue
                }
				notifMsg, err := json.Marshal(notif)
				// log.Println("notif msg: ", notifMsg)
				if err != nil {
					log.Fatalln("Could not marshal the notification to a json")
				} else {
					select {
					case cp.messages <- notifMsg:
					default:
						close(cp.messages)
						delete(h.chats[cp.chatId].cp, cp)
					}
				}
			}

			log.Println(adminFound)

			if !adminFound {
				for cp := range adminMap {
					// log.Println(cp.userId)
					if cp.role == "admin" {
						adminFound = true
					}
					notifMsg, err := json.Marshal(notif)
					// log.Println("notif msg: ", notifMsg)
					if err != nil {
						log.Fatalln("Could not marshal the notification to a json")
					} else {
						select {
						case cp.messages <- notifMsg:
						default:
							close(cp.messages)
							delete(h.chats[cp.chatId].cp, cp)
						}
					}
				}
			}
		case message := <-h.broadcast:
			// log.Println("Message: ", message.userId, message.chatId)
			// log.Println("Message ChatId: ", message.chatId)
			participants := h.chats[message.ChatId].cp

			dbMessage := models.NewMessage(message.ChatId, message.SocketUserId, []byte(message.Message))
			err := h.dbConn.Create(&dbMessage)

			if err.Error != nil {
				log.Fatalln("Could not create the message :[")
			} else {
				// log.Println("participant list")

				for cp := range participants {
					// log.Println(cp.userId)
					select {
					case cp.messages <- []byte(message.Message):
					default:
						close(cp.messages)
						delete(h.chats[cp.chatId].cp, cp)
					}
				}
				// log.Println("end")
			}
		}
	}
}

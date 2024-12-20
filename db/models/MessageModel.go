package models

import (
	"RPJ_Overseas_Exim/go_mod_home/utils"

	"gorm.io/gorm"
)

type Message struct{
    gorm.Model
    Id string `gorm:"primaryKey"`
    SocketUserId string
    ChatId string
    Message string
}

func NewMessage(chatId , userId string, msg []byte) *Message{
    return &Message{
        Id:*utils.GenNanoid(),
        SocketUserId: userId,
        ChatId: chatId,
        Message: string(msg),
    }
}

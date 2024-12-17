package models

import (
	"RPJ_Overseas_Exim/go_mod_home/utils"
	"time"

	"gorm.io/gorm"
)

type Participant struct {
    gorm.Model
    Id string `gorm:"primaryKey"`
    SocketUserId string
    ChatId string
    lastSeen time.Time
}

func NewParticipant(userId, chatId string) *Participant{
    return &Participant{
        Id: *utils.GenNanoid(),
        SocketUserId: userId,
        ChatId: chatId,
    }
}

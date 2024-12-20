package models

import (
	"RPJ_Overseas_Exim/go_mod_home/utils"
	"time"

	"gorm.io/gorm"
)

type Chat struct {
    gorm.Model
    Id string `gorm:"primaryKey"`
    Participant []Participant
    Message []Message
    LastMessageTime time.Time
}

func NewChat() *Chat{
    return &Chat{
        Id: *utils.GenNanoid(),
    }
}

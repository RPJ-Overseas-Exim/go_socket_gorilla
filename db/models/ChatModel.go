package models

import (
	"RPJ_Overseas_Exim/go_mod_home/utils"

	"gorm.io/gorm"
)

type Chat struct {
    gorm.Model
    Id string `gorm:"primaryKey"`;
    Participant []Participant;
    Message []Message;
}

func NewChat() *Chat{
    return &Chat{
        Id: *utils.GenNanoid(),
    }
}

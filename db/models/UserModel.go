package models

import (
	"RPJ_Overseas_Exim/go_mod_home/utils"

	"gorm.io/gorm"
)

type SocketUser struct {
    gorm.Model
    Id string `gorm:"primaryKey"`
    Email string `gorm:"uniqueIndex"`
    Role string
    Participant []Participant
    Message []Message
    Online bool
}


func NewSocketuser (email string) *SocketUser{
    return &SocketUser{
        Id:*utils.GenNanoid(),
        Email:email,
        Role: "user",
    }
}

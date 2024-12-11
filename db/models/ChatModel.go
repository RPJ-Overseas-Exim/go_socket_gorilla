package models

import (
	"time"

	"gorm.io/gorm"
)

type Participant struct {
    gorm.Model
    Id string 
    UserId string
    ChatId string
    lastSeen time.Time
}

type Chat struct {
    gorm.Model
    Id string;
    Participant []Participant;
    Message []Message;
}

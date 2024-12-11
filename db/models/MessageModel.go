package models

import "gorm.io/gorm"

type Message struct{
    gorm.Model
    Id string
    UserId string
    ChatId string
    Message string
    SenderId int
}

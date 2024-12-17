package models

import "gorm.io/gorm"

type Message struct{
    gorm.Model
    Id string `gorm:"primaryKey"`
    SocketUserId string
    ChatId string
    Message string
    SenderId int
}

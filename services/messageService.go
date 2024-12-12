package services

import (
	"RPJ_Overseas_Exim/go_mod_home/db/models"
	"log"

	"gorm.io/gorm"
)

type MessageService struct {
	dbConn *gorm.DB
}

func (ms *MessageService) GetMessages(chatId string) *[]models.Message {
	var msgs []models.Message
	ms.dbConn.
		InnerJoins("Chat", ms.dbConn.Where(&models.Chat{Id: chatId})).
		Find(&msgs)
	return &msgs
}

func (ms *MessageService) SendMessage(chatId, userId, content string) error {
    log.Println(chatId, userId, content)
    return nil
}

func NewMessageService(dbConn *gorm.DB) *MessageService {
    return &MessageService{dbConn}
}

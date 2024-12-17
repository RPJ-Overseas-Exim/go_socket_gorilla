package services

import (
	"RPJ_Overseas_Exim/go_mod_home/db/models"
	"log"

	"gorm.io/gorm"
)

type ChatService struct {
	dbConn *gorm.DB
}

func (cs *ChatService) GetAllChats() *[]models.Chat {
	var chats []models.Chat

	cs.dbConn.
        Model(&models.Chat{}).
		InnerJoins("inner join participants on chats.id = participants.chat_id").
		InnerJoins("inner join socket_users on participants.socket_user_id = socket_users.id").
        Group("chats.id").
		Find(&chats)

	return &chats
}

func (cs *ChatService) GetChatAndUserId(email string) (string, string) {

	var participant models.Participant

	log.Println("Eamil: ", email)
	cs.dbConn.Model(&models.Participant{}).
		InnerJoins("inner join socket_users on participants.socket_user_id = socket_users.id and socket_users.email = ?", email).
		First(&participant)

	log.Println("participant: ", participant)
	if participant.ChatId != "" {
		return participant.ChatId, participant.SocketUserId
	} else {
		user := models.NewSocketuser(email)
		cs.dbConn.Create(&user)

		var admin models.SocketUser
		cs.dbConn.First(&admin, "role = ?", "admin")

		insertedChat := *models.NewChat()
		cs.dbConn.Create(&insertedChat)

		participant := []models.Participant{*models.NewParticipant(user.Id, insertedChat.Id), *models.NewParticipant(admin.Id, insertedChat.Id)}
		cs.dbConn.CreateInBatches(&participant, 2)

		return insertedChat.Id, user.Id
	}
}

func NewChatService(dbConn *gorm.DB) *ChatService {
	return &ChatService{
		dbConn,
	}
}

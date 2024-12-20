package services

import (
	"RPJ_Overseas_Exim/go_mod_home/db"
	"RPJ_Overseas_Exim/go_mod_home/db/models"
	"log"
	"time"

	// "log"

	"gorm.io/gorm"
)

type ChatService struct {
	dbConn *gorm.DB
}

func (cs *ChatService) GetAllChats() (*[]db.ResultsType, *map[string]time.Time) {

    var results []db.ResultsType;

	cs.dbConn.
        Table("socket_users").
		InnerJoins("inner join participants on participants.socket_user_id = socket_users.id and socket_users.role <> ?", "admin").
        InnerJoins("inner join chats on chats.id = participants.chat_id ").
        Order("participants.chat_id").
		Scan(&results)

    var adminChats []models.Participant

    cs.dbConn.
        Model(&models.Participant{}).
        Select("chat_id", "last_seen").
        InnerJoins("inner join socket_users on socket_users.id = participants.socket_user_id" ).
        Find(&adminChats, "socket_users.role = ?", "admin")

    seenMap  := make(map[string]time.Time)

    for _,v := range adminChats{
        seenMap[v.ChatId] = v.LastSeen
    }

    // log.Println("participants: ", results)

	return &results, &seenMap
}

func (cs *ChatService) GetChatAndUserId(email string) (string, string) {

	var participant models.Participant

	// log.Println("Email: ", email)
	cs.dbConn.Model(&models.Participant{}).
		InnerJoins("inner join socket_users on participants.socket_user_id = socket_users.id and socket_users.email = ?", email).
		First(&participant)

	// log.Println("participant: ", participant)
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

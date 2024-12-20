package db

import (
	"RPJ_Overseas_Exim/go_mod_home/db/models"
	"log"
	"time"

	"github.com/aidarkhanov/nanoid"
	"gorm.io/gorm"
)

func GenerateNanoid() string {
	alphabets := nanoid.DefaultAlphabet
	id, err := nanoid.Generate(alphabets, 12)
	if err != nil {
		log.Fatal("Could not generate nanoid")
	}
	return id
}

// func populateTempData(db *gorm.DB) {
// 	zolpidemId := GenerateNanoid()
// 	alprazolamId := GenerateNanoid()
// 	clonazepamId := GenerateNanoid()
// 	lorazepamId := GenerateNanoid()
// 	priceQty, err := models.NewPriceQty(zolpidemId, []string{GenerateNanoid(), GenerateNanoid()}, []int16{245, 360}, []int16{90, 180})
// 	priceQty1, _ := models.NewPriceQty(zolpidemId, []string{GenerateNanoid(), GenerateNanoid()}, []int16{245, 360}, []int16{90, 180})
// 	priceQty2, _ := models.NewPriceQty(zolpidemId, []string{GenerateNanoid(), GenerateNanoid()}, []int16{245, 360}, []int16{90, 180})
// 	priceQty3, _ := models.NewPriceQty(zolpidemId, []string{GenerateNanoid(), GenerateNanoid()}, []int16{245, 360}, []int16{90, 180})
//
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
//
// 	product := []*models.Product{
// 		models.NewProduct(zolpidemId, "zolpidem", priceQty),
// 		models.NewProduct(alprazolamId, "alprazolam", priceQty1),
// 		models.NewProduct(clonazepamId, "clonazepam", priceQty2),
// 		models.NewProduct(lorazepamId, "lorazepam", priceQty3),
// 	}
//
// 	results := db.Create(product)
//
// 	if results.Error != nil {
// 		log.Fatal("Product initialization failed")
// 	} else {
// 		log.Println("Products initialized successfully")
// 	}
// }

func migrateDb(db *gorm.DB){
     log.Println("Running migration")
     db.AutoMigrate(
         models.SocketUser{},
         models.Chat{},
         models.Participant{},
         models.Message{},
     )
     Session := db.Session(&gorm.Session{PrepareStmt: true})
     if Session != nil {
         log.Println("Migration successful")
     }
    // db.AutoMigrate(&models.Chat{})
    // db.AutoMigrate(&models.SocketUser{})
    // db.AutoMigrate(&models.Participant{})
    // db.AutoMigrate(&models.Message{})
}

type ResultsType struct {
    Email, 
    ChatId string
    Online bool
    LastSeen,
    LastMessageTime time.Time
}


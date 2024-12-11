package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Id string
    Name string
    Email string
    Role string
    Participant []Participant
    Message []Message
    Online bool
}

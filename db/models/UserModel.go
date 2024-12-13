package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Id string
    Name string
    Email string
    Number int
    Role string
    Participant []Participant
    Message []Message
    Online bool
}

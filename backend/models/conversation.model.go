package models

import (
	"gorm.io/gorm"
)


type Conversation struct {
    gorm.Model
    UserIDs       []uint        `gorm:"type:json"`
    Messages      []Message     `gorm:"foreignKey:ConversationID"`
}
package models

import (
	"time"
)

type Sessions struct {
	ID        string              `bson:"_id,omitempty" json:"id"`
	UserId 	  string              `bson:"userId" json:"userId" binding:"required"`
	ValidTill time.Time  `bson:"validTill" json:"validTill" binding:"required"`
}
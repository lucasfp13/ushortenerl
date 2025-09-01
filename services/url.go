package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Original  string             `json:"original" bson:"original" validate:"required"`
	Short     string             `json:"short" bson:"short" validate:"required"`
	Clicks    int64              `json:"clicks" bson:"clicks"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	ExpiresAt time.Time          `json:"expires_at,omitempty" bson:"expires_at,omitempty"`
}

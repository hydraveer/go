package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" validate:"required, min=2, max=100"`
	LastName     *string            `json:"last_name" validate:"required, min=2, max=100"`
	Password     *string            `json:"password" validate:"required, min=6, max=20"`
	Email        *string            `json:"email" validate:"required, min=4, max=20"`
	Phone        *string            `json:"phone" validate:"required, min=10"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	RefreshToken *string            `json:"refresh_token" validate:"required"`
	CraetedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
	UserId       string             `json:"user_id" validate:"required, min=2, max=100"`
}


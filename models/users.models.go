package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User is the model that governs all notes objects retrived or inserted into the DB
type User struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	First_Name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_Name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"Password" validate:"required,min=6""`
	Email         *string            `json:"email" validate:"email,required"`
	Phone         *string            `json:"phone" validate:"required"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	Created_At    time.Time          `json:"created_at"`
	Updated_At    time.Time          `json:"updated_at"`
	User_Id       int                `json:"user_id"`
}

type UserTokens struct {
	Token        *string `json:"token"`
	RefreshToken *string `json:"refresh_token"`
}

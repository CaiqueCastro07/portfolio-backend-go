package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID      primitive.ObjectID `bson:"_id"`
	Message string             `bson:"message"`
	Email   string             `bson:"email"`
	Phone   string             `bson:"phone"`
	Name    string             `bson:"name"`
}

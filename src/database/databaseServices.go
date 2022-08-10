package database

import (
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterMessage(task *Message) (bool, string) {

	_, err := MessagesCollection.InsertOne(Ctx, task)

	if err != nil {
		return false, err.Error()
	}

	return true, ""

}

func GetUser(user string) (bool, Task) {

	var document Task

	err := TasksCollection.FindOne(Ctx, bson.D{{Key: "user", Value: user}}).Decode(&document)

	if err != nil {
		return false, document
	}

	return true, document

}

func GetUserByEmail(email string) (bool, Task) {

	var document Task

	err := TasksCollection.FindOne(Ctx, bson.D{{Key: "email", Value: email}}).Decode(&document)

	if err != nil {
		return false, document
	}

	return true, document

}

func CreateUser(user string, email string, password string) (bool, string) {

	var newUser Task

	newUser.User = user
	newUser.Email = email
	newUser.Password = password
	newUser.Tasks = []struct {
		Task     string `bson:"task"`
		Done     bool   `bson:"done"`
		Priority int    `bson:"priority"`
	}{}

	_, err := TasksCollection.InsertOne(Ctx, newUser)

	if err != nil {
		return false, err.Error()
	}

	return true, ""

}

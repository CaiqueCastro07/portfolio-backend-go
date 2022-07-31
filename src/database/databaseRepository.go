package database

type Message struct {
	Message string `bson:"message"`
	Email   string `bson:"email"`
	Phone   string `bson:"phone"`
	Name    string `bson:"name"`
}

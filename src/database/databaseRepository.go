package database

type Message struct {
	Message string `bson:"message"`
	Email   string `bson:"email"`
	Phone   string `bson:"phone"`
	Name    string `bson:"name"`
}

type Task struct {
	User     string `bson:"user"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
	Tasks    []struct {
		Task     string `bson:"task"`
		Done     bool   `bson:"done"`
		Priority int    `bson:"priority"`
	} `bson:"tasks"`
}

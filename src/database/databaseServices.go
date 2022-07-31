package database

func RegisterMessage(task *Message) (bool, string) {

	_, err := Collection.InsertOne(Ctx, task)

	if err != nil {
		return false, err.Error()
	}

	return true, ""

}

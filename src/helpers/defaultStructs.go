package helpers

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data"`
}

type ErrorJson struct {
	Value  string      `json:"value"`
	Struct string      `json:"struct"`
	Field  bool        `json:"field"`
	Data   interface{} `json:"data"`
}

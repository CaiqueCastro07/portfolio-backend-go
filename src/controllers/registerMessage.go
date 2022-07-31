package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/CaiqueCastro07/portfolio-backend-go/src/database"
	"github.com/CaiqueCastro07/portfolio-backend-go/src/helpers"

	"github.com/gin-gonic/gin"
)

func RegisterMessage(c *gin.Context) {

	var newMessage database.Message

	defer c.Request.Body.Close()

	err := c.ShouldBindJSON(&newMessage)

	if err != nil {

		var errMap helpers.ErrorJson
		data, _ := json.Marshal(err)
		json.Unmarshal(data, &errMap)

		c.IndentedJSON(http.StatusBadRequest, helpers.Response{http.StatusBadRequest, err.Error(), true, errMap})
		return
	}

	if emailValidation, errMsg := helpers.ValidateEmail(newMessage.Email); !emailValidation {
		c.IndentedJSON(http.StatusBadRequest, helpers.Response{http.StatusBadRequest, errMsg, true, newMessage.Email})
		return
	}

	if messageValidation, errMsg := helpers.ValidateMessage(newMessage.Message); !messageValidation {
		c.IndentedJSON(http.StatusBadRequest, helpers.Response{http.StatusBadRequest, errMsg, true, "*Mensagem inválida*"})
		return
	}

	if messageRegistered, errMsg := database.RegisterMessage(&newMessage); messageRegistered == false {
		c.IndentedJSON(http.StatusBadRequest, helpers.Response{http.StatusBadRequest, errMsg, true, "Erro ao registrar mensagem no banco de dados."})
		return
	}

	c.IndentedJSON(http.StatusOK, helpers.Response{http.StatusOK, "Mensagem registrada com sucesso.", false, map[string]string{}})
}

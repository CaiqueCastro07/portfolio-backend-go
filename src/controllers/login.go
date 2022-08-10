package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/CaiqueCastro07/portfolio-backend-go/src/database"
	"github.com/CaiqueCastro07/portfolio-backend-go/src/helpers"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	User     string      `json:"user"`
	Password map[int]int `json:"password"`
	R        string      `json:r`
}

func Login(c *gin.Context) {

	var credentials Credentials

	defer c.Request.Body.Close()

	if err := c.ShouldBindJSON(&credentials); err != nil {

		var errMap helpers.ErrorJson
		data, _ := json.Marshal(err)
		json.Unmarshal(data, &errMap)

		c.IndentedJSON(http.StatusBadRequest, helpers.Response{http.StatusBadRequest, err.Error(), true, errMap})
		return
	}

	ok, password := helpers.DecodePassword(credentials.Password, credentials.R)

	if !ok {
		c.IndentedJSON(500, helpers.Response{500, "Erro interno", true, map[string]interface{}{}})
		return
	}

	ok2, data := database.GetUser(credentials.User)

	if !ok2 {
		c.IndentedJSON(http.StatusBadRequest, helpers.Response{http.StatusBadRequest, "Usuário não existe.", true, map[string]string{}})
		return
	}

	if data.Password != password {
		c.IndentedJSON(http.StatusBadRequest, helpers.Response{http.StatusBadRequest, "Senha incorreta.", true, map[string]string{}})
		return
	}

	c.IndentedJSON(http.StatusOK, helpers.Response{http.StatusOK, "Usuário autorizado", false, credentials.User})
}

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CaiqueCastro07/portfolio-backend-go/src/database"
	"github.com/CaiqueCastro07/portfolio-backend-go/src/helpers"

	"github.com/gin-gonic/gin"
)

type RegisterCredentials struct {
	User     string      `json:"user"`
	Password map[int]int `json:"password"`
	Email    string      `json:"email"`
	R        string      `json:"r"`
}

func Register(c *gin.Context) {

	var credentials RegisterCredentials

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

	exists, _ := database.GetUser(credentials.User)

	if exists {
		c.IndentedJSON(http.StatusBadRequest, helpers.Response{http.StatusBadRequest, "Usuário já existe.", true, map[string]string{}})
		return
	}

	exists2, _ := database.GetUserByEmail(credentials.Email)

	if exists2 {
		c.IndentedJSON(http.StatusBadRequest, helpers.Response{http.StatusBadRequest, "O email já está em uso.", true, map[string]string{}})
		return
	}

	if ok, msg := database.CreateUser(credentials.User, credentials.Email, password); !ok {
		fmt.Println(msg)
		c.IndentedJSON(http.StatusBadRequest, helpers.Response{http.StatusBadRequest, "Erro ao registrar usuário.", true, map[string]string{}})
		return
	}

	c.IndentedJSON(http.StatusOK, helpers.Response{http.StatusOK, "Usuário autorizado", false, map[string]string{}})
}

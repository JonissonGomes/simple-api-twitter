package handlers

import (
	"net/http"

	"github.com/JonissonGomes/simple-api-twitter/model"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) SignUp(c echo.Context) (err error) {

	// User connect
	user := &model.User{ID: bson.NewObjectId()}
	if err = c.Bind(user); err != nil {
		return
	}

	// User validation
	if user.Email == "" || user.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Email ou senha não conferem"}
	}

	return
}

package handlers

import (
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

	return
}

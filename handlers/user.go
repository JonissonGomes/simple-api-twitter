package handlers

import (
	"net/http"

	"github.com/JonissonGomes/simple-api-twitter/model"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2"
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

	// Save user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter-api").C("users").Insert(user); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// User connect
	user := &model.User{ID: bson.NewObjectId()}
	if err = c.Bind(user); err != nil {
		return
	}

	// Search user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter-api").C("users").Find(bson.M{"email": user.Email, "password": user.Password}).One(user); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Email ou senha não conferem"}
		}
		return
	}

	return
}

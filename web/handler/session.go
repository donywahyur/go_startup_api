package handler

import (
	"go_startup_api/user"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	userService user.Service
}

func NewSessionHandler(userService user.Service) *SessionHandler {
	return &SessionHandler{userService}
}

func (h *SessionHandler) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "session_new.html", nil)
}

func (h *SessionHandler) Process(c *gin.Context) {
	var input user.FormLoginInput
	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "session_new.html", nil)
		return
	}

	var inputLogin user.LoginInput
	inputLogin.Email = input.Email
	inputLogin.Password = input.Password

	user, err := h.userService.Login(inputLogin)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "session_new.html", nil)
		return
	}

	if user.ID == 0 || user.Role != "admin" {
		c.HTML(http.StatusInternalServerError, "session_new.html", nil)
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Save()

	c.Redirect(http.StatusFound, "/users")
}

func (h *SessionHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}

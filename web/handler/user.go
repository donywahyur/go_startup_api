package handler

import (
	"fmt"
	"go_startup_api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{
		userService,
	}
}

func (h *userHandler) Index(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "user_index.html", gin.H{"users": users})
}

func (h *userHandler) New(c *gin.Context) {

	c.HTML(http.StatusOK, "user_new.html", nil)
}

func (h *userHandler) Create(c *gin.Context) {
	var input user.FormCreateUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusInternalServerError, "user_new.html", input)
		return
	}

	var registerInput user.RegisterUserInput
	registerInput.Name = input.Name
	registerInput.Email = input.Email
	registerInput.Occupation = input.Occupation
	registerInput.Password = input.Password

	_, err = h.userService.RegisterUser(registerInput)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusInternalServerError, "user_new.html", input)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) Edit(c *gin.Context) {
	var input user.FormDetailUserInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", nil)
		return
	}

	userGet, err := h.userService.GetUserByID(input.ID)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", nil)
		return
	}

	var inputUser user.FormUpdateUserInput
	inputUser.ID = input.ID
	inputUser.Name = userGet.Name
	inputUser.Email = userGet.Email
	inputUser.Occupation = userGet.Occupation
	inputUser.Error = nil

	c.HTML(http.StatusOK, "user_edit.html", inputUser)
}

func (h *userHandler) Update(c *gin.Context) {
	var input user.FormDetailUserInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", nil)
		return
	}

	var inputUpdate user.FormUpdateUserInput
	inputUpdate.ID = input.ID
	err = c.ShouldBind(&inputUpdate)
	if err != nil {
		inputUpdate.Error = err
		c.HTML(http.StatusInternalServerError, "user_edit.html", inputUpdate)
		return
	}

	_, err = h.userService.Update(inputUpdate)
	if err != nil {
		inputUpdate.Error = err
		c.HTML(http.StatusInternalServerError, "user_edit.html", inputUpdate)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

func (h *userHandler) Avatar(c *gin.Context) {
	var input user.FormDetailUserInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", nil)
		return
	}

	user, err := h.userService.GetUserByID(input.ID)
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", nil)
		return
	}

	c.HTML(http.StatusOK, "user_avatar.html", user)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	var input user.FormDetailUserInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "user_avatar.html", input)
		return
	}

	path := fmt.Sprintf("images/avatar/%d-%s", input.ID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "user_avatar.html", input)
		return
	}

	_, err = h.userService.SaveAvatar(input.ID, path)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "user_avatar.html", input)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

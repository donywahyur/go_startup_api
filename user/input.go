package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
type FormCreateUserInput struct {
	Name       string `form:"name" binding:"required"`
	Occupation string `form:"occupation" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Password   string `form:"password" binding:"required"`
	Error      error
}

type FormDetailUserInput struct {
	ID int `uri:"id" binding:"required"`
}

type FormUpdateUserInput struct {
	ID         int
	Name       string `form:"name" binding:"required"`
	Occupation string `form:"occupation" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Error      error
}

type FormLoginInput struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

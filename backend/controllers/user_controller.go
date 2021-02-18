package controllers

import (
	"net/http"
	"strconv"

	"github.com/shunyaYoshimra/day27/backend/middleware"

	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/repositories"
)

type UserController struct {
	Repository repositories.UserRepository
}

func NewUserController() UserController {
	return UserController{
		Repository: repositories.NewUserRepository(),
	}
}

func (uc *UserController) Index(c *gin.Context) {
	users := uc.Repository.RetrieveUsers()
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if user, err := uc.Repository.FindByID(id); err != nil {
		res := response.NotFound("This user was not found")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = user
		c.JSON(res.Status, res)
	}
}

func (uc *UserController) GetMe(c *gin.Context) {
	userID := middleware.GetSession(c)
	if user, err := uc.Repository.FindByID(userID); err != nil {
		res := response.NotFound("This user was not found")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = user
		c.JSON(res.Status, res)
	}
}

func (uc *UserController) Signup(c *gin.Context) {
	var user entity.User
	email := c.PostForm("email")
	name := c.PostForm("name")
	pass := c.PostForm("password")

	if name == "" || email == "" || pass == "" {
		res := response.BadRequest("全ての情報を入力してください(All information should be filled in.)")
		c.JSON(res.Status, res)
	} else if _, err := uc.Repository.FindByEmail(email); err == nil {
		res := response.Conflict("このメールアドレスは既に登録されています(This e-mail has already been registered)")
		c.JSON(res.Status, res)
	} else {
		passwordEncrypt, _ := middleware.PasswordEncrypt(pass)
		user = entity.User{
			Name:     name,
			Email:    email,
			Password: passwordEncrypt,
		}
		if err := uc.Repository.Create(&user); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		} else {
			userID := user.ID
			middleware.Login(c, userID)
			res := response.SuccessResponse("")
			c.JSON(res.Status, res)
		}
	}
}

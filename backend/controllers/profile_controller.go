package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/middleware"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/shunyaYoshimra/day27/backend/utils/response"
)

type ProfileController struct {
	Repository repositories.ProfileRepository
}

func NewProfileController() ProfileController {
	return ProfileController{
		Repository: repositories.NewProfileRepository(),
	}
}

func (pc *ProfileController) Create(c *gin.Context) {
	description := c.PostForm("description")
	subject := c.PostForm("subject")
	userID := middleware.GetSession(c)
	if subject == "" {
		res := response.BadRequest("学科は必ず入力してください(Subject should be filled in)")
		c.JSON(res.Status, res)
	} else if _, err := pc.Repository.FindByUser(userID); err == nil {
		res := response.Conflict("プロフィールは既に登録されています(Your profile has already been registerd)")
		c.JSON(res.Status, res)
	} else {
		profile := entity.Profile{
			Description: description,
			Subject:     subject,
			UserID:      userID,
		}
		if err := pc.Repository.Create(&profile); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		} else {
			res := response.SuccessResponse("")
			res.Data = profile
			c.JSON(res.Status, res)
		}
	}
}

func (pc *ProfileController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if profile, err := pc.Repository.FindByUser(id); err != nil {
		res := response.NotFound("プロフィールが見つかりませんでした(This Profile wasn't Found)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = profile
		c.JSON(res.Status, res)
	}
}

func (pc *ProfileController) GetMyProfile(c *gin.Context) {
	userID := middleware.GetSession(c)
	if profile, err := pc.Repository.FindByUser(userID); err != nil {
		res := response.NotFound("プロフィールが見つかりませんでした(This Profile wasn't Found)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = profile
		c.JSON(res.Status, res)
	}
}

func (pc *ProfileController) Update(c *gin.Context) {
	userID := middleware.GetSession(c)
	description := c.PostForm("description")
	subject := c.PostForm("subject")
	if profile, err := pc.Repository.FindByUser(userID); err != nil {
		res := response.NotFound("プロフィールが見つかりませんでした(This Profile wasn't Found)")
		c.JSON(res.Status, res)
	} else if err := pc.Repository.Update(profile, description, subject); err != nil {
		res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = profile
		c.JSON(res.Status, res)
	}
}

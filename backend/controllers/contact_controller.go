package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/middleware"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/shunyaYoshimra/day27/backend/utils/response"
)

type ContactController struct {
	Repository repositories.ContactRepository
}

func NewContactController() ContactController {
	return ContactController{
		Repository: repositories.NewContactRepository(),
	}
}

func (cc *ContactController) Create(c *gin.Context) {
	media := c.PostForm("media")
	account := c.PostForm("account")
	var userID int
	if test := c.PostForm("test"); test == "" {
		userID = middleware.GetSession(c)
	} else {
		testID, _ := strconv.Atoi(test)
		userID = testID
	}
	if _, err := cc.Repository.FindByUser(userID); err == nil {
		res := response.Conflict("コンタクトは既に登録されています(Your contact has already been registerd)")
		c.JSON(res.Status, res)
	} else {
		contact := entity.Contact{
			Media:   media,
			Account: account,
			UserID:  userID,
		}
		if err := cc.Repository.Create(&contact); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		} else {
			res := response.SuccessResponse("")
			res.Data = contact
			c.JSON(res.Status, res)
		}
	}
}

func (cc *ContactController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if contact, err := cc.Repository.FindByUser(id); err != nil {
		res := response.NotFound("コンタクトが見つかりませんでした(This contact wasn't found)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = contact
		c.JSON(res.Status, res)
	}
}

func (cc *ContactController) GetMyContact(c *gin.Context) {
	userID := middleware.GetSession(c)
	if contact, err := cc.Repository.FindByUser(userID); err != nil {
		res := response.NotFound("コンタクトが見つかりませんでした(This contact wasn't found)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = contact
		c.JSON(res.Status, res)
	}
}

func (cc *ContactController) Update(c *gin.Context) {
	media := c.PostForm("media")
	account := c.PostForm("account")
	var userID int
	if test := c.PostForm("test"); test == "" {
		userID = middleware.GetSession(c)
	} else {
		testID, _ := strconv.Atoi(test)
		userID = testID
	}
	if contact, err := cc.Repository.FindByUser(userID); err != nil {
		res := response.NotFound("コンタクトが見つかりませんでした(This contact wasn't Found)")
		c.JSON(res.Status, res)
	} else if err := cc.Repository.Update(contact, media, account); err != nil {
		res := response.SuccessResponse("")
		res.Data = contact
		c.JSON(res.Status, res)
	}
}

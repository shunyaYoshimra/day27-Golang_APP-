package controllers

import (
	"net/http"
	"strconv"

	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/middleware"
	"github.com/shunyaYoshimra/day27/backend/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/repositories"
)

type AbroadController struct {
	Repository repositories.AbroadRepository
}

func NewAbroadController() AbroadController {
	return AbroadController{
		Repository: repositories.NewAbroadRepository(),
	}
}

func (ac *AbroadController) Index(c *gin.Context) {
	abroads := ac.Repository.RetrieveAbroads()
	c.JSON(http.StatusOK, abroads)
}

func (ac *AbroadController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if abroad, err := ac.Repository.FindByUser(id); err != nil {
		res := response.NotFound("海外情報が見つかりませんでした(Not Found)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = abroad
		c.JSON(res.Status, res)
	}
}

func (ac *AbroadController) MyAbroad(c *gin.Context) {
	userID := middleware.GetSession(c)
	if abroad, err := ac.Repository.FindByUser(userID); err != nil {
		res := response.NotFound("海外情報が見つかりませんでした(Not Found)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = abroad
		c.JSON(res.Status, res)
	}
}

func (ac *AbroadController) Create(c *gin.Context) {
	country := c.PostForm("country")
	college := c.PostForm("college")
	description := c.PostForm("description")
	var userID int
	if test := c.PostForm("test"); test == "" {
		userID = middleware.GetSession(c)
	} else {
		testID, _ := strconv.Atoi(test)
		userID = testID
	}
	if country == "" && college == "" && description == "" {
		res := response.BadRequest("情報を入力してください(Please fill in your information)")
		c.JSON(res.Status, res)
	} else {
		abroad := entity.Abroad{
			Country:     country,
			College:     college,
			Description: description,
			UserID:      userID,
		}
		if err := ac.Repository.Create(&abroad); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		} else {
			res := response.SuccessResponse("")
			res.Data = abroad
			c.JSON(res.Status, res)
		}
	}
}

func (ac *AbroadController) Update(c *gin.Context) {
	country := c.PostForm("country")
	college := c.PostForm("college")
	description := c.PostForm("description")
	var userID int
	if test := c.PostForm("test"); test == "" {
		userID = middleware.GetSession(c)
	} else {
		testID, _ := strconv.Atoi(test)
		userID = testID
	}
	if abroad, err := ac.Repository.FindByUser(userID); err != nil {
		res := response.NotFound("海外情報が見つかりませんでした(Not Found)")
		c.JSON(res.Status, res)
	} else if err := ac.Repository.Update(abroad, country, college, description); err != nil {
		res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = abroad
		c.JSON(res.Status, res)
	}
}

package controllers

import (
	"net/http"
	"strconv"

	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/middleware"
	"github.com/shunyaYoshimra/day27/backend/repositories"
)

type CheckController struct {
	Repository repositories.CheckRepository
}

func NewCheckController() CheckController {
	return CheckController{
		Repository: repositories.NewCheckRepository(),
	}
}

func (cc *CheckController) ChecksOfArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	checks := cc.Repository.ChecksOfArticle(id)
	c.JSON(http.StatusOK, checks)
}

func (cc *CheckController) ChecksOfUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	checks := cc.Repository.ChecksOfUser(id)
	c.JSON(http.StatusOK, checks)
}

func (cc *CheckController) Create(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))
	userID := middleware.GetSession(c)
	if _, err := cc.Repository.CheckUnique(articleID, userID); err == nil {
		res := response.NotFound("This check was not found")
		c.JSON(res.Status, res)
	} else {
		check := entity.Check{
			ArticleID: articleID,
			UserID:    userID,
		}
		if err := cc.Repository.Create(&check); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		} else {
			res := response.SuccessResponse("")
			res.Data = check
			c.JSON(res.Status, res)
		}
	}
}

func (cc *CheckController) Delete(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))
	userID := middleware.GetSession(c)
	if check, err := cc.Repository.CheckUnique(articleID, userID); err != nil {
		res := response.Conflict("was not unique")
		c.JSON(res.Status, res)
	} else if err := cc.Repository.Delete(check); err != nil {
		res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		c.JSON(res.Status, res)
	}
}

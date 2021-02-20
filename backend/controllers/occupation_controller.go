package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/middleware"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/shunyaYoshimra/day27/backend/utils/response"
)

type OccupationController struct {
	Repository repositories.OccupationRepository
}

func NewOccupationController() OccupationController {
	return OccupationController{
		Repository: repositories.NewOccupationRepository(),
	}
}

func (oc *OccupationController) Index(c *gin.Context) {
	occupations := oc.Repository.RetrieveOccuaptions()
	c.JSON(http.StatusOK, occupations)
}

func (oc *OccupationController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if occupation, err := oc.Repository.FindByUser(id); err != nil {
		res := response.NotFound("職業情報が見つかりませんでした(Not Found)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = occupation
		c.JSON(res.Status, res)
	}
}

func (oc *OccupationController) MyOccupation(c *gin.Context) {
	userID := middleware.GetSession(c)
	if occupation, err := oc.Repository.FindByUser(userID); err != nil {
		res := response.NotFound("職業情報が見つかりませんでした(Not Found)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = occupation
		c.JSON(res.Status, res)
	}
}

func (oc *OccupationController) Create(c *gin.Context) {
	kind := c.PostForm("kind")
	company := c.PostForm("company")
	description := c.PostForm("description")
	var userID int
	if test := c.PostForm("test"); test == "" {
		userID = middleware.GetSession(c)
	} else {
		testID, _ := strconv.Atoi(test)
		userID = testID
	}
	if kind == "" && company == "" && description == "" {
		res := response.BadRequest("情報を入力してください(Please fill in your information)")
		c.JSON(res.Status, res)
	} else {
		occupation := entity.Occupation{
			Kind:        kind,
			Company:     company,
			Description: description,
			UserID:      userID,
		}
		if err := oc.Repository.Create(&occupation); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		} else {
			res := response.SuccessResponse("")
			res.Data = occupation
			c.JSON(res.Status, res)
		}
	}
}

func (oc *OccupationController) Update(c *gin.Context) {
	kind := c.PostForm("kind")
	company := c.PostForm("company")
	description := c.PostForm("description")
	var userID int
	if test := c.PostForm("test"); test == "" {
		userID = middleware.GetSession(c)
	} else {
		testID, _ := strconv.Atoi(test)
		userID = testID
	}
	if occupation, err := oc.Repository.FindByUser(userID); err != nil {
		res := response.NotFound("職業情報が見つかりませんでした(Not Found)")
		c.JSON(res.Status, res)
	} else if err := oc.Repository.Update(occupation, kind, company, description); err != nil {
		res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = occupation
		c.JSON(res.Status, res)
	}
}

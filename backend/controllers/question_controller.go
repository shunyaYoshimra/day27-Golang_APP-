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

type QuestionController struct {
	Repository repositories.QuestionRepository
}

func NewQuestionController() QuestionController {
	return QuestionController{
		Repository: repositories.NewQuestionRepository(),
	}
}

func (qc *QuestionController) Index(c *gin.Context) {
	questions := qc.Repository.RetrieveQuestions()
	c.JSON(http.StatusOK, questions)
}

func (qc *QuestionController) SearchedIndex(c *gin.Context) {
	about := c.Param("about")
	questions := qc.Repository.FindByAbout(about)
	c.JSON(http.StatusOK, questions)
}

func (qc *QuestionController) UserQuestions(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	questions := qc.Repository.FindByUser(id)
	c.JSON(http.StatusOK, questions)
}

func (qc *QuestionController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if question, err := qc.Repository.FindByID(id); err != nil {
		res := response.NotFound("質問が見つかりませんでした(The question was not found)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = question
		c.JSON(res.Status, res)
	}
}

func (qc *QuestionController) Create(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	about := c.PostForm("about")
	var userID int
	if test := c.PostForm("test"); test == "" {
		userID = middleware.GetSession(c)
	} else {
		testID, _ := strconv.Atoi(test)
		userID = testID
	}
	if title == "" || content == "" || about == "" {
		res := response.BadRequest("全ての情報を入力して下さい(All information should be filled in)")
		c.JSON(res.Status, res)
	} else {
		question := entity.Question{
			Title:   title,
			Content: content,
			About:   about,
			UserID:  userID,
		}
		if err := qc.Repository.Create(&question); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		} else {
			res := response.SuccessResponse("")
			res.Data = question
			c.JSON(res.Status, res)
		}
	}
}

func (qc *QuestionController) Update(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	id, _ := strconv.Atoi(c.Param("id"))
	if question, err := qc.Repository.FindByID(id); err != nil {
		res := response.NotFound("質問が見つかりませんでした(This question was not found)")
		c.JSON(res.Status, res)
	} else if err := qc.Repository.Update(question, title, content); err != nil {
		res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = question
		c.JSON(res.Status, res)
	}
}

func (qc *QuestionController) ChangeStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if question, err := qc.Repository.FindByID(id); err != nil {
		res := response.NotFound("質問が見つかりませんでした(This question was not found)")
		c.JSON(res.Status, res)
	} else if err := qc.Repository.ChangeStatus(question); err != nil {
		res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		res.Data = question
		c.JSON(res.Status, res)
	}
}

func (qc *QuestionController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if question, err := qc.Repository.FindByID(id); err != nil {
		res := response.NotFound("質問が見つかりませんでした(This question was not found)")
		c.JSON(res.Status, res)
	} else if err := qc.Repository.Delete(question); err != nil {
		res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		c.JSON(res.Status, res)
	}
}

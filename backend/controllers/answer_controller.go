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

type AnswerController struct {
	Repository repositories.AnswerRepository
}

func NewAnswerController() AnswerController {
	return AnswerController{
		Repository: repositories.NewAnswerRepository(),
	}
}

func (ac *AnswerController) Index(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	answers := ac.Repository.RetrieveAnswers(id)
	c.JSON(http.StatusOK, answers)
}

func (ac *AnswerController) Create(c *gin.Context) {
	content := c.PostForm("content")
	questionID, _ := strconv.Atoi(c.Param("id"))
	var userID int
	if test := c.PostForm("test"); test == "" {
		userID = middleware.GetSession(c)
	} else {
		testID, _ := strconv.Atoi(test)
		userID = testID
	}
	if content == "" {
		res := response.BadRequest("内容を入力してください(Content should be filled in)")
		c.JSON(res.Status, res)
	} else {
		answer := entity.Answer{
			Content:    content,
			QuestionID: questionID,
			UserID:     userID,
		}
		if err := ac.Repository.Create(&answer); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		} else {
			res := response.SuccessResponse("")
			res.Data = answer
			c.JSON(res.Status, res)
		}
	}
}

func (ac *AnswerController) Update(c *gin.Context) {
	content := c.PostForm("content")
	answerID, _ := strconv.Atoi(c.PostForm("answerID"))
	questionID, _ := strconv.Atoi(c.PostForm("questionID"))
	if content == "" {
		res := response.BadRequest("内容は必ず入力して下さい(Content should be filled in)")
		c.JSON(res.Status, res)
	} else if answer, err := ac.Repository.FindByID(answerID); err != nil {
		res := response.NotFound("回答が見つかりませんでした(Answer was not found)")
		c.JSON(res.Status, res)
	} else if err := ac.Repository.Update(answer, content); err != nil {
		res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
		c.JSON(res.Status, res)
	} else {
		answers := ac.Repository.RetrieveAnswers(questionID)
		c.JSON(http.StatusOK, answers)
	}
}

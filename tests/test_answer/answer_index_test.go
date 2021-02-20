package test_answer

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestAnswerIndex(id string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	answerRepository := repositories.AnswerRepository{Conn: database.GetDB().Table("answers")}
	answerRepository.Create(&entity.Answer{
		ID:         1,
		Content:    "test content",
		UserID:     1,
		QuestionID: 1,
	})
	answerRepository.Create(&entity.Answer{
		ID:         2,
		Content:    "test content",
		UserID:     1,
		QuestionID: 1,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/test/answers/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestAnswerIndex(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestAnswerIndex("1")
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

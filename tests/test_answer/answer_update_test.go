package test_answer

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestAnswerUpdate(values []string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	answerRepository := repositories.AnswerRepository{Conn: database.GetDB().Table("answers")}
	answerRepository.Create(&entity.Answer{
		ID:         1,
		Content:    "test content",
		QuestionID: 1,
	})

	form := url.Values{}
	form.Add("content", values[0])
	form.Add("answerID", values[1])
	form.Add("questionID", values[2])
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/test/answers", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestAnswerUpdate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update content", "1", "1"}
		w := InitTestAnswerUpdate(values)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return Not Found with invalid id", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update content", "2", "1"}
		w := InitTestAnswerUpdate(values)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("it should return Bad Request with empty content", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"", "1", "1"}
		w := InitTestAnswerUpdate(values)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

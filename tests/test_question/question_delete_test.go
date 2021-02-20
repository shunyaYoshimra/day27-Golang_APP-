package test_question

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shunyaYoshimra/day27/backend/apps"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"

	"github.com/gin-gonic/gin"
)

func InitTestQuestionDelete(id string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	questionRepository := repositories.QuestionRepository{Conn: database.GetDB().Table("questions")}
	questionRepository.Create(&entity.Question{
		ID:      1,
		Title:   "test title",
		Content: "test content",
		About:   "test about",
		UserID:  1,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/test/questions/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestQuestionDelete(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestQuestionDelete("1")
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return Not Found with invalid id", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestQuestionDelete("2")
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

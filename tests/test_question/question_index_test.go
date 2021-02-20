package test_question

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/stretchr/testify/assert"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/repositories"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestQuestionIndex() *httptest.ResponseRecorder {
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
	questionRepository.Create(&entity.Question{
		ID:      2,
		Title:   "test title",
		Content: "test content",
		About:   "test about",
		UserID:  2,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/test/questions", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestQuestionIndex(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestQuestionIndex()
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

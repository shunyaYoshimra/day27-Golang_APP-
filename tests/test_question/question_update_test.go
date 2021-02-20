package test_question

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/utils/response"
	"github.com/stretchr/testify/assert"

	"github.com/shunyaYoshimra/day27/backend/database/entity"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/repositories"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestQuestionUpdate(values []string) *httptest.ResponseRecorder {
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

	form := url.Values{}
	form.Add("title", values[0])
	form.Add("content", values[1])
	form.Add("about", values[2])
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/test/questions/"+values[3], body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestQuestionUpdate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update title", "update content", "update about", "1"}
		w := InitTestQuestionUpdate(values)
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, actual.Data)
	})
	t.Run("it should return Not Found with invalid id", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update title", "update content", "update about", "2"}
		w := InitTestQuestionUpdate(values)
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, actual.Data)
	})
}

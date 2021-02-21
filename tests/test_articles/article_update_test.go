package test_articles

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/shunyaYoshimra/day27/backend/utils/response"
	"github.com/stretchr/testify/assert"
)

func InitTestArticleUpdate(values []string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	articleRepository := repositories.ArticleRepository{Conn: database.GetDB().Table("articles")}
	articleRepository.Create(&entity.Article{
		ID:     1,
		Title:  "test title",
		UserID: 1,
	})

	form := url.Values{}
	form.Add("title", values[0])
	form.Add("lines", values[1])
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/test/articles/"+values[2], body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestArticleUpdate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update title", "update linse", "1"}
		w := InitTestArticleUpdate(values)
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, actual.Data)
	})
	t.Run("it should return Not Found with invalid id", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update title", "update linse", "2"}
		w := InitTestArticleUpdate(values)
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, actual.Data)
	})
}

package test_articles

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

func InitTestArticleIndex() *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	articleRepository := repositories.ArticleRepository{Conn: database.GetDB().Table("articles")}
	articleRepository.Create(&entity.Article{
		ID:     1,
		Title:  "test title",
		UserID: 1,
	})
	articleRepository.Create(&entity.Article{
		ID:     2,
		Title:  "test title",
		UserID: 2,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/test/articles", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestArticleIndex(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestArticleIndex()
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

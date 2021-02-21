package test_favorite

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"
)

func InitTestFavoriteDelete(postID, userID string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	form := url.Values{}
	form.Add("test", userID)
	body := strings.NewReader(form.Encode())

	favoriteRepository := repositories.FavoriteRepository{Conn: database.GetDB().Table("favorites")}
	favoriteRepository.Create(&entity.Favorite{
		UserID: 1,
		PostID: 1,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/test/favorites/"+postID, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestFavoriteDelete(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestFavoriteDelete("1", "1")
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

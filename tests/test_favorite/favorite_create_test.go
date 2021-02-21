package test_favorite

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/shunyaYoshimra/day27/backend/utils/response"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestFavoriteCreate(postID, userID string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	form := url.Values{}
	form.Add("test", userID)
	body := strings.NewReader(form.Encode())

	favoriteRepository := repositories.FavoriteRepository{Conn: database.GetDB().Table("favorites")}
	favoriteRepository.Create(&entity.Favorite{
		UserID: 2,
		PostID: 2,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/test/favorites/"+postID, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestFavoriteCreate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestFavoriteCreate("1", "1")
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, actual.Data)
	})
	t.Run("it should return Conflict with duplicate id", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestFavoriteCreate("2", "2")
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Empty(t, actual.Data)
	})
}

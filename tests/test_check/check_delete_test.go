package test_check

import (
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
	"github.com/stretchr/testify/assert"
)

func InitTestCheckDelete(articleID, userID string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	form := url.Values{}
	form.Add("test", userID)
	body := strings.NewReader(form.Encode())

	checkRepository := repositories.CheckRepository{Conn: database.GetDB().Table("checks")}
	checkRepository.Create(&entity.Check{
		UserID:    1,
		ArticleID: 1,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/test/checks/"+articleID, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestFavoriteDelete(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestCheckDelete("1", "1")
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

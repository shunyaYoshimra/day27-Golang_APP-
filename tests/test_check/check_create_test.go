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

func InitTestCheckCreate(articleID, userID string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	form := url.Values{}
	form.Add("test", userID)
	body := strings.NewReader(form.Encode())

	checkRepository := repositories.CheckRepository{Conn: database.GetDB().Table("checks")}
	checkRepository.Create(&entity.Check{
		UserID:    2,
		ArticleID: 2,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/test/checks/"+articleID, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestCheckCreate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestCheckCreate("1", "1")
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return Conflict with duplicate id", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestCheckCreate("2", "2")
		assert.Equal(t, http.StatusConflict, w.Code)
	})
}

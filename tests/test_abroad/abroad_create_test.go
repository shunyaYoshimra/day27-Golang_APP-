package test_abroad

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestAbroadCreate(values []string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	form := url.Values{}
	form.Add("country", values[0])
	form.Add("college", values[1])
	form.Add("description", values[2])
	form.Add("test", values[3])
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/test/abroads", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestAbroadCreate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"teset country", "test college", "test description", "10"}
		w := InitTestAbroadCreate(values)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return Bad Request with invalid params", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"", "", "", "12"}
		w := InitTestAbroadCreate(values)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

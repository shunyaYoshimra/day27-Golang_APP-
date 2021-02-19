package test_profile

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shunyaYoshimra/day27/backend/database"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestProfileCreate(values []string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	form := url.Values{}
	form.Add("description", values[0])
	form.Add("subject", values[1])
	form.Add("test", values[2])
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/test/profiles", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestProfileCreate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"test description", "test subject", "5"}
		w := InitTestProfileCreate(values)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return bad request with empty subject", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"test description", "", "6"}
		w := InitTestProfileCreate(values)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("it should return conflict with duplicate user id", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"test description", "test subject", "1"}
		w := InitTestProfileCreate(values)
		assert.Equal(t, http.StatusConflict, w.Code)
	})
}

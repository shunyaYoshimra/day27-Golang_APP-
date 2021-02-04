package user_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestUserCreate(body map[string]interface{}) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	w := httptest.NewRecorder()
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/api/users", strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestUserCreate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		body := map[string]interface{}{
			"name":     "User Test",
			"email":    "user@email.com",
			"password": "password",
		}
		w := InitTestUserCreate(body)
		t.Log(w)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

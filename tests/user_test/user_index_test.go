package user_test

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

func InitTestUserIndex() *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	userRepository := repositories.UserRepository{Conn: database.GetDB().Table("users")}
	userRepository.Create(&entity.User{
		ID:       1,
		Name:     "shunya",
		Email:    "maoorgri1015@gmail.com",
		Password: "password",
	})
	userRepository.Create(&entity.User{
		ID:       2,
		Name:     "saya",
		Email:    "saya@gmail.com",
		Password: "password",
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/users", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestUserIndex(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestUserIndex()
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

package user_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestUserShow(id string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	userRepository := repositories.UserRepository{Conn: database.GetDB().Table("users")}
	userRepository.Create(entity.User{
		ID:       1,
		Name:     "shunya",
		Email:    "maoorgri1015@gmail.com",
		Password: "password",
	})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/user/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestUserShow(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestUserShow("1")
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("is should return 404 with invalid params url", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestUserShow("2")
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

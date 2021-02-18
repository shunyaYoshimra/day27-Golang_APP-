package user_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestUserCreate(values []string, duplicate bool) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	if duplicate == true {
		userRepository := repositories.UserRepository{Conn: database.GetDB().Table("users")}
		userRepository.Create(&entity.User{
			ID:       1,
			Name:     "test user1",
			Email:    "duplicate@gmail.com",
			Password: "password",
		})
	}

	form := url.Values{}
	form.Add("name", values[0])
	form.Add("email", values[1])
	form.Add("password", values[2])
	form.Add("test", "test")
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/test/signup", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestUserCreate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"User Test", "test@gmail.com", "password"}
		w := InitTestUserCreate(values, false)
		t.Log(w)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return 400 with invalid data", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"", "", ""}
		w := InitTestUserCreate(values, false)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("it should return 400 with duplicate email", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"User Test", "duplicate@gmail.com", "password"}
		w := InitTestUserCreate(values, true)
		assert.Equal(t, http.StatusConflict, w.Code)
	})
}

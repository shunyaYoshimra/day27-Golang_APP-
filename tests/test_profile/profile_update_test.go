package test_profile

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/stretchr/testify/assert"

	"github.com/shunyaYoshimra/day27/backend/database"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestProfileUpdate(values []string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	profileRepository := repositories.ProfileRepository{Conn: database.GetDB().Table("profiles")}
	profileRepository.Create(&entity.Profile{
		ID:          11,
		Description: "test description",
		Subject:     "test subject",
	})

	form := url.Values{}
	form.Add("description", values[0])
	form.Add("subject", values[1])
	form.Add("test", values[2])
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/test/profiles", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestprofileUpdate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update description", "update subject", "11"}
		w := InitTestProfileUpdate(values)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return Not Found with duplicate user id", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update description", "update subject", "12"}
		w := InitTestProfileUpdate(values)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

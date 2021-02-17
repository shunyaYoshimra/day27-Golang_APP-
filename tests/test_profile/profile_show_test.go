package test_profile

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

func InitTestProfileShow(id string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	profileRepository := repositories.ProfileRepository{Conn: database.GetDB().Table("profiles")}
	profileRepository.Create(&entity.Profile{
		ID:          1,
		Description: "test description",
		Subject:     "test subject",
		UserID:      1,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/test/profile/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestProfileShow(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestProfileShow("1")
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return 404 with invalid params url", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestProfileShow("2")
		t.Log(w)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

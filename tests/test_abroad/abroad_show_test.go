package test_abroad

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

func InitTestAbroadShow(id string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	abraodRepository := repositories.AbroadRepository{Conn: database.GetDB().Table("abroads")}
	abraodRepository.Create(&entity.Abroad{
		ID:      1,
		Country: "test country",
		College: "test college",
		UserID:  1,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/test/abroad/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestAbroadShow(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestAbroadShow("1")
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return 404 with invalid params url", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestAbroadShow("3")
		t.Log(w)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

package test_abroad

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"
)

func InitTestAbroadIndex() *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	abroadRepository := repositories.AbroadRepository{Conn: database.GetDB().Table("abroads")}
	abroadRepository.Create(&entity.Abroad{
		ID:          1,
		Country:     "test country",
		College:     "test college",
		Description: "test description",
		UserID:      1,
	})
	abroadRepository.Create(&entity.Abroad{
		ID:          2,
		Country:     "test country",
		College:     "test college",
		Description: "test description",
		UserID:      2,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/test/abroads", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestAbroadIndex(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestAbroadIndex()
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

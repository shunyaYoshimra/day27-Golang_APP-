package test_contact

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

func InitTestContactShow(id string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	contactRepository := repositories.ContactRepository{Conn: database.GetDB().Table("contacts")}
	contactRepository.Create(&entity.Contact{
		ID:      1,
		Media:   "test media",
		Account: "test account",
		UserID:  1,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/test/contact/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestContactShow(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestContactShow("1")
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return 404 with invalid params url", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestContactShow("2")
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

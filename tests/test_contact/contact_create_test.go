package test_contact

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestContactCreate(values []string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	form := url.Values{}
	form.Add("medial", values[0])
	form.Add("account", values[1])
	form.Add("test", values[2])
	body := strings.NewReader(form.Encode())

	contactRepository := repositories.ContactRepository{Conn: database.GetDB().Table("contacts")}
	contactRepository.Create(&entity.Contact{
		UserID: 2,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/test/contacts", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestContactCreate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"test media", "test account", "1"}
		w := InitTestContactCreate(values)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return bad request with invalid params", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"test media", "test account", "2"}
		w := InitTestContactCreate(values)
		assert.Equal(t, http.StatusConflict, w.Code)
	})
}

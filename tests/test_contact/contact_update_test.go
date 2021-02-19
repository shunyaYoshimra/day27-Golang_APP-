package test_contact

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/stretchr/testify/assert"
)

func InitTestContactUpdate(values []string) *httptest.ResponseRecorder {
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
		UserID: 1,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/test/contacts", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestContactUpdate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update media", "update account", "1"}
		w := InitTestContactUpdate(values)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it should return Not Found with invalid user id", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update media", "update account", "2"}
		w := InitTestContactUpdate(values)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

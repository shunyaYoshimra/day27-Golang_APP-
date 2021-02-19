package test_abroad

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/repositories"
	"github.com/shunyaYoshimra/day27/backend/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestAbroadUpdate(values []string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	abroadRepository := repositories.AbroadRepository{Conn: database.GetDB().Table("abroads")}
	abroadRepository.Create(&entity.Abroad{
		ID:      21,
		Country: "test country",
		College: "test college",
		UserID:  21,
	})

	form := url.Values{}
	form.Add("country", values[0])
	form.Add("college", values[1])
	form.Add("description", values[2])
	form.Add("test", values[3])
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/test/abroads", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestAbroadUpdate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update country", "update college", "update description", "21"}
		w := InitTestAbroadUpdate(values)
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, actual.Data)
	})
	t.Run("it should return Not Found", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update country", "update college", "update description", "22"}
		w := InitTestAbroadUpdate(values)
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, actual.Data)
	})
}

package test_occupation

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/shunyaYoshimra/day27/backend/apps"
	"github.com/shunyaYoshimra/day27/backend/utils/response"

	"github.com/gin-gonic/gin"
)

func InitTestOccupationCreate(values []string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	form := url.Values{}
	form.Add("kind", values[0])
	form.Add("company", values[1])
	form.Add("description", values[2])
	form.Add("test", values[3])
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/test/occupations", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestOccupationCreate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		values := []string{"test kind", "test college", "test description", "1"}
		w := InitTestOccupationCreate(values)
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, actual.Data)
	})
	t.Run("it should return Bad Request with empty params", func(t *testing.T) {
		values := []string{"", "", "", "1"}
		w := InitTestOccupationCreate(values)
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Empty(t, actual.Data)
	})
}

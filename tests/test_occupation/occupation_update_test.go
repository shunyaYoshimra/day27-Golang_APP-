package test_occupation

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/utils/response"
	"github.com/stretchr/testify/assert"

	"github.com/shunyaYoshimra/day27/backend/repositories"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestOccupationUpdate(values []string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	occupationRepository := repositories.OccupationRepository{Conn: database.GetDB().Table("occupations")}
	occupationRepository.Create(&entity.Occupation{
		ID:          1,
		Kind:        "test kind",
		Company:     "test company",
		Description: "test description",
		UserID:      1,
	})

	form := url.Values{}
	form.Add("kind", values[0])
	form.Add("company", values[1])
	form.Add("description", values[2])
	form.Add("test", values[3])
	body := strings.NewReader(form.Encode())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/test/occupations", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(w, req)
	return w
}

func TestOccupationUpdate(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update kind", "update company", "update description", "1"}
		w := InitTestOccupationUpdate(values)
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, actual.Data)
	})
	t.Run("it should return Not Found with invalid params", func(t *testing.T) {
		defer database.DropAllTable()
		values := []string{"update kind", "update company", "update description", "2"}
		w := InitTestOccupationUpdate(values)
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, actual.Data)
	})
}

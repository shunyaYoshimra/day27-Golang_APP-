package test_occupation

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/stretchr/testify/assert"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/repositories"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestOccupationIndex() *httptest.ResponseRecorder {
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
	occupationRepository.Create(&entity.Occupation{
		ID:          2,
		Kind:        "test kind",
		Company:     "test company",
		Description: "test description",
		UserID:      2,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/test/occupations", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestOccupationIndex(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestOccupationIndex()
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

package test_occupation

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

func InitTestOccupationShow(id string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	aboadRepository := repositories.OccupationRepository{Conn: database.GetDB().Table("occupations")}
	aboadRepository.Create(&entity.Occupation{
		ID:          3,
		Kind:        "test kind",
		Company:     "test company",
		Description: "test description",
		UserID:      3,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/test/occupation/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestOccupationShow(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestOccupationShow("3")
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("it shoule return 404 with invalid prams url", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestOccupationShow("4")
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

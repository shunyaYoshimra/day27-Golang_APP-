package test_post

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shunyaYoshimra/day27/backend/database/entity"
	"github.com/shunyaYoshimra/day27/backend/utils/response"
	"github.com/stretchr/testify/assert"

	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/repositories"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/apps"
)

func InitTestPostShow(id string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	postRepository := repositories.PostRepository{Conn: database.GetDB().Table("posts")}
	postRepository.Create(&entity.Post{
		ID:      1,
		Content: "test content",
		Tags:    "test tags",
		UserID:  1,
	})
	imageRepository := repositories.ImageRepository{Conn: database.GetDB().Table("images")}
	imageRepository.Create(&entity.Image{
		ID:       1,
		FileName: "test file name",
		Time:     "test time",
		PostID:   1,
	})
	imageRepository.Create(&entity.Image{
		ID:       2,
		FileName: "test file name",
		Time:     "test time",
		PostID:   1,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/test/post/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestPostShow(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestPostShow("1")
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, actual.Data)
	})
	t.Run("it should return Not Found with invalid id", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestPostShow("2")
		actual := response.TestResponse{}
		if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
			panic(err)
		}
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Empty(t, actual.Data)
	})
}

package test_post

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

func InitTestUserPosts(id string) *httptest.ResponseRecorder {
	r := gin.Default()
	app := new(apps.Application)
	app.CreateTest(r)

	postRepository := repositories.PostRepository{Conn: database.GetDB().Table("posts")}
	postRepository.Create(&entity.Post{
		ID:      1,
		Content: "test content",
		Tags:    "test, tags",
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
	req, _ := http.NewRequest(http.MethodGet, "/api/test/posts/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	return w
}

func TestUserPosts(t *testing.T) {
	t.Run("it should return success", func(t *testing.T) {
		defer database.DropAllTable()
		w := InitTestUserPosts("1")
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

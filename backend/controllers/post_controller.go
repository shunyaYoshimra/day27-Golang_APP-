package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/shunyaYoshimra/day27/backend/middleware"
	"github.com/shunyaYoshimra/day27/backend/utils/response"

	"github.com/shunyaYoshimra/day27/backend/database/entity"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/repositories"
)

type PostController struct {
	PostRepository  repositories.PostRepository
	ImageRepository repositories.ImageRepository
}

func NewPostController() PostController {
	return PostController{
		PostRepository:  repositories.NewPostRepository(),
		ImageRepository: repositories.NewImageRepository(),
	}
}

func (pc *PostController) Index(c *gin.Context) {
	posts := pc.PostRepository.RetrievePosts()
	images := pc.ImageRepository.RetrieveImages()
	c.JSON(http.StatusOK, gin.H{
		"posts":  posts,
		"images": images,
	})
}

func (pc *PostController) UserPosts(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	posts := pc.PostRepository.PostsByUser(id)
	var images []entity.Image
	for _, post := range posts {
		value := pc.ImageRepository.ImagesByPost(post.ID)
		images = append(images, value...)
	}
	c.JSON(http.StatusOK, gin.H{
		"posts":  posts,
		"images": images,
	})
}

func (pc *PostController) SearchedIndex(c *gin.Context) {
	keyword := c.Param("key-word")
	posts := pc.PostRepository.PostsByKeyword(keyword)
	var images []entity.Image
	for _, post := range posts {
		value := pc.ImageRepository.ImagesByPost(post.ID)
		images = append(images, value...)
	}
	c.JSON(http.StatusOK, gin.H{
		"posts":  posts,
		"images": images,
	})
}

func (pc *PostController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if post, err := pc.PostRepository.FindByID(id); err != nil {
		res := response.NotFound("投稿が見つかりませんでした(This post was not found)")
		c.JSON(res.Status, res)
	} else {
		images := pc.ImageRepository.ImagesByPost(post.ID)
		res := response.SuccessResponse("")
		c.JSON(res.Status, gin.H{
			"post":   post,
			"images": images,
		})
	}
}

func (pc *PostController) Create(c *gin.Context) {
	form, _ := c.MultipartForm()
	content := c.PostForm("content")
	tags := c.PostForm("tags")
	time := c.PostForm("time")
	userID := middleware.GetSession(c)
	if content == "" {
		res := response.BadRequest("内容は必ず記入してください(Content should be filled in)")
		c.JSON(res.Status, res)
	} else {
		post := entity.Post{
			Content: content,
			Tags:    tags,
			UserID:  userID,
		}
		if err := pc.PostRepository.Create(&post); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		} else {
			files := form.File["file"]
			if files != nil {
				for n, file := range files {
					id := strconv.Itoa(n)
					if err := c.SaveUploadedFile(file, "frontend/dist/img/"+time+"_"+id+".png"); err != nil {
						res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
						c.JSON(res.Status, res)
					} else {
						image := entity.Image{
							FileName: time + "_" + id + ".png",
							Time:     time,
							PostID:   post.ID,
						}
						if err := pc.ImageRepository.Create(&image); err != nil {
							res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
							c.JSON(res.Status, res)
						} else {
							res := response.SuccessResponse("")
							c.JSON(res.Status, res)
						}
					}
				}
			}
		}
	}
}

func (pc *PostController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := middleware.GetSession(c)
	if post, err := pc.PostRepository.FindByID(id); err != nil {
		res := response.NotFound("この投稿は見つかりませんでした(This post was not found)")
		c.JSON(res.Status, res)
	} else if userID != post.UserID {
		res := response.Forbidden("権限がありません(Out of authorization)")
		c.JSON(res.Status, res)
	} else {
		images := pc.ImageRepository.ImagesByPost(post.ID)
		for _, image := range images {
			os.Remove(fmt.Sprintf("frontend/dist/img/%s", image.FileName))
		}
	}
	res := response.SuccessResponse("")
	c.JSON(res.Status, res)
}

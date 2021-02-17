package controllers

import (
	"net/http"
	"strconv"

	"github.com/shunyaYoshimra/day27/backend/database/entity"

	"github.com/shunyaYoshimra/day27/backend/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/middleware"
	"github.com/shunyaYoshimra/day27/backend/repositories"
)

type FavoriteController struct {
	Repository repositories.FavoriteRepository
}

func NewFavoriteController() FavoriteController {
	return FavoriteController{
		Repository: repositories.NewFavoriteRepository(),
	}
}

func (fc *FavoriteController) MyFavorites(c *gin.Context) {
	userID := middleware.GetSession(c)
	favorites := fc.Repository.FavroitesOfUser(userID)
	c.JSON(http.StatusOK, favorites)
}

func (fc *FavoriteController) FavoritesOfUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	favorites := fc.Repository.FavroitesOfUser(userID)
	c.JSON(http.StatusOK, favorites)
}

func (fc *FavoriteController) Create(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))
	userID := middleware.GetSession(c)
	if _, err := fc.Repository.CheckUnique(postID, userID); err == nil {
		res := response.Conflict("Conflict!!!")
		c.JSON(res.Status, res)
	} else {
		favorite := entity.Favorite{
			PostID: postID,
			UserID: userID,
		}
		if err := fc.Repository.Create(&favorite); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		} else {
			res := response.SuccessResponse("")
			res.Data = favorite
			c.JSON(res.Status, res)
		}
	}
}

func (fc *FavoriteController) Delete(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))
	userID := middleware.GetSession(c)
	if favorite, err := fc.Repository.CheckUnique(postID, userID); err != nil {
		res := response.Conflict("was not unique")
		c.JSON(res.Status, res)
	} else if err := fc.Repository.Delete(favorite); err != nil {
		res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
		c.JSON(res.Status, res)
	} else {
		res := response.SuccessResponse("")
		c.JSON(res.Status, res)
	}
}

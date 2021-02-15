package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/shunyaYoshimra/day27/backend/database/entity"

	"github.com/shunyaYoshimra/day27/backend/middleware"

	"github.com/shunyaYoshimra/day27/backend/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/shunyaYoshimra/day27/backend/repositories"
)

type ArticleController struct {
	ArticleRepository     repositories.ArticleRepository
	ArticleLineRepository repositories.ArticleLineRepository
	BoldNumberRepository  repositories.BoldNumberRepository
	LinkNumberRepository  repositories.LinkNumberRepository
}

func NewArticleController() ArticleController {
	return ArticleController{
		ArticleRepository:     repositories.NewArticleRepository(),
		ArticleLineRepository: repositories.NewArticleLineRepository(),
		BoldNumberRepository:  repositories.NewBoldNumberRepository(),
		LinkNumberRepository:  repositories.NewLinkNumberRepository(),
	}
}

func (ac *ArticleController) Index(c *gin.Context) {
	articles := ac.ArticleRepository.RetrieveArticles()
	c.JSON(http.StatusOK, articles)
}

func (ac *ArticleController) UserArticles(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	articles := ac.ArticleRepository.ArticlesOfUser(id)
	c.JSON(http.StatusOK, articles)
}

func (ac *ArticleController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if article, err := ac.ArticleRepository.FindByID(id); err != nil {
		res := response.NotFound("投稿が見つかりませんでした(This article was not found)")
		c.JSON(res.Status, res)
	} else {
		articleLines := ac.ArticleLineRepository.ArticleLinesOfArticle(id)
		boldNumbers := ac.BoldNumberRepository.BoldNumbersOfArticle(id)
		linkNumbers := ac.LinkNumberRepository.LinkNumbersOfArticle(id)
		c.JSON(http.StatusOK, gin.H{
			"article":      article,
			"articleLines": articleLines,
			"boldNumbers":  boldNumbers,
			"linkNumbers":  linkNumbers,
		})
	}
}

func (ac *ArticleController) Create(c *gin.Context) {
	title := c.PostForm("title")
	lines := strings.Split(c.PostForm("lines"), "&divide;,")
	boldNumbers := strings.Split(c.PostForm("bold-numbers"), ",")
	linkNumbers := strings.Split(c.PostForm("link-numbers"), ",")
	userID := middleware.GetSession(c)
	if title == "" || lines == nil {
		res := response.BadRequest("タイトルと内容は必ず記入してください(Title and Content should be filled in)")
		c.JSON(res.Status, res)
	} else {
		article := entity.Article{
			Title:  title,
			UserID: userID,
		}
		if err := ac.ArticleRepository.Create(&article); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		} else {
			var articleLine entity.ArticleLine
			var boldNumber entity.BoldNumber
			var linkNumber entity.LinkNumber
			var number int
			for i := 0; i < len(lines); i++ {
				articleLine = entity.ArticleLine{
					Line:      lines[i],
					ArticleID: article.ID,
				}
				if err := ac.ArticleLineRepository.Create(&articleLine); err != nil {
					res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
					c.JSON(res.Status, res)
				}
			}
			for i := 0; i < len(boldNumbers); i++ {
				number, _ = strconv.Atoi(boldNumbers[i])
				boldNumber = entity.BoldNumber{
					Number:    number,
					ArticleID: article.ID,
				}
				if err := ac.BoldNumberRepository.Create(&boldNumber); err != nil {
					res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
					c.JSON(res.Status, res)
				}
			}
			for i := 0; i < len(linkNumbers); i++ {
				number, _ = strconv.Atoi(linkNumbers[i])
				linkNumber = entity.LinkNumber{
					Number:    number,
					ArticleID: article.ID,
				}
				if err := ac.LinkNumberRepository.Create(&linkNumber); err != nil {
					res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
					c.JSON(res.Status, res)
				}
			}
		}
		res := response.SuccessResponse("")
		res.Data = article
		c.JSON(res.Status, res)
	}
}

func (ac *ArticleController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.PostForm("title")
	lines := strings.Split(c.PostForm("lines"), "&divide;,")
	if article, err := ac.ArticleRepository.FindByID(id); err != nil {
		res := response.NotFound("投稿が見つかりませんでした(This article was not found)")
		c.JSON(res.Status, res)
	} else {
		if err := ac.ArticleRepository.Update(article, title); err != nil {
			res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
			c.JSON(res.Status, res)
		}
		articleLines := ac.ArticleLineRepository.ArticleLinesOfArticle(id)
		for n, articleLine := range articleLines {
			if err := ac.ArticleLineRepository.Update(articleLine, lines[n]); err != nil {
				res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
				c.JSON(res.Status, res)
			}
		}
		res := response.SuccessResponse("")
		res.Data = lines
		c.JSON(res.Status, res)
	}
}

func (ac *ArticleController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	articleLines := ac.ArticleLineRepository.ArticleLinesOfArticle(id)
	boldNumbers := ac.BoldNumberRepository.BoldNumbersOfArticle(id)
	linkNumbers := ac.LinkNumberRepository.LinkNumbersOfArticle(id)
	if article, err := ac.ArticleRepository.FindByID(id); err != nil {
		res := response.NotFound("投稿が見つかりませんでした(This article was not found)")
		c.JSON(res.Status, res)
	} else if err := ac.ArticleRepository.Delete(article); err != nil {
		res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
		c.JSON(res.Status, res)
	} else {
		for i := 0; i < len(articleLines); i++ {
			if err := ac.ArticleLineRepository.Delete(articleLines[i]); err != nil {
				res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
				c.JSON(res.Status, res)
			}
		}
		for i := 0; i < len(boldNumbers); i++ {
			if err := ac.BoldNumberRepository.Delete(boldNumbers[i]); err != nil {
				res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
				c.JSON(res.Status, res)
			}
		}
		for i := 0; i < len(linkNumbers); i++ {
			if err := ac.LinkNumberRepository.Delete(linkNumbers[i]); err != nil {
				res := response.BadRequest("予期せぬエラーが発生しました(An unexpected error has occured)")
				c.JSON(res.Status, res)
			}
		}
	}
	res := response.SuccessResponse("")
	c.JSON(res.Status, res)
}

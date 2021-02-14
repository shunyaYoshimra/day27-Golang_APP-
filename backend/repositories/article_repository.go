package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

type ArticleRepository struct {
	Conn *gorm.DB
}

type ArticleLineRepository struct {
	Conn *gorm.DB
}

type BoldNumberRepository struct {
	Conn *gorm.DB
}

type LinkNumberRepository struct {
	Conn *gorm.DB
}

func NewArticleRepository() ArticleRepository {
	return ArticleRepository{Conn: database.GetDB().Table("articles")}
}

func NewArticleLineRepository() ArticleLineRepository {
	return ArticleLineRepository{Conn: database.GetDB().Table("article_lines")}
}

func NewBoldNumberRepository() BoldNumberRepository {
	return BoldNumberRepository{Conn: database.GetDB().Table("bold_numbers")}
}

func NewLinkNumberRepository() LinkNumberRepository {
	return LinkNumberRepository{Conn: database.GetDB().Table("link_numbers")}
}

func (ar *ArticleRepository) RetrieveArticles() (articles []entity.Article) {
	ar.Conn.Order("created_at").Find(&articles)
	return
}

func (ar *ArticleLineRepository) ArticlesOfUser(id int) (articles []entity.Article) {
	ar.Conn.Where("user_id = ?", id).Order("created_at").Find(&articles)
	return
}

func (ar *ArticleLineRepository) FindByID(id int) (article entity.Article, err error) {
	err = ar.Conn.First(article, id).Error
	return
}

func (alr *ArticleLineRepository) ArticleLinesOfArticle(id int) (articleLines []entity.ArticleLine) {
	alr.Conn.Where("article_id = ?", id).Find(&articleLines)
	return
}

func (br *BoldNumberRepository) BoldNumbersOFArticle(id int) (boldNumbers []entity.BoldNumber) {
	br.Conn.Where("article_id = ?", id).Find(&boldNumbers)
	return
}

func (lr *LinkNumberRepository) LinkNumbersOfArticle(id int) (LinkNumbers []entity.LinkNumber) {
	lr.Conn.Where("article_id = ?", id).Find(&LinkNumbers)
	return
}

// functions for creating articles

func (ar *ArticleRepository) Create(article *entity.Article) (err error) {
	err = ar.Conn.Create(article).Error
	return
}

func (alr *ArticleLineRepository) Create(articleLine *entity.ArticleLine) (err error) {
	err = alr.Conn.Create(articleLine).Error
	return
}

func (br *BoldNumberRepository) Create(boldNumber *entity.BoldNumber) (err error) {
	err = br.Conn.Create(boldNumber).Error
	return
}

func (lr *LinkNumberRepository) Create(linkNumber *entity.LinkNumber) (err error) {
	err = lr.Conn.Create(linkNumber).Error
	return
}

// functions for updating articles

func (ar *ArticleRepository) Update(article entity.Article, title string) (err error) {
	err = ar.Conn.Model(&article).Update(entity.Article{Title: title}).Error
	return
}

func (alr *ArticleLineRepository) Update(articleLine entity.ArticleLine, line string) (err error) {
	err = alr.Conn.Model(&articleLine).Update(entity.ArticleLine{Line: line}).Error
	return
}

func (br *BoldNumberRepository) Update(boldNumber entity.BoldNumber, number int) (err error) {
	err = br.Conn.Model(&boldNumber).Update(entity.BoldNumber{Number: number}).Error
	return
}

func (lr *LinkNumberRepository) Update(linkNumber entity.LinkNumber, number int) (err error) {
	err = lr.Conn.Model(&linkNumber).Update(entity.LinkNumber{Number: number}).Error
	return
}

// functions for deleting articles

func (ar *ArticleRepository) Delete(article entity.Article) (err error) {
	err = ar.Conn.Delete(article).Error
	return
}

func (alr *ArticleLineRepository) Delete(articleLine entity.ArticleLine) (err error) {
	err = alr.Conn.Delete(articleLine).Error
	return
}

func (br *BoldNumberRepository) Delete(boldNumber entity.BoldNumber) (err error) {
	err = br.Conn.Delete(boldNumber).Error
	return
}

func (lr *LinkNumberRepository) Delete(linkNumber entity.LinkNumber) (err error) {
	err = lr.Conn.Delete(linkNumber).Error
	return
}

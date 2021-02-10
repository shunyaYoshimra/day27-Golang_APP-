package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

type PostRepository struct {
	Conn *gorm.DB
}

type ImageRepository struct {
	Conn *gorm.DB
}

func NewPostRepository() PostRepository {
	return PostRepository{Conn: database.GetDB().Table("posts")}
}

func NewImageRepository() ImageRepository {
	return ImageRepository{Conn: database.GetDB().Table("images")}
}

// functions for PostRepository

func (pr *PostRepository) RetrievePosts() (posts []entity.Post) {
	pr.Conn.Order("created_at desc").Find(&posts)
	return
}

func (pr *PostRepository) PostsByUser(id int) (posts []entity.Post) {
	pr.Conn.Where("user_id = ?", id).Order("created_at desc").Find(&posts)
	return
}

func (pr *PostRepository) PostsByKeyword(keyword string) (posts []entity.Post) {
	pr.Conn.Where("keyword = ?", keyword).Order("created_at desc").Find(&posts)
	return
}

func (pr *PostRepository) FindByID(id int) (post []entity.Post, err error) {
	err = pr.Conn.First(&post, id).Error
	return
}

func (pr *PostRepository) Create(post *[]entity.Post) (err error) {
	err = pr.Conn.Create(post).Error
	return
}

func (pr *PostRepository) Delete(post []entity.Post) (err error) {
	err = pr.Conn.Delete(&post).Error
	return
}

// functions for ImageRepository

func (ir *ImageRepository) RetrieveImages() (images []entity.Image) {
	ir.Conn.Order("crated_at desc").Find(&images)
	return
}

func (ir *ImageRepository) Create(image *[]entity.Image) (err error) {
	err = ir.Conn.Create(image).Error
	return
}

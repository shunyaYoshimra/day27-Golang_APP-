package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/shunyaYoshimra/day27/backend/database"
	"github.com/shunyaYoshimra/day27/backend/database/entity"
)

type FavoriteRepository struct {
	Conn *gorm.DB
}

func NewFavoriteRepository() FavoriteRepository {
	return FavoriteRepository{Conn: database.GetDB().Table("favorites")}
}

func (fr *FavoriteRepository) FavoritesOfPost(id int) (favorites []entity.Favorite) {
	fr.Conn.Where("post_id = ?", id).Find(&favorites)
	return
}

func (fr *FavoriteRepository) FavroitesOfUser(id int) (favorites []entity.Favorite) {
	fr.Conn.Where("user_id = ?", id).Find(&favorites)
	return
}

func (fr *FavoriteRepository) CheckUnique(postID, userID int) (err error) {
	var f entity.Favorite
	err = fr.Conn.Where("post_id = ? AND user_id = ?", postID, userID).First(&f).Error
	return
}

func (fr *FavoriteRepository) FindByID(id int) (favorite entity.Favorite, err error) {
	fr.Conn.First(&favorite, id)
	return
}

func (fr *FavoriteRepository) Create(favorite *entity.Favorite) (err error) {
	err = fr.Conn.Create(favorite).Error
	return
}

func (fr *FavoriteRepository) Delete(favorite entity.Favorite) (err error) {
	err = fr.Conn.Delete(&favorite).Error
	return
}

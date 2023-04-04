package service

import (
	"errors"
	"my-admin/global"
	"my-admin/model"

	"gorm.io/gorm"
)

func CreatePhoto(photo model.Photo) error {
	return insertPhoto(photo)
}

func DeletePhoto(photo model.Photo) error {
	return deletePhoto(photo)
}

func insertPhoto(phote model.Photo) error {
	return global.DBClient.Create(&phote).Error
}

func deletePhoto(photo model.Photo) error {
	return global.DBClient.Delete(photo).Error
}

func isPhotoOutOfCapacity() bool {
	return false
}

func ListPhoto() []model.Photo {
	var photos []model.Photo
	global.DBClient.Find(&photos)
	return photos
}

func GetPhotoByUUID(uuid string) (model.Photo, bool) {
	var photo model.Photo
	result := global.DBClient.Where("uuid = ?", uuid).First(&photo)
	return photo, !errors.Is(result.Error, gorm.ErrRecordNotFound)
}

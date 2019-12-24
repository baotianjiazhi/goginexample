package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

type Tag struct {
	gorm.Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	Status     string `json:"status"`
}

func ExistByName(name string) (bool, error) {
	var tag Tag
	err := DB.Select("id").Where("name = ? And status = ?", name, Active).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	log.Println(tag)
	if tag.ID > 0 {
		err := errors.New("tag已经存在")
		return true, err
	}

	return false, nil
}

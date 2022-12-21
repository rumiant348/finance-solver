package models

import (
	"github.com/jinzhu/gorm"
)

type listGorm struct {
	db *gorm.DB
}

var _ ListDB = &listGorm{}

func (lg listGorm) ByID(id uint) (*List, error) {
	var l List
	db := lg.db.Where("id = ?", id)
	err := First(db, &l)
	if err != nil {
		return nil, err
	}
	return &l, err
}

func (lg listGorm) ByUserID(id uint) ([]List, error) {
	var l []List
	db := lg.db.Where("user_id = ?", id)
	err := db.Find(&l).Error
	if err != nil {
		return nil, err
	}
	return l, err
}

func (lg listGorm) Create(list *List) error {
	return lg.db.Create(list).Error
}

func (lg listGorm) Update(list *List) error {
	return lg.db.Save(list).Error
}

func (lg listGorm) Delete(id uint) error {
	l := List{
		Model: gorm.Model{ID: id},
	}
	return lg.db.Delete(&l).Error
}

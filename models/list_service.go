package models

import "github.com/jinzhu/gorm"

type ListService interface {
	ListDB
}

type listService struct {
	ListDB
}

func NewListService(db *gorm.DB) ListService {
	return &listService{
		&listValidator{
			&listGorm{db: db},
		},
	}
}

var _ ListService = &listService{}

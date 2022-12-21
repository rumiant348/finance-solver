package models

import (
	"github.com/jinzhu/gorm"
)

func NewServices(cfgs ...ServicesConfig) (*Services, error) {

	var s Services

	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}

	return &s, nil
}

type Services struct {
	List ListService
	User UserService
	db   *gorm.DB
}

type ServicesConfig func(*Services) error

func WithGorm(dialect, connectionInfo string) ServicesConfig {
	return func(s *Services) error {
		db, err := gorm.Open(dialect, connectionInfo)
		if err != nil {
			return err
		}
		s.db = db
		return nil
	}
}

func WithLogMode(mode bool) ServicesConfig {
	return func(s *Services) error {
		s.db.LogMode(mode)
		return nil
	}
}

func WithUser(pepper, hmacKey string) ServicesConfig {
	return func(s *Services) error {
		s.User = NewUserService(s.db, pepper, hmacKey)
		return nil
	}
}

func WithList() ServicesConfig {
	return func(s *Services) error {
		s.List = NewListService(s.db)
		return nil
	}
}

// Close the database connection
func (s *Services) Close() error {
	return s.db.Close()
}

// AutoMigrate will attempt to automatically migrate all tables
func (s *Services) AutoMigrate() error {
	return s.db.AutoMigrate(&User{}, &List{}).Error
}

// DestructiveReset drops all tables and rebuilds them
func (s *Services) DestructiveReset() error {
	err := s.db.DropTableIfExists(&User{}, &List{}).Error
	if err != nil {
		return err
	}
	return s.AutoMigrate()
}

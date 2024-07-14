package config

import "gorm.io/gorm"

var (
	database *gorm.DB
)

func Init() error {
	return nil
}
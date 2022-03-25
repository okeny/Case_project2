package database

import "gorm.io/gorm"

type DatabaseInterface interface {
	GetConnection() *gorm.DB
}

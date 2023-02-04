package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	db  *gorm.DB
	log *logrus.Entry
}

func NewRepository(db *gorm.DB, log *logrus.Entry) *Repository {
	return &Repository{
		db:  db,
		log: log,
	}
}

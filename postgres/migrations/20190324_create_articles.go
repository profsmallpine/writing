package migrations

import (
	"github.com/jinzhu/gorm"
)

func (ArticlesTable) TableName() string {
	return "articles"
}

type ArticlesTable struct {
	gorm.Model
	Title   string `sql:"not null"`
	Content string `sql:"not null"`
}

func CreateArticlesTable(db *gorm.DB) error {
	return db.CreateTable(&ArticlesTable{}).Error
}

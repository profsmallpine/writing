package migrations

import "gorm.io/gorm"

type ArticlesTable struct {
	gorm.Model
	Body    string `gorm:"not null"`
	Slug    string `gorm:"not null;index"`
	Summary string `gorm:"not null"`
	Title   string `gorm:"not null"`
}

func (ArticlesTable) TableName() string {
	return "articles"
}

func createArticlesTable(tx *gorm.DB) error {
	return tx.AutoMigrate(&ArticlesTable{})
}

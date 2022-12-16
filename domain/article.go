package domain

import (
	"html/template"

	"gorm.io/gorm"
)

const ArticlePerPage = 5

type Article struct {
	gorm.Model
	Body    string
	Slug    string
	Summary string
	Title   string
}

func (a *Article) CreatedAtDate() string {
	return a.CreatedAt.Format("1/2/2006")
}

func (a *Article) BodyHTML() template.HTML {
	return template.HTML(a.Body)
}

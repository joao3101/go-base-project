package mysql

import (
	"gorm.io/gorm"

	"github.com/joao3101/go-base-project/domain"
)

type ArticleRepository struct {
	Conn *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewArticleRepository(Conn *gorm.DB) domain.ArticleRepository {
	return &ArticleRepository{Conn}
}

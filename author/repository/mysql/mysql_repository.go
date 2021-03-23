package mysql

import (
	"gorm.io/gorm"

	"github.com/joao3101/go-base-project/domain"
)

type AuthorRepo struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) domain.AuthorRepository {
	return &AuthorRepo{
		DB: db,
	}
}

func (m *AuthorRepo) GetByID(id int64) (res domain.Author, err error) {
	err = m.DB.First(&res).Error
	if err != nil {
		return domain.Author{}, err
	}
	return
}

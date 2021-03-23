package domain

// Author ...
type Author struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// AuthorRepository represent the author's repository contract
type AuthorRepository interface {
	GetByID(id int64) (Author, error)
}

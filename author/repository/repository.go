package repository

import "github.com/nenjotsu/go-clean-arch/author"

type AuthorRepository interface {
	GetByID(id int64) (*author.Author, error)
}

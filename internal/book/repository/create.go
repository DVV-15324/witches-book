package repository

import (
	"context"
	domainBook "new_example/internal/shared/domain"
	"new_example/internal/book/mapping"
)

func (r *BookRepository) Create(ctx context.Context, e *domainBook.Book) (error) {
	model := mapping.FromDomainToModelBook(e)
	err := r.db.WithContext(ctx).Create(model).Error
	return err
}
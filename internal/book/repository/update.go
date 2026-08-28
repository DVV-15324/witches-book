package repository

import (
	"context"
	"new_example/internal/book/mapping"
	domainBook "new_example/internal/shared/domain"
)

func (r *BookRepository) Update(ctx context.Context, e *domainBook.Book, id int) error {
	model := mapping.FromDomainToModelBook(e)
	return r.db.WithContext(ctx).Save(model).Error
}
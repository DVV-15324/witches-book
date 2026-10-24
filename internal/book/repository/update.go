package repository

import (
	"context"
	"github.com/DVV-15324/witches-book/internal/book/mapping"
	domainBook "github.com/DVV-15324/witches-book/pkg/shared/domain"
)

func (r *BookRepository) Update(ctx context.Context, e *domainBook.Book, id int) error {
	model := mapping.FromDomainToModelBook(e)
	return r.core.DB.WithContext(ctx).Save(model).Error
}

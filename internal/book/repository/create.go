package repository

import (
	"context"
	"github.com/DVV-15324/witches-book/internal/book/mapping"
	domainBook "github.com/DVV-15324/witches-book/pkg/shared/domain"
)

func (r *BookRepository) Create(ctx context.Context, e *domainBook.Book) error {
	model := mapping.FromDomainToModelBook(e)
	err := r.core.DB.WithContext(ctx).Create(model).Error
	return err
}

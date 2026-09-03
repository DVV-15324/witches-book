package repository

import (
	"context"
	modelBook "github.com/DVV-15324/witches-book/internal/book/model"
)

func (r *BookRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&modelBook.Book{}, id).Error
}

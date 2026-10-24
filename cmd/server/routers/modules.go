package routers

import (
	"github.com/DVV-15324/witches-book/cmd/server/core"
	"github.com/DVV-15324/witches-book/internal/book"
)

type Modules struct {
	Book *book.BookModule
}

func InitModules(core *core.CoreServices) *Modules {
	bookModule := book.NewBookModule(core)

	return &Modules{Book: bookModule}
}

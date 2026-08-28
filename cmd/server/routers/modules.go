package routers

import (
	"new_example/cmd/server/core"
	"new_example/internal/book"
)

type Modules struct {
	Book *book.BookModule
}

func InitModules(core *core.CoreServices) *Modules {
	bookModule := book.NewBookModule(core)

	return &Modules{Book: bookModule}
}

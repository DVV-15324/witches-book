package routers

import (
	"github.com/DVV-15324/witches-book/cmd/server/core"
)

type Modules struct {
}

func InitModules(core *core.CoreServices) *Modules {

	return &Modules{}
}

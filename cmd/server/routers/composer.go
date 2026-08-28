package routers

import (
	wcmd_utils "github.com/DVV-15324/witches/cmd/utils"
	"github.com/gin-gonic/gin"
	"new_example/cmd/server/core"
)

func Start(r *gin.Engine, cfg *wcmd_utils.Config) {
	core := core.NewCoreServices(cfg)
	modules := InitModules(core)
	RegisterRoutes(r, core, modules)
}

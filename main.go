package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/linux-http/api"
	"github.com/lwinmgmg/linux-http/env"
	"github.com/lwinmgmg/linux-http/middleware"
	"github.com/lwinmgmg/linux-http/services"
)

var Env env.Settings = env.NewEnv()

func main() {
	app := gin.New()
	app.Use(gin.CustomRecovery(middleware.PanicMiddleware))
	ctrl := api.ControllerV1{
		App: app.Group(""),
		DB:  services.DB,
	}
	ctrl.HandleRoutes()
	app.Run(fmt.Sprintf("%v:%v", Env.LH_HOST, Env.LH_PORT))
}

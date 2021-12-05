package main

import (
	"github.com/unrolled/render"
	"log"
	"net/http"
	"os"
	"zmakers-backend/controllers"
	"zmakers-backend/logger"
	"zmakers-backend/routers"
	"zmakers-backend/service"
)

func main() {
	service.DbInit()
	defer service.DbClose()

	bc := controllers.BaseController{
		Render: render.New(),
		Log: logger.MustNew("backendAPI", os.Getenv("SENTRY_DSN")),
	}

	collectionController := controllers.CollectionController{
		&bc,
	}

	mediaController := controllers.MediaController{
		&bc,
	}

	userController := controllers.UserController{
		&bc,
	}
	err := http.ListenAndServe(":3000", routers.RegisterServices(&collectionController,
		&mediaController, &userController))
	if err != nil {
		log.Fatal("system error ", err)
	}
}

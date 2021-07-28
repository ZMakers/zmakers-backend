package main

import (
	"github.com/unrolled/render"
	"log"
	"net/http"
	"os"
	"zmakers-backend/controllers"
	"zmakers-backend/routers"
	"zmakers-backend/service"
	"zmakers-backend/logger"
)

func main() {
	service.DbInit()
	defer service.DbClose()

	collectionController := controllers.CollectionController{
		Render: render.New(),
		Log: logger.MustNew("collectionAPI", os.Getenv("SENTRY_DSN")),
	}

	err := http.ListenAndServe(":3000", routers.RegisterCollectionService(&collectionController))
	if err != nil {
		log.Fatal("system error ", err)
	}

}
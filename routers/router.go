package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	. "zmakers-backend/controllers"
)


var router *chi.Mux

func RegisterCollectionService(CollectionController *CollectionController) *chi.Mux {
	router = chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/collections", CollectionController.NewCollection)
	return router
}

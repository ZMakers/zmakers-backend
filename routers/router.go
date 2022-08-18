package routers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"os"
	. "zmakers-backend/controllers"
)


var router *chi.Mux = chi.NewRouter()

func RegisterServices(controllers ...interface{}) *chi.Mux {
	router.Use(middleware.Logger)
	for _, controller := range controllers {
		switch controller.(type) {
		case *CollectionController:
			op, ok := controller.(*CollectionController)
			if ok {
				fmt.Println("registerCollectionService")
				RegisterCollectionService(op)
			} else {
				fmt.Println(12312)
				os.Exit(010)
			}
		case *MediaController:
			op, ok := controller.(*MediaController)
			if ok {
				fmt.Println("registerMediaService")
				RegisterMediaService(op)
			} else {
				os.Exit(010)
			}
		case *UserController:
			op, ok := controller.(*UserController)
			if ok {
				fmt.Println("registerUserService")
				RegisterUserService(op)
			} else {
				os.Exit(010)
			}
		case *NftController:
			op, ok := controller.(*NftController)
			if ok {
				fmt.Println("registerNftController")
				RegisterNftService(op)
			} else {
				os.Exit(010)
			}
		}
	}
	return router
}

func RegisterCollectionService(CollectionController *CollectionController) *chi.Mux {
	router.Post("/collections", CollectionController.NewCollection)
	return router
}

func RegisterMediaService(MediaController *MediaController) *chi.Mux {
	router.Post("/media", MediaController.NewMedia)
	router.Post("/media/buy", MediaController.BuyMedia)
	return router
}

func RegisterUserService(UserController *UserController) *chi.Mux {
	router.Post("/user", UserController.NewUser)
	router.Post("/user/login", UserController.Login)
	return router
}

func RegisterNftService(nftController *NftController) *chi.Mux {
	router.Post("/nft", nftController.MintNft)
	router.Post("/nft/tx", nftController.ParseTx)
	return router
}

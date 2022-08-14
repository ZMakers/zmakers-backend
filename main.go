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

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")  // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		f(w, r)
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Golang"))
}

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

	nftController := controllers.NftController{
		&bc,
	}

	http.HandleFunc("/", cors(index))
	err := http.ListenAndServe(":3000", routers.RegisterServices(&collectionController,
		&mediaController, &userController, &nftController))
	if err != nil {
		log.Fatal("system error ", err)
	}
}

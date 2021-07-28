package controllers

import (
	"github.com/unrolled/render"
	"net/http"
	"zmakers-backend/lib"
	"zmakers-backend/logger"
	"zmakers-backend/models"
)

type CollectionController struct {
	Render *render.Render
	Log    *logger.Clog
	Collection models.Collection
}

func (cc *CollectionController) catchError(w http.ResponseWriter, r *http.Request, infos ...interface{}) {
	if err := recover(); err != nil {
		cc.Log.Errorf(r, infos[0].(string), err)
		cc.Render.JSON(w, http.StatusBadRequest, infos[1])
		return
	}
	cc.Render.JSON(w, http.StatusOK, infos[2])
}

func (cc *CollectionController) NewCollection(w http.ResponseWriter, r *http.Request) {
	queryValues := r.Header
	creator := queryValues.Get("creator")
	metaPath := queryValues.Get("metaPath")

	if !lib.IsValidAddress(creator) {
		cc.Log.Errorf(r, "Invalid request data %v", cc.Collection)
		cc.Render.JSON(w, http.StatusBadRequest, "Invalid request data")
	}

	cc.Collection = models.Collection{
		Owner:    creator,
		MetaPath: metaPath,
	}

	infos := [3]interface{}{"Create collection error %v\n", "Failed to create collection", cc.Collection}
	defer cc.catchError(w, r, infos)
	cc.Collection.CreateCollection()
}


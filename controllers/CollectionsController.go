package controllers

import (
	"encoding/json"
	"github.com/unrolled/render"
	"net/http"
	"zmakers-backend/lib"
	"zmakers-backend/logger"
	"zmakers-backend/models"
)

type CollectionController struct {
	Render *render.Render
	Log    *logger.Clog
	Collection *models.Collection
}

func (cc *CollectionController) catchError(w http.ResponseWriter, r *http.Request, infos []interface{}) {
	if err := recover(); err != nil {
		cc.Log.Errorf(r, infos[0].(string), err)
		cc.Render.JSON(w, http.StatusBadRequest, infos[1])
		return
	}
	cc.Render.JSON(w, http.StatusOK, infos[2])
}

func (cc *CollectionController) NewCollection(w http.ResponseWriter, r *http.Request) {
	var payload models.Collection
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		cc.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request. Incorrect post body"})
		return
	}
	if  payload.Owner == "" || payload.MetaPath == "" || payload.TxHash == "" {
		cc.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request. Incorrect MetaPath/Owner/TxHash"})
		return
	}

	if !lib.IsValidAddress(payload.Owner) {
		cc.Log.Errorf(r, "Invalid request data %v", cc.Collection)
		cc.Render.JSON(w, http.StatusBadRequest, "Invalid request data")
	}

	cc.Collection = &payload

	infos := make([]interface{}, 3)
	infos[0] = "Create collection error\n"
	infos[1] = "Failed to create collection"
	infos[2] = *cc.Collection

	defer cc.catchError(w, r, infos)

	cc.Collection.CreateCollection()
}


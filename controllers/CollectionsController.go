package controllers

import (
	"encoding/json"
	"net/http"
	"zmakers-backend/lib"
	"zmakers-backend/models"
)

type CollectionController struct {
	BC	*BaseController
}


func (cc *CollectionController) catchError(w http.ResponseWriter, r *http.Request, infos []interface{}) {
	if err := recover(); err != nil {
		cc.BC.Log.Errorf(r, infos[0].(string), err)
		cc.BC.Render.JSON(w, http.StatusBadRequest, infos[1])
		return
	}
	cc.BC.Render.JSON(w, http.StatusOK, infos[2])
}

func (cc *CollectionController) NewCollection(w http.ResponseWriter, r *http.Request) {

	err := cc.BC.CheckToken(r)
	if err != nil {
		cc.BC.Render.JSON(w, http.StatusBadRequest, map[string]string{"token error": err.Error()})
		return
	}

	var payload models.Collection
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		cc.BC.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request. Incorrect post body "+err.Error()})
		return
	}
	if  payload.Owner == "" || payload.MetaPath == "" || payload.TxHash == "" {
		cc.BC.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request. Incorrect MetaPath/Owner/TxHash"})
		return
	}

	if !lib.IsValidAddress(payload.Owner) {
		cc.BC.Log.Errorf(r, "Invalid request data %v", payload.Owner)
		cc.BC.Render.JSON(w, http.StatusBadRequest, "Invalid request data, owner address is not valid")
		return
	}

	if !lib.IsValidTxHash(payload.TxHash) {
		cc.BC.Log.Errorf(r, "Invalid request data %v", payload.TxHash)
		cc.BC.Render.JSON(w, http.StatusBadRequest, "Invalid request data, tx hash is not valid")
		return
	}

	infos := make([]interface{}, 3)
	infos[0] = "Create collection error\n"
	infos[1] = "Failed to create collection"
	infos[2] = payload

	defer cc.catchError(w, r, infos)
	payload.CreateCollection()
}

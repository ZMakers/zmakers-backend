package controllers

import (
	"encoding/json"
	"net/http"
	"zmakers-backend/lib"
	"zmakers-backend/models"
	"zmakers-backend/models/utility"
)

type CollectionController struct {
	BC	*BaseController
}


func (cc *CollectionController) catchError(w http.ResponseWriter, r *http.Request, infos []interface{}) {
	var jsonResult utility.Response
	if err := recover(); err != nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = infos[1]
		cc.BC.Log.Errorf(r, infos[0].(string), err)
		cc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	jsonResult.Code = http.StatusOK
	jsonResult.Data = infos[2]
	cc.BC.Render.JSON(w, http.StatusOK, jsonResult)
}

func (cc *CollectionController) NewCollection(w http.ResponseWriter, r *http.Request) {

	err := cc.BC.CheckToken(r)
	var jsonResult utility.Response
	if err != nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"token error": err.Error()}
		cc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}

	var payload models.Collection
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "Bad request. Incorrect post body "+err.Error()}
		cc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	if  payload.Owner == "" || payload.MetaPath == "" || payload.TxHash == "" {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "Bad request. Incorrect MetaPath/Owner/TxHash"}
		cc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}

	if !lib.IsValidAddress(payload.Owner) {
		cc.BC.Log.Errorf(r, "Invalid request data %v", payload.Owner)
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = "Invalid request data, owner address is not valid"
		cc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}

	if !lib.IsValidTxHash(payload.TxHash) {
		cc.BC.Log.Errorf(r, "Invalid request data %v", payload.TxHash)
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = "Invalid request data, tx hash is not valid"
		cc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}

	infos := make([]interface{}, 3)
	infos[0] = "Create collection error\n"
	infos[1] = "Failed to create collection"
	infos[2] = payload

	defer cc.catchError(w, r, infos)
	payload.CreateCollection()
}

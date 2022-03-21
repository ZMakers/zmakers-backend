package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zmakers-backend/models"
	"zmakers-backend/models/utility"
)

type MediaController struct {
	BC    *BaseController
}

func (mc *MediaController) catchError(w http.ResponseWriter, r *http.Request, infos []interface{}) {
	var jsonResult utility.Response
	if err := recover(); err != nil {
		mc.BC.Log.Errorf(r, infos[0].(string), err)
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = infos[1]
		mc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	jsonResult.Code = http.StatusOK
	jsonResult.Data = infos[2]
	mc.BC.Render.JSON(w, http.StatusOK, infos[2])
}

func (mc *MediaController) NewMedia(w http.ResponseWriter, r *http.Request) {
	var jsonResult utility.Response
	err := mc.BC.CheckToken(r)
	if err != nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"token error": err.Error()}
		mc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}

	var payload models.Media
	fmt.Println(r.Body)
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "Bad request. Incorrect post body "+err.Error()}
		mc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}

	infos := make([]interface{}, 3)
	infos[0] = "Create Media error\n"
	infos[1] = "Failed to create media"
	infos[2] = payload

	defer mc.catchError(w,r,infos)
	payload.CreateMedia()
}

func (mc *MediaController) BuyMedia(w http.ResponseWriter, r *http.Request) {
	err := mc.BC.CheckToken(r)
	var jsonResult utility.Response
	if err != nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"token error": err.Error()}
		mc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}

	var payload models.Transfer
	fmt.Println(r.Body)
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "Bad request. Incorrect post body "+err.Error()}
		mc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}

	infos := make([]interface{}, 3)
	infos[0] = "Buy Media error\n"
	infos[1] = "Failed to buy media"
	infos[2] = payload

	defer mc.catchError(w,r,infos)
	payload.BuyMedia()
}


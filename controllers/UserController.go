package controllers

import (
	"encoding/json"
	"net/http"
	"zmakers-backend/models"
	"zmakers-backend/models/utility"
)

type UserController struct {
	BC    *BaseController
}

func (uc *UserController) NewUser (w http.ResponseWriter, r *http.Request)  {
	var payload models.User
	var jsonResult utility.Response
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		uc.BC.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request. Incorrect post body "+err.Error()})
		return
	}
	err = payload.CreatUser()
	if err != nil {
		jsonResult.Code = http.StatusInternalServerError
		jsonResult.Data = map[string]string{"error": err.Error()}
		uc.BC.Render.JSON(w, http.StatusInternalServerError, jsonResult)
		return
	}
	jsonResult.Code = http.StatusOK
	jsonResult.Data = payload
	uc.BC.Render.JSON(w, http.StatusOK, jsonResult)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var payload models.User
	var jsonResult utility.Response
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "Bad request. Incorrect post body "+err.Error()}
		uc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	user := payload.SearchUser()
	if user == nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "User not found!"}
		uc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	err = user.GenerateToken()
	if err != nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "Generate user token error "+ err.Error()}
		uc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	user = user.UpdateUserToken()
	if user == nil {
		jsonResult.Code = http.StatusBadRequest
		jsonResult.Data = map[string]string{"error": "Update User error"}
		uc.BC.Render.JSON(w, http.StatusBadRequest, jsonResult)
		return
	}
	jsonResult.Code = http.StatusOK
	jsonResult.Data = user
	uc.BC.Render.JSON(w, http.StatusOK, jsonResult)
}


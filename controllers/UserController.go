package controllers

import (
	"encoding/json"
	"net/http"
	"zmakers-backend/models"
)

type UserController struct {
	BC    *BaseController
}

func (uc *UserController) NewUser (w http.ResponseWriter, r *http.Request)  {
	var payload models.User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		uc.BC.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request. Incorrect post body "+err.Error()})
		return
	}
	err = payload.CreatUser()
	if err != nil {
		uc.BC.Render.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to register user!"})
		return
	}
	uc.BC.Render.JSON(w, http.StatusOK, payload)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var payload models.User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		uc.BC.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request. Incorrect post body "+err.Error()})
		return
	}
	user := payload.SearchUser()
	if user == nil {
		uc.BC.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "User unfound!"})
		return
	}
	err = user.GenerateToken()
	if err != nil {
		uc.BC.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Generate user token error"})
		return
	}
	user = user.UpdateUserToken()
	if user == nil {
		uc.BC.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Update User error"})
		return
	}
	uc.BC.Render.JSON(w, http.StatusOK, user)
}


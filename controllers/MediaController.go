package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zmakers-backend/models"
)

type MediaController struct {
	BC    *BaseController
}

func (mc *MediaController) NewMedia(w http.ResponseWriter, r *http.Request) {

	err := mc.BC.CheckToken(r)
	if err != nil {
		mc.BC.Render.JSON(w, http.StatusBadRequest, map[string]string{"token error": err.Error()})
		return
	}

	var payload models.Media
	fmt.Println(r.Body)
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		mc.BC.Render.JSON(w, http.StatusBadRequest, map[string]string{"error": "Bad request. Incorrect post body "+err.Error()})
		return
	}

	payload.CreateMedia()
}

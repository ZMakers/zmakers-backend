package controllers

import (
	"github.com/unrolled/render"
	"net/http"
	_interface "zmakers-backend/controllers/interface"
	"zmakers-backend/logger"
	"zmakers-backend/models"
)

type BaseController struct {
	Render *render.Render
	Log    *logger.Clog
	_interface.IBaseController
}

func (BC *BaseController) CheckToken(r *http.Request)  error {
	token := r.Header.Get("token")
	user := models.User{
		Token:    token,
	}
	return user.VerifyToken(token)
}
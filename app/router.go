package app

import (
	"halo-suster/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userCtl controller.UserCtl) *httprouter.Router {
	router := httprouter.New()

	router.POST("/v1/user/it/register", userCtl.Register)
	router.POST("/v1/user/it/login", userCtl.Login)

	return router
}

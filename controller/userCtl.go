package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserCtl interface {
	Register(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

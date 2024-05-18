package controller

import (
	"halo-suster/helper"
	"halo-suster/model/web"
	"halo-suster/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserCtlImpl struct {
	UserSvc service.UserSvc
}

func NewUserCtl(userSvc service.UserSvc) UserCtl {
	return &UserCtlImpl{
		UserSvc: userSvc,
	}
}

func (controller *UserCtlImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userCreateRequest := web.UserRegisterReq{}
	helper.ReadFromRequestBody(r, &userCreateRequest)
	userResponse, err := controller.UserSvc.Register(r.Context(), userCreateRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		webResponse := web.WebResponse{
			Message: err.Error(),
		}
		helper.WriteToResponseBody(w, webResponse)
	} else {
		w.WriteHeader(http.StatusCreated)
		webResponse := web.WebResponse{
			Message: "User registered successfully",
			Data:    userResponse,
		}
		helper.WriteToResponseBody(w, webResponse)
	}
}

func (controller *UserCtlImpl) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userLoginRequest := web.UserLoginReq{}
	helper.ReadFromRequestBody(r, &userLoginRequest)
	userResponse, err := controller.UserSvc.Login(r.Context(), userLoginRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		webResponse := web.WebResponse{
			Message: err.Error(),
		}
		helper.WriteToResponseBody(w, webResponse)
	} else {
		w.WriteHeader(http.StatusOK)
		webResponse := web.WebResponse{
			Message: "User logged successfully",
			Data:    userResponse,
		}
		helper.WriteToResponseBody(w, webResponse)
	}
}

package service

import (
	"context"
	"halo-suster/model/web"
)

type UserSvc interface {
	Register(ctx context.Context, request web.UserRegisterReq) (web.UserRes, error)
	Login(ctx context.Context, request web.UserLoginReq) (web.UserRes, error)
}

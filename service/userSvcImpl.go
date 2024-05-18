package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"halo-suster/config"
	"halo-suster/helper"
	"halo-suster/model/domain"
	"halo-suster/model/web"
	"halo-suster/pkg/errorwrapper"
	"halo-suster/repository"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserSvcImpl struct {
	UserRepo repository.UserRepo
	DB       *sql.DB
	Validate *validator.Validate
}

func NewUserSvc(userRepository repository.UserRepo, DB *sql.DB, validate *validator.Validate) UserSvc {
	return &UserSvcImpl{
		UserRepo: userRepository,
		DB:       DB,
		Validate: validate,
	}
}

func (service *UserSvcImpl) GenerateToken(ctx context.Context, tx *sql.Tx, user domain.User) (string, error) {
	expTime := time.Now().Add(time.Hour * 8)
	claims := &config.JWTClaim{
		Name:   user.Name,
		UserId: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(os.Getenv("JWT_SECRET"))
	token, err := generateToken.SignedString(key)
	if err != nil {
		return "", errors.New("error generate token")
	}
	// err = service.SessionRepo.Save(ctx, tx, user, token)
	// if err != nil {
	// 	return "", errors.New("error save session")
	// }

	return token, nil
}

func (service *UserSvcImpl) Register(ctx context.Context, request web.UserRegisterReq) (web.UserRes, error) {
	err := service.Validate.Struct(request)
	fmt.Println("test")
	if err != nil {
		fmt.Println(err)
		return web.UserRes{}, err
	}
	if !isNipValid(request.Nip) {
		return web.UserRes{}, errorwrapper.New(http.StatusBadRequest, errors.New("invalid NIP"), "")
	}
	tx, err := service.DB.Begin() // transaction db
	if err != nil {
		fmt.Println(err)
		return web.UserRes{}, errorwrapper.New(http.StatusInternalServerError, err, "error database transaction")
	}
	defer helper.CommitOrRollback(tx)
	// hash password
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return web.UserRes{}, errorwrapper.New(http.StatusInternalServerError, errors.New("error generate password"), "")
	}
	request.Password = string(bytes)

	user := domain.User{
		Nip:      request.Nip,
		Password: request.Password,
		Name:     request.Name,
	}
	user, err = service.UserRepo.Save(ctx, tx, user)
	if err != nil {
		fmt.Println(err)
		return web.UserRes{}, errorwrapper.New(http.StatusInternalServerError, errors.New("error save data to database"), "")
	}
	fmt.Println("user:", user.Id)
	token, err := service.GenerateToken(ctx, tx, user)
	if err != nil {
		fmt.Println(err)
		return web.UserRes{}, errorwrapper.New(http.StatusInternalServerError, errors.New("error generate token"), "")
	}
	return helper.ToCategoryResponseUser(user, token), nil
}

func (service *UserSvcImpl) Login(ctx context.Context, request web.UserLoginReq) (web.UserRes, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		fmt.Println(err)
		return web.UserRes{}, err
	}
	helper.PanicIfError(err)
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	user, err := service.UserRepo.FindByNip(ctx, tx, request.Nip)
	if err != nil {
		fmt.Println(err)
		return web.UserRes{}, errorwrapper.New(http.StatusNotFound, err, "")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		fmt.Println(err)
		return web.UserRes{}, errorwrapper.New(http.StatusBadRequest, errors.New("invalid password"), "")
	}
	token, err := service.GenerateToken(ctx, tx, user)
	if err != nil {
		fmt.Println(err)
		return web.UserRes{}, errorwrapper.New(http.StatusInternalServerError, err, "")
	}
	return helper.ToCategoryResponseUser(user, token), nil
}

func isNipValid(phoneNumber string) bool {
	// This is a simple regex for validating an international phone number, which allows for country codes starting with '+'
	// followed by up to 15 digits. This may not cover all possible international phone number formats.
	// You may need to adjust this regex to suit your specific needs.
	regex := `^\+[1-9]{1}[0-9]{9,15}$`
	match, _ := regexp.MatchString(regex, phoneNumber)
	return match
}

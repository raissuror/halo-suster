package main

import (
	// "fmt"
	// "halo-suster/app"
	// "halo-suster/controller"
	// "halo-suster/exception"
	// "halo-suster/repository"
	// "halo-suster/service"
	// "net/http"

	// "github.com/go-playground/validator/v10"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// err := godotenv.Load(".env")
	// validate := validator.New()

	// app.ConnectToPostgres()
	// db, err := app.ConnectToPostgres()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//staff
	// staffRepository := repository.NewStaffRepo()
	// staffService := service.NewStaffService(staffRepository, db, validate)
	// staffController := controller.NewStaffController(staffService)

	// router := app.NewRouter()
	// router.PanicHandler = exception.ErrorHandler
	// server := http.Server{
	// 	Addr:    "localhost:8080",
	// 	Handler: router,
	// }
	// err = server.ListenAndServe()
}

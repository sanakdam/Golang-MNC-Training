package main

import (
	"fmt"
	"net/http"
	"time"

	"Assignment_6/service"

	"github.com/gorilla/mux"
)

func main() {
	var db []*service.User
	userSvc := service.NewUserService(db)
	fmt.Println("http server running on localhost:8080")

	r := mux.NewRouter()
	r.HandleFunc("/register", userSvc.RegisterHandler)
	r.HandleFunc("/user", userSvc.GetUserHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	_ = srv.ListenAndServe()
}

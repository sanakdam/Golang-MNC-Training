package main

import (
	"Golang-MNC-Training/Kode-Assignment-2/server"
	"Golang-MNC-Training/Kode-Assignment-2/server/controllers"
	"Golang-MNC-Training/Kode-Assignment-2/server/db"
)

func main() {
	db := db.ConnectGorm()
	peopleController := controllers.NewOrderController(db)
	router := server.NewRouter(peopleController)
	router.Start(":4000")
}

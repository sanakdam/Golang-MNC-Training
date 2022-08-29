package main

import (
	"fmt"

	"github.com/sanakdam/Golang-MNC-Training/Assignment_5/service"
)

func main() {
	userSvc := service.NewUserService()
	res := userSvc.Register(&service.User{Nama: "budi"})
	fmt.Println(res)
}

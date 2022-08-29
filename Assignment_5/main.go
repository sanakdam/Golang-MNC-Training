package main

import (
	"fmt"

	"Assignment_5/service"
)

func main() {
	userSvc := service.NewUserService()
	res := userSvc.Register(&service.User{Nama: "budi"})
	fmt.Println(res)
	res2 := userSvc.Register(&service.User{Nama: "ana"})
	fmt.Println(res2)
	users := userSvc.GetUser()
	fmt.Println("---Hasil get user---")
	for _, user := range users {
		fmt.Println(user.Nama)
	}
}

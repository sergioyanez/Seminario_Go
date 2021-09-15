package main

import (
	"fmt"
	"tudai2021/services"
)

func main() {
	var userSrv services.UserService    // Defino la variable userSrv como la interfaz UserService
	var userSrvAWS services.UserService // Defino la variable userSrvAWS como la interfaz UserService

	userSrv, err := services.NewUserServiceLocal()   //Llamo a la función constructora que me devuelve la implementación de la interfaz
	userSrvAWS, err1 := services.NewUserServiceAWS() //Llamo a la función constructora que me devuelve la implementación de la interfaz

	if err1 != nil {
		panic(err)
	}
	uAWS, err1 := userSrvAWS.Login("jppAWS", "654321")
	u, err := userSrv.Login("jpp", "123456")
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
	fmt.Println(uAWS)
}

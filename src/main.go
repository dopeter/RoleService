package main


import (
	"fmt"
	"domain"
)

func main(){
	fmt.Println("main func start")

	//todo tmp test put here.
	role,_:=domain.BuildRole("Role1","Rolr1")

	fmt.Println(role.RoleName)
}
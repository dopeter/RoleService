package main


import (
	"fmt"
	"domain/role"
)

func main(){
	fmt.Println("main func start")

	//todo tmp test put here.
	role,_:=role.BuildRole("Role1","Rolr1")

	fmt.Println(role.RoleName)
}
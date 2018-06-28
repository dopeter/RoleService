package main

import (
	"encoding/json"
	"fmt"
)

type roleRequest struct{
	RoleName string
	Desc string
	Permissions []*permissionRequest
}

type permissionRequest struct{
	Name string
	Operation string
	Path string
	Desc string
	RoleId int
}


func main(){
	body:=`{
		"RoleName":"test1",
		"Permissions":	[
			{"Name":"perm1"},
			{"Name":"perm2"}
		]
}`

	var role roleRequest

	err:=json.Unmarshal([]byte(body),&role)
	if err!=nil{
		fmt.Println(err.Error())
	}


}

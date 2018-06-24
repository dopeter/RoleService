package role

import (
	"toolkit"
	"errors"
)



var(
	ADMIN_ROLE_NAME="Admin"
	VALIDATION_ROLE_LENGTH=15
	VALIDATION_ROLE_LENGTH_ERROR="The length of role name must be less than"+toolkit.ToString(VALIDATION_ROLE_LENGTH)
	VALIDATION_PERMISSION_LENGTH_ERROR="The length of permissions must be greater than 0."
)

type Role struct{
	RoleId int
	RoleName string
	Desc string
	Permissions []Permission
}


func BuildAdmin()*Role{
	role:=&Role{
		RoleName:ADMIN_ROLE_NAME,
		Desc:"Admin",
		Permissions:nil,
	}

	return role
}

func BuildRole(roleName string,desc string) (*Role,error){

	if len(roleName)>VALIDATION_ROLE_LENGTH{
		return nil,errors.New(VALIDATION_ROLE_LENGTH_ERROR)
	}

	role:=&Role{
		RoleName:roleName,
		Desc:desc,
		Permissions:nil,
	}

	return role,nil
}

func(role *Role) RefreshPermission(permissions []Permission) error{
	if permissions==nil || len(permissions)==0{
		return errors.New(VALIDATION_PERMISSION_LENGTH_ERROR)
	}

	role.Permissions=permissions

	return nil
}
package session

import (
	"errors"
	"github.com/satori/go.uuid"
	"toolkit"
)


type UserSession struct{
	Id string
	UserId string
	RoleId	 int
	Identity string
}

var(
	ADMIN_ROLE_NAME="Admin"
	VALIDATION_USERID_EMPTY_ERROR="Please specify user id."
	ERROR_UUID_GENERATE="Generate uuid error."
)


func BuildSession(userId string,roleId int) (*UserSession,error){

	if(userId==""){
		return nil,errors.New(VALIDATION_USERID_EMPTY_ERROR)
	}

	uid,err:=uuid.NewV4()

	if(err!=nil){
		//todo log record error
		return nil,errors.New(ERROR_UUID_GENERATE)
	}

	return &UserSession{
		Id:BuildSessionId(userId,roleId),
		UserId:userId,
		RoleId:roleId,
		Identity:uid.String(),
	},nil

}

func BuildSessionId(userId string,roleId int)string{
	return userId+toolkit.ToString(roleId)
}
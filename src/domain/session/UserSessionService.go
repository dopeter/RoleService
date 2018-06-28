package session

import "errors"

var(
	ERROR_BUSINESS_SESSION_EXIST="User has logged once."
)

type UserSessionService struct {
	repository UserSessionRepository
}

func BuildUserSessionService(repositoryImpl UserSessionRepository) *UserSessionService{
	return &UserSessionService{repository:repositoryImpl}
}


func (service *UserSessionService) BuildUserSeesion(userId string,roleId int) error{

	//ensure is singleton login.
	existSession:=service.repository.RetrieveSessionByUserAndRole(userId,roleId)

	if(existSession!=nil){
		errors.New(ERROR_BUSINESS_SESSION_EXIST)
	}

	//todo check role is exist?

	newSession,err:=BuildSession(userId,roleId)
	if err!=nil{
		return err
	}

	service.repository.SaveSession(newSession)

	return nil
}

func (service *UserSessionService) RemoveUserSession(userId string,roleId int) {
	service.repository.RemoveSession(userId,roleId)
}

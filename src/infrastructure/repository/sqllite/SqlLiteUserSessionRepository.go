package sqllite

import (
	. "domain/session"
	"github.com/jinzhu/gorm"
)

type SqlLiteUserSessionRepository struct{
	db *gorm.DB

}

func BuildSqlLiteUserSessionRepository() *SqlLiteUserSessionRepository{

	//build table
	GetDbEngine().Engine.CreateTable(&UserSession{})

	return &SqlLiteUserSessionRepository{
		db:GetDbEngine().Engine,
	}
}

func(repository *SqlLiteUserSessionRepository) retrieveSession(userId string,roleId int) *UserSession{

	session:=&UserSession{UserId:"-1"}

	repository.db.Where(&UserSession{UserId:userId,RoleId:roleId}).First(session)

	if session.UserId=="-1"{
		return nil
	}
	return session
}

func(repository *SqlLiteUserSessionRepository) RetrieveSessionByUserAndRole(userId string,roleId int) *UserSession{
	return repository.retrieveSession(userId,roleId)
}

func(repository *SqlLiteUserSessionRepository) SaveSession(session *UserSession){
	//check session is exist
	if repository.retrieveSession(session.UserId,session.RoleId) ==nil {
		repository.db.Create(session)
	}else{
		repository.db.Save(session)
	}
}

func(repository *SqlLiteUserSessionRepository) RemoveSession(userId string,roleId int){
	if session:=repository.retrieveSession(userId,roleId); session!=nil{
		repository.db.Delete(session)
	}
}
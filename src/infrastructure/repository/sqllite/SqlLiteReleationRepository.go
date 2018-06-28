package sqllite

import (
	"github.com/jinzhu/gorm"
	"domain/userrole"
)

type SqlLiteReleationRepository struct {
	db *gorm.DB
}

func BuildSqlLiteReleationRepository()*SqlLiteReleationRepository{

	GetDbEngine().Engine.CreateTable(&userrole.Releation{})

	return &SqlLiteReleationRepository{
		db:GetDbEngine().Engine,
	}
}

func (repository *SqlLiteReleationRepository)AddReleation (userId string,roleIds []int){

	releation:=&userrole.Releation{UserId:userId}
	releation.RoleIds=releation.JoinRoleId(roleIds)
	repository.db.Create(releation)
}

func (repository *SqlLiteReleationRepository)RemoveReleation(userId string){
	repository.db.Where(&userrole.Releation{UserId:userId}).Delete(&userrole.Releation{})
}

func (repository *SqlLiteReleationRepository)RetrieveReleation(userId string)*userrole.Releation{
	var releation =&userrole.Releation{UserId:"-1"}

	repository.db.Where(&userrole.Releation{UserId:userId}).First(releation)

	if releation.UserId=="-1"{
		return nil
	}

	return releation
}

func (repository *SqlLiteReleationRepository) UpdateReleation(releation *userrole.Releation){
	repository.db.Save(releation)
}
package sqllite

import (
	"github.com/jinzhu/gorm"
	"domain/role"
)


type SqlLiteRoleRepository struct {
	db *gorm.DB
}

func BuildSqlLiteRoleRepository()*SqlLiteRoleRepository{

	GetDbEngine().Engine.CreateTable(&role.Role{})
	GetDbEngine().Engine.CreateTable(&role.Permission{})

	return &SqlLiteRoleRepository{
		db:GetDbEngine().Engine,
	}
}

func(repository *SqlLiteRoleRepository) SaveRole(role *role.Role){
	repository.db.Save(role)
}

func(repository *SqlLiteRoleRepository) RemoveRole(role *role.Role){
	repository.db.Delete(role)
}

func(repository *SqlLiteRoleRepository) RetrieveRole(roleId int) *role.Role{
	retrieveRole:=&role.Role{RoleName:"-1"}

	repository.db.Preload("Permissions").First(retrieveRole,roleId)

	if(retrieveRole.RoleName=="-1"){
		return nil
	}

	return retrieveRole
}

func(repository *SqlLiteRoleRepository) RetrieveRoleByIds(roleIds []int) []*role.Role{
	roles:=[]*role.Role{}
	repository.db.Where(roleIds).Preload("Permissions").Find(&roles)

	return roles
}

func(repository *SqlLiteRoleRepository) AddRole(role *role.Role){
/*
	t := time.NewTimer(10 * time.Second)
	expire := <- t.C
	log.Println("Expiration time: %v.\n", expire)
*/
	repository.db.Create(role)
}



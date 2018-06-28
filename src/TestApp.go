package main

import (
	"infrastructure/repository/sqllite"
	"domain/userrole"
	"domain/session"
	"domain/role"
	"fmt"
)

func main(){

	engine:=sqllite.GetDbEngine().Engine

	engine.CreateTable(&userrole.Releation{})
	engine.CreateTable(&session.UserSession{})
	engine.CreateTable(&role.Role{})
	engine.CreateTable(&role.Permission{})

	userId:="user1"
	userId2:="user2"
	//test releation
	releationRepo:=sqllite.BuildSqlLiteReleationRepository()

	releationRepo.AddReleation(userId,[]int{1,2,3,4,5})
	releation:=releationRepo.RetrieveReleation(userId)
	fmt.Println(releation.RoleIds)

	releation=releationRepo.RetrieveReleation(userId)
	releation.UserId=userId2
	engine.Save(releation)
	//updateReleation:=&userrole.Releation{}
	//updateReleation.ID=1
	//updateReleation.UserId="user2"
	//releationRepo.UpdateReleation(updateReleation)

	tReleation:=userrole.Releation{}
	engine.First(&tReleation)
	tReleation.UserId="user3"
	engine.Update(tReleation)

	releationRepo.RemoveReleation(userId2)
	//test session
	sessionRepo:=sqllite.BuildSqlLiteUserSessionRepository()
	sessionRepo.SaveSession(&session.UserSession{UserId:userId,RoleId:1})

	sessionRepo.RetrieveSessionByUserAndRole(userId,1)
	sessionRepo.RemoveSession(userId,1)

	//test role
	roleRepo:=sqllite.BuildSqlLiteRoleRepository()

	newRole:=&role.Role{
		RoleName:"role1",
		Desc:"desc1",
		Permissions:[]*role.Permission{
			&role.Permission{
				Name:"permissino1",
			},
			&role.Permission{
				Name:"permissino2",
			},
		},
	}


	roleRepo.AddRole(newRole)
	retrieveRole:=roleRepo.RetrieveRole(1)
	retrieveRoles:=roleRepo.RetrieveRoleByIds([]int{1})

	//retrieveRole.Permissions[0].Name="per2"
	newRole.RoleName="roleNameNew2"
	roleRepo.SaveRole(newRole)

	retrieveRole.RoleName="roleNameNew3"
	roleRepo.SaveRole(retrieveRole)
/*
	retrieveRole.Permissions[0].Name="PerNewName1"

	for _,per:=range retrieveRole.Permissions{
		now:=time.Now()
		per.DeletedAt=&now
	}
*/

	engine.Model(retrieveRole).Association("Permissions").
		Replace([]*role.Permission{
		&role.Permission{
			Name:"newpem1",
		},
		&role.Permission{
			Name:"newpem2",
		},
	})

	engine.Model(retrieveRole).Association("Permissions").
		Append(
			[]*role.Permission{
				&role.Permission{
					Name:"newpem3",
				},
				&role.Permission{
					Name:"newpem4",
				},
			})

	roleRepo.SaveRole(retrieveRole)

	retrieveRole=roleRepo.RetrieveRole(1)
	retrieveRoles=roleRepo.RetrieveRoleByIds([]int{1})

	fmt.Println(retrieveRole.RoleName)
	fmt.Println(len(retrieveRoles))
}

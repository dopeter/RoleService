package application

import(
	"domain/role"
	"domain/userrole"
	"domain/session"

	"infrastructure/repository/sqllite"
)


type RoleAppService interface {
	AddRole(role *role.Role)
	EditRole(role *role.Role)
	RemoveRole(role *role.Role)

	BuildUserSession(userId string,roleId int)

	AddReleation (userId string,roleIds []int)

	RemoveReleation(userId string)

	RetrieveReleation(userId string) *userrole.Releation

	UpdateReleation(releation *userrole.Releation)
}

type roleAppService struct {
	roleRepo role.RoleRepository
	sessionRepo session.UserSessionRepository
	releationRepo userrole.ReleationRespisotry
	sessionService *session.UserSessionService
}
//todo replace sqllite to mysql
func BuildRoleAppService() RoleAppService{
	var roleApp RoleAppService
	roleApp=&roleAppService{
		roleRepo:sqllite.BuildSqlLiteRoleRepository(),
		sessionRepo:sqllite.BuildSqlLiteUserSessionRepository(),
		releationRepo:sqllite.BuildSqlLiteReleationRepository(),
		sessionService:session.BuildUserSessionService(sqllite.BuildSqlLiteUserSessionRepository()),
	}

	return roleApp
}

func(service *roleAppService)AddRole(role *role.Role){
	service.roleRepo.AddRole(role)
}

func(service *roleAppService)EditRole(role *role.Role){

}
func(service *roleAppService)RemoveRole(role *role.Role){

}

func(service *roleAppService) BuildUserSession(userId string,roleId int){

}

func(service *roleAppService) AddReleation (userId string,roleIds []int){

}

func(service *roleAppService) RemoveReleation(userId string){

}

func(service *roleAppService) RetrieveReleation(userId string) *userrole.Releation{
	return nil
}

func(service *roleAppService) UpdateReleation(releation *userrole.Releation){

}

package role

type RoleService struct {
	reposiory RoleRepository
}

func BuildRoleService(roleSeviceImpl RoleRepository) *RoleService{
	return &RoleService{roleSeviceImpl}
}













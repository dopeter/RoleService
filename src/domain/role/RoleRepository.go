package role

type RoleRepository interface {
	SaveRole(role *Role)

	RemoveRole(role *Role)

	RetrieveRole(roleId int) *Role

	RetrieveRoleByIds(roleIds []int) []*Role

	AddRole(role *Role)
}

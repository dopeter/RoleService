package role

type RoleRepository interface {
	SaveRole(role *Role)

	EditRole(role *Role)

	RemoveRole(role *Role)
}

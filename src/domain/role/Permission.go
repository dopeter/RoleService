package role


type Permission struct{
	Id int
	Name string
	ObjectOperationMap map[PermissionObject]Operation
}
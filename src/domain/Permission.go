package domain


type Permission struct{
	Id int
	Name string
	ObjectOperationMap map[PermissionObject]Operation
}
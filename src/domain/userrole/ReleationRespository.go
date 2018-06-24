package userrole


type ReleationRespisotry interface {
	AddReleation (userId string,roleId int)

	RemoveReleation(userId string,roleId int)

	RetrieveReleation(userId string)
}




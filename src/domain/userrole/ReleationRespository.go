package userrole


type ReleationRespisotry interface {
	AddReleation (userId string,roleIds []int)

	RemoveReleation(userId string)

	RetrieveReleation(userId string) *Releation

	UpdateReleation(releation *Releation)
}




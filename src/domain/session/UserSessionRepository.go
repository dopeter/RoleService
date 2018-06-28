package session

type UserSessionRepository interface {
	SaveSession(session *UserSession)

	RetrieveSessionByUserAndRole(userId string,roleId int) *UserSession

	RemoveSession(userId string,roleId int)
}


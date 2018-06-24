package session

type UserSessionRepository interface {
	SaveSession(session *UserSession)

	RetrieveSessionByUserAndRole(userId string,roleId int) *UserSession

	RemoveSession(userId string,roleId int)
}

type UserSessionInMemoryRepository struct{
	SessionCache map[string]*UserSession
}

func BuildUserSessionInMemoryRepository() *UserSessionInMemoryRepository{
		return &UserSessionInMemoryRepository{
		SessionCache:make(map[string]*UserSession),//build empty cache
	}
}

func(repository *UserSessionInMemoryRepository) retrieveSessionFromCache(userId string,roleId int) *UserSession{
	if session,exist:=repository.SessionCache[BuildSessionId(userId,roleId)] ; exist{
		return session
	}

	return nil
}

func(repository *UserSessionInMemoryRepository) RetrieveSessionByUserAndRole(userId string,roleId int) *UserSession{
	return repository.retrieveSessionFromCache(userId,roleId)
}

func(repository *UserSessionInMemoryRepository) SaveSession(session *UserSession){
	//check session is exist
	if repository.retrieveSessionFromCache(session.UserId,session.RoleId) ==nil{
		repository.SessionCache[BuildSessionId(session.UserId,session.RoleId)]=session
	}
}

func(repository *UserSessionInMemoryRepository) RemoveSession(userId string,roleId int){
	if repository.retrieveSessionFromCache(userId,roleId) !=nil{
		delete(repository.SessionCache,BuildSessionId(userId,roleId))
	}
}
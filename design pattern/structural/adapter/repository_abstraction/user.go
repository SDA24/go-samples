package repo_abstraction

type User struct {
	// some fields
}
type Iuser_repository interface {
	GetByID(id uint) (*User, error)
}

type User_repository struct {
	I_user_repository Iuser_repository
}

func New_user_repository(i Iuser_repository) User_repository {
	return User_repository{
		I_user_repository: i,
	}
}

func (u User_repository) GetByID(id uint) (*User, error) {
	return u.I_user_repository.GetByID(id)
}

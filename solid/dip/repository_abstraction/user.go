package repo_abstraction

type User struct {
	// some fields
}

type UserRepository interface {
	GetByID(id uint) (*User, error)
}

package user

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepo interface {
	GetAuthenticated(usr User) (*User, bool, int, error)
	GetUserID(username string) (*string, error)
}

package domain

type User struct {
	ID       int
	UserName string
	Password string
	Role     string
	Balance  int
}

type IUserService interface {
	Register(username, password, role string) error
	Login(username, password string) (*User, error)
	ChangeBalance(user *User, oldBalance, newBalance int) (*User, error)
	DeleteUser(UserId int) error
	ListUser() ([]User, error)
}

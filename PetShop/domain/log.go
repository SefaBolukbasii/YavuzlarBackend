package domain

type Log struct {
	ID          int
	UserId      int
	Transaction string
}

type ILog interface {
	LogAdd(UserId int, Transaction string) error
}

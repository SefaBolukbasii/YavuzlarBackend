package domain

type Item struct {
	ID    int
	Name  string
	Price int
}

type Iitem interface {
	AddItem(item Item) error
	DeleteItem(id int) error
	UpdateItem(item Item) error
	ListItems() ([]Item, error)
	ToBuy(Item *Item, User *User) error
	ListBuy(User *User) ([]Item, error)
}

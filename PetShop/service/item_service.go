package service

import (
	"petshop/domain"

	jsondb "github.com/SefaBolukbasii/JsonDB"
)

type ItemService struct {
	db *jsondb.Database
}

func CreateItemService(db *jsondb.Database) domain.Iitem {
	return &ItemService{db: db}
}
func (Is *ItemService) AddItem(item domain.Item) error {

	if err := Is.db.Insert("Items", map[string]any{
		"ItemName": item.Name,
		"Price":    item.Price,
	}); err != nil {
		return err
	}
	return nil
}

func (Is *ItemService) DeleteItem(id int) error {

	if err := Is.db.Delete("Items", "id", id); err != nil {
		return err
	}
	return nil
}

func (Is *ItemService) UpdateItem(item *domain.Item, newPrice int) error {
	oldPrice := item.Price

	if err := Is.db.Update("Items", "Price", oldPrice, newPrice); err != nil {
		return err
	}
	item.Price = newPrice
	return nil
}

func (Is *ItemService) ListItems() ([]domain.Item, error) {

	Items, err := Is.db.Select("Items")
	if err != nil {
		return nil, err
	}
	var ItemArray []domain.Item
	for _, item := range Items {
		a := domain.Item{
			ID:    item["id"].(int),
			Name:  item["ItemName"].(string),
			Price: item["Price"].(int),
		}
		ItemArray = append(ItemArray, a)
	}
	return ItemArray, nil
}

func (Is *ItemService) ToBuy(Item *domain.Item, User *domain.User) error {

	if err := Is.db.Insert("ItemUser", map[string]any{
		"ItemId":     Item.ID,
		"UserId":     User.ID,
		"TotalPrice": Item.Price,
	}); err != nil {
		return err
	}
	userService := UserService{}
	err := userService.ChangeBalance(User, User.Balance, User.Balance-Item.Price)
	if err != nil {
		return err
	}
	return nil

}
func (Is *ItemService) ListBuy(User *domain.User) ([]domain.Item, error) {
	var ItemsId []int
	var BuyItemsArray []domain.Item
	var ItemsArray []domain.Item

	BuyItems, err := Is.db.Select("ItemUser")
	if err != nil {
		return nil, err
	}
	for _, BuyItem := range BuyItems {
		if BuyItem["UserId"] == User.ID {
			ItemsId = append(ItemsId, BuyItem["ItemId"].(int))
		}
	}
	ItemsArray, err = Is.ListItems()
	if err != nil {
		return nil, err
	}
	for _, Item_ := range ItemsArray {
		for _, a := range ItemsId {
			if Item_.ID == a {
				BuyItemsArray = append(BuyItemsArray, Item_)
			}
		}
	}
	return BuyItemsArray, err

}

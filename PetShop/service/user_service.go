package service

import (
	"petshop/Database"
	"petshop/domain"
	"petshop/pkg"
)

type UserService struct{}

func (us *UserService) Register(username, password, role string) error {
	// Kullanıcı kayıt işlemi
	Veritabani, err := Database.Connect()
	if err != nil {
		return err
	}
	password, err = pkg.HashPassword(password)
	if err != nil {
		return err
	}
	if err := Veritabani.Db.Insert("Users", map[string]any{
		"username": username,
		"password": password,
		"role":     role,
		"balance":  0,
	}); err != nil {
		return err
	}

	return nil
}

func (us *UserService) Login(username, password string) (*domain.User, error) {
	Veritabani, err := Database.Connect()
	if err != nil {
		return nil, err
	}
	users, err := Veritabani.Db.Select("Users")
	if err != nil {
		return nil, err
	}
	passwordHash, err := pkg.HashPassword(password)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user["username"] == username && user["password"] == passwordHash {
			return &domain.User{
				ID:       user["id"].(int),
				UserName: user["username"].(string),
				Password: user["password"].(string),
				Role:     user["role"].(string),
				Balance:  user["balance"].(int),
			}, nil
		}
	}
	return nil, nil
}

func (us *UserService) ChangeBalance(user *domain.User, oldBalance, newBalance int) error {
	Veritabani, err := Database.Connect()
	if err != nil {
		return err
	}
	if err := Veritabani.Db.Update("Users", "balance", oldBalance, newBalance); err != nil {
		return err
	}
	user.Balance = newBalance
	return nil
}
func (us *UserService) DeleteUser(UserId int) error {
	Veritabani, err := Database.Connect()
	if err != nil {
		return err
	}
	if err := Veritabani.Db.Delete("Users", "id", UserId); err != nil {
		return err
	}
	return nil
}
func (us *UserService) ListUser() ([]domain.User, error) {
	Veritabani, err := Database.Connect()
	if err != nil {
		return nil, err
	}
	Users, err := Veritabani.Db.Select("Users")
	if err != nil {
		return nil, err
	}
	var UsersArray []domain.User
	for _, user := range Users {
		UsersArray = append(UsersArray, domain.User{
			ID:       user["id"].(int),
			UserName: user["username"].(string),
			Role:     user["role"].(string),
			Balance:  user["balance"].(int),
		})
	}
	return UsersArray, nil
}

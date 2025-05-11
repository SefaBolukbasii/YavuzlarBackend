package services

import (
	"spotifyAPI/models"
	"spotifyAPI/repositories"
)

func Login(username string, password string) (models.User, error) {
	User, err := repositories.Login(username, password)
	if err != nil {
		return models.User{}, err
	}
	return User, nil
}
func Register(user models.User) error {
	err := repositories.Register(user)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(id string) (models.User, error) {
	User, err := repositories.GetUser(id)
	if err != nil {
		return models.User{}, err
	}
	return User, nil
}
func UpdateUser(id string, username string, password string) error {
	err := repositories.UpdateUser(id, username, password)
	if err != nil {
		return err
	}
	return nil
}
func DeleteUser(id string) error {
	err := repositories.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
func GetAllUsers() ([]models.User, error) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
func GetPremium(id string) error {
	err := repositories.GetPremium(id)
	if err != nil {
		return err
	}
	return nil
}

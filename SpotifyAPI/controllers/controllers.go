package controllers

import (
	"spotifyAPI/models"
	"spotifyAPI/services"
	"spotifyAPI/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store

func Login(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	LoginUser, err := services.Login(user.Username, user.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	store = session.New()
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create session")
	}
	sess.Set("userId", LoginUser.Id)
	sess.Set("username", LoginUser.Username)
	sess.Set("accountType", LoginUser.AccountType)
	sess.Set("cash", LoginUser.Cash)
	sess.Set("roles", LoginUser.Roles)
	sess.Save()
	return c.SendString("Login successful")
}
func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err := services.Register(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendString("User registered successfully")
}
func GetUser(c *fiber.Ctx) error {

	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	userId := sess.Get("userId")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not logged in")
	}
	UserInfo, err := services.GetUser(userId.(string))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	utils.WriteJson(c, UserInfo)
	return nil
}
func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		return err
	}
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	userId := sess.Get("userId")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not logged in")
	}
	err = services.UpdateUser(userId.(string), user.Username, user.Password)
	if err != nil {
		fiber.NewError(fiber.StatusNotFound, "User not found")
		return err
	}
	utils.WriteJson(c, "Başarı ile güncellendi")
	return nil
}
func DeleteUser(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	userId := sess.Get("userId")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not logged in")
	}
	err = services.DeleteUser(userId.(string))
	if err != nil {
		fiber.NewError(fiber.StatusNotFound, "User not found")
		return err
	}
	utils.WriteJson(c, "Başarı ile silindi")
	return nil
}
func GetAllUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get users")
	}
	utils.WriteJson(c, users)
	return nil
}
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := services.GetUser(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}
	utils.WriteJson(c, user)
	return nil
}
func UpdateUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	err := services.UpdateUser(id, user.Username, user.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}
	utils.WriteJson(c, "Başarı ile güncellendi")
	return nil
}
func DeleteUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	err := services.DeleteUser(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}
	utils.WriteJson(c, "Başarı ile silindi")
	return nil
}
func GetPremium(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	userId := sess.Get("userId")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not logged in")
	}
	accountType := sess.Get("accountType")
	cash := sess.Get("cash")
	if accountType == "free" && cash.(int) >= 50 {
		err = services.GetPremium(userId.(string))
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		utils.WriteJson(c, "Başarı ile premium alındı")
		return nil
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "You don't have enough cash or you are already premium")
	}
}

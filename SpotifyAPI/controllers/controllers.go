package controllers

import (
	"spotifyAPI/models"
	"spotifyAPI/services"
	"spotifyAPI/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
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
	redisStore := redis.New(redis.Config{
		Host:     "redis",
		Port:     6379,
		Password: "",
		Database: 0,
	})
	store = session.New(session.Config{
		Storage: redisStore,
	})
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
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	role := sess.Get("roles")
	if role == "admin" {
		users, err := services.GetAllUsers()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get users")
		}
		utils.WriteJson(c, users)
		return nil
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource")
	}
}
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	role := sess.Get("roles")
	if role == "admin" {
		user, err := services.GetUser(id)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		utils.WriteJson(c, user)
		return nil
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource")
	}
}
func UpdateUserById(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	role := sess.Get("roles")
	if role == "admin" {
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
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource")
	}
}
func DeleteUserById(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	role := sess.Get("roles")
	if role == "admin" {
		id := c.Params("id")
		err := services.DeleteUser(id)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		utils.WriteJson(c, "Başarı ile silindi")
		return nil
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource")
	}
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
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	cuponId, ok := data["cuponId"].(string)
	if !ok {
		if accountType == "free" && cash.(int) >= 50 {
			err = services.GetPremium(userId.(string), "nul")
			if err != nil {
				return fiber.NewError(fiber.StatusNotFound, "User not found")
			}
			utils.WriteJson(c, "Başarı ile premium alındı")

			now := time.Now()
			expire := now.Add(30 * 24 * time.Hour)
			sess.Set("accountType", "premium")
			sess.Set("premium_start", now.Format(time.RFC3339))
			sess.Set("premium_expire", expire.Format(time.RFC3339))
			sess.Save()
			return nil
		} else {
			return fiber.NewError(fiber.StatusBadRequest, "You don't have enough cash or you are already premium")
		}
	}
	if accountType == "free" && cash.(int) >= 50 {
		err = services.GetPremium(userId.(string), cuponId)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		utils.WriteJson(c, "Başarı ile premium alındı")

		now := time.Now()
		expire := now.Add(30 * 24 * time.Hour)
		sess.Set("accountType", "premium")
		sess.Set("premium_start", now.Format(time.RFC3339))
		sess.Set("premium_expire", expire.Format(time.RFC3339))
		sess.Save()
		return nil
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "You don't have enough cash or you are already premium")
	}

}
func AddSong(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	role := sess.Get("roles")
	if role == "admin" {
		var song models.Song
		if err := c.BodyParser(&song); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}
		err := services.AddSong(song)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to add song")
		}
		utils.WriteJson(c, "Başarı ile şarkı eklendi")
		return nil
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource")
	}
}
func DeleteSong(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	role := sess.Get("roles")
	if role == "admin" {
		id := c.Params("id")
		err := services.DeleteSong(id)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "Song not found")
		}
		utils.WriteJson(c, "Başarı ile şarkı silindi")
		return nil
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource")
	}
}
func UpdateSong(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	role := sess.Get("roles")
	if role == "admin" {
		id := c.Params("id")
		var song models.Song
		if err := c.BodyParser(&song); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}
		err = services.UpdateSong(song, id)
		if err != nil {
			return err
		}
		utils.WriteJson(c, "Başarı ile şarkı güncellendi")
		return nil
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource")
	}
}
func GetAllSongs(c *fiber.Ctx) error {
	pageParam := c.Query("page", "1") // default: "1"
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 0 {
		page = 0
	}
	songs, err := services.GetAllSongs(page)
	if err != nil {
		return err
	}
	utils.WriteJson(c, songs)
	return nil
}
func GetSong(c *fiber.Ctx) error {
	id := c.Params("id")
	song, err := services.GetSong(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Song not found")
	}
	utils.WriteJson(c, song)
	return nil
}
func GetMyPlaylists(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	id := sess.Get("userId")
	playlists, err := services.GetMyPlaylists(id.(string))
	if err != nil {
		return err
	}
	utils.WriteJson(c, playlists)
	return nil
}
func GetMyPlaylistID(c *fiber.Ctx) error {
	id := c.Params("playlistID")
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	userId := sess.Get("userId")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not logged in")
	}
	playlist, err := services.GetMyPlaylistID(id, userId.(string))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Playlist not found")
	}
	utils.WriteJson(c, playlist)
	return nil
}
func AddPlaylist(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	userId := sess.Get("userId")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not logged in")
	}
	var playlist models.Playlist
	if err := c.BodyParser(&playlist); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	err = services.AddPlaylist(playlist, userId.(string))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to add playlist")
	}
	utils.WriteJson(c, "Başarı ile playlist eklendi")
	return nil
}
func DeletePlaylist(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	userId := sess.Get("userId")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not logged in")
	}
	id := c.Params("playlistID")
	err = services.DeletePlaylist(id, userId.(string))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Playlist not found")
	}
	utils.WriteJson(c, "Başarı ile playlist silindi")
	return nil
}
func PlaylistAddSong(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	userId := sess.Get("userId")
	account_type := sess.Get("accountType")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not logged in")
	}
	playlistID := c.Params("playlistID")
	songID := c.Params("songID")
	expired := sess.Get("premium_expire")
	expireStr, ok := expired.(string)
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError, "Session veri tipi hatalı")
	}

	expireTime, err := time.Parse(time.RFC3339, expireStr)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Tarih biçimi hatalı")
	}

	// Şu anki zamanla karşılaştır
	if time.Now().After(expireTime) {
		err = services.PlaylistAddSong(playlistID, songID, userId.(string), "free")
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
	}
	err = services.PlaylistAddSong(playlistID, songID, userId.(string), account_type.(string))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	utils.WriteJson(c, "Başarı ile şarkı eklendi")
	return nil
}
func GetPlaylistUser(c *fiber.Ctx) error {
	userID := c.Params("userID")
	playlists, err := services.GetPlaylistUser(userID)
	if err != nil {
		return err
	}
	utils.WriteJson(c, playlists)
	return nil
}
func GetPlaylistSongUser(c *fiber.Ctx) error {
	songID := c.Params("songID")
	song, err := services.GetPlaylistSongUser(songID)
	if err != nil {
		return err
	}
	utils.WriteJson(c, song)
	return nil
}
func CreateCupon(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	role := sess.Get("roles")
	if role == "admin" {
		var cupon models.Cupon
		if err := c.BodyParser(&cupon); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}
		err := services.CreateCupon(cupon)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		utils.WriteJson(c, "Başarı ile kupon oluşturuldu")
		return nil
	}
	return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource")
}
func AssignCupon(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	role := sess.Get("roles")
	if role == "admin" {
		var data map[string]interface{}
		if err := c.BodyParser(&data); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
		}
		cuponId, ok := data["cuponId"].(string)
		if !ok {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid cuponId")
		}
		userId, ok := data["userId"].(string)
		if !ok {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid userId")
		}
		err = services.AssignCupon(cuponId, userId)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to assign cupon")
		}
		utils.WriteJson(c, "Başarı ile kupon atandı")
		return nil
	}
	return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource")
}
func GetCupon(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	userId := sess.Get("userId")
	if userId == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "User not logged in")
	}
	cupons, err := services.GetCupon(userId.(string))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Cupon not found")
	}
	utils.WriteJson(c, cupons)
	return nil
}
func ListCupons(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to get session")
	}
	role := sess.Get("roles")
	if role == "admin" {
		cupons, err := services.ListCupons()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		utils.WriteJson(c, cupons)
		return nil
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "You are not authorized to access this resource")
	}
}

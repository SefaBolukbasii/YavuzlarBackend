package routes

import (
	"spotifyAPI/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes() {
	app := fiber.New()
	app.Post("/api/v1/login", controllers.Login)
	app.Post("/api/v1/register", controllers.Register)
	app.Get("/api/v1/user", controllers.GetUser)
	app.Put("/api/v1/user", controllers.UpdateUser)
	app.Delete("/api/v1/user", controllers.DeleteUser)
	app.Get("/api/v1/admin/user", controllers.GetAllUsers)
	app.Get("/api/v1/admin/user/:id", controllers.GetUserById)
	app.Put("/api/v1/admin/user/:id", controllers.UpdateUserById)
	app.Delete("/api/v1/admin/user/:id", controllers.DeleteUserById)
	//app.Get("/api/v1/getPremium", controllers.GetPremium)

	app.Post("/api/v1/admin/song/", controllers.AddSong)
	app.Delete("/api/v1/admin/song/:id", controllers.DeleteSong)
	app.Put("/api/v1/admin/song/:id", controllers.UpdateSong)

	app.Get("/api/v1/user/song/", controllers.GetAllSongs)
	app.Get("/api/v1/user/song/:id", controllers.GetSong)

	app.Get("/api/v1/user/playlist", controllers.GetMyPlaylists)
	app.Get("/api/v1/user/playlist/:playlistID", controllers.GetMyPlaylistID)
	app.Post("/api/v1/user/playlist", controllers.AddPlaylist)
	app.Delete("/api/v1/user/playlist/:playlistID", controllers.DeletePlaylist)
	app.Post("/api/v1/user/playlist/:playlistID/:songID", controllers.PlaylistAddSong)
	app.Get("/api/v1/user/playlist/:userID", controllers.GetPlaylistUser)
	app.Get("/api/v1/user/playlist/:userID/:songID", controllers.GetPlaylistSongUser)

	app.Post("/api/v1/admin/cupon", controllers.CreateCupon)
	app.Post("/api/v1/admin/cupon/assign", controllers.AssignCupon)
	app.Get("/api/v1/admin/cupon", controllers.ListCupons)
	app.Get("/api/v1/user/cupon", controllers.GetCupon)
	app.Post("/api/v1/user/premium", controllers.GetPremium)

	app.Listen(":8080")
}

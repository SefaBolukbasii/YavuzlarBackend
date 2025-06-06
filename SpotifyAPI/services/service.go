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
func GetPremium(id string, cuponId string) error {
	err := repositories.GetPremium(id, cuponId)
	if err != nil {
		return err
	}
	return nil
}
func AddSong(song models.Song) error {
	err := repositories.AddSong(song)
	if err != nil {
		return err
	}
	return nil
}
func DeleteSong(id string) error {
	err := repositories.DeleteSong(id)
	if err != nil {
		return err
	}
	return nil
}
func UpdateSong(song models.Song, id string) error {
	err := repositories.UpdateSong(song, id)
	if err != nil {
		return err
	}
	return nil
}
func GetAllSongs(page int) ([]models.Song, error) {
	songs, err := repositories.GetAllSongs(page)
	if err != nil {
		return nil, err
	}
	return songs, nil
}
func GetSong(id string) (models.Song, error) {
	song, err := repositories.GetSong(id)
	if err != nil {
		return models.Song{}, err
	}
	return song, nil
}
func GetMyPlaylists(id string) ([]models.Playlist, error) {
	playlist, err := repositories.GetMyPlaylists(id)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}
func GetMyPlaylistID(id string, userId string) (models.Playlist, error) {
	playlist, err := repositories.GetMyPlaylistID(id, userId)
	if err != nil {
		return models.Playlist{}, err
	}
	return playlist, nil
}
func AddPlaylist(playlist models.Playlist, userId string) error {
	err := repositories.AddPlaylist(playlist, userId)
	if err != nil {
		return err
	}
	return nil
}
func DeletePlaylist(id string, userId string) error {
	err := repositories.DeletePlaylist(id, userId)
	if err != nil {
		return err
	}
	return nil
}
func PlaylistAddSong(playlistID string, songID string, userId string, account_type string) error {
	err := repositories.PlaylistAddSong(playlistID, songID, userId, account_type)
	if err != nil {
		return err
	}
	return nil
}
func GetPlaylistUser(userId string) ([]models.Playlist, error) {
	playlists, err := repositories.GetPlaylistUser(userId)
	if err != nil {
		return nil, err
	}
	return playlists, nil
}
func GetPlaylistSongUser(songID string) (models.Song, error) {
	song, err := repositories.GetPlaylistSongUser(songID)
	if err != nil {
		return models.Song{}, err
	}
	return song, nil
}
func CreateCupon(cupon models.Cupon) error {
	err := repositories.CreateCupon(cupon)
	if err != nil {
		return err
	}
	return nil
}
func AssignCupon(cuponID string, userId string) error {
	err := repositories.AssignCupon(cuponID, userId)
	if err != nil {
		return err
	}
	return nil
}
func GetCupon(userId string) ([]models.Cupon, error) {
	cupons, err := repositories.GetCupon(userId)
	if err != nil {
		return nil, err
	}
	return cupons, nil
}
func ListCupons() ([]models.Cupon, error) {
	cupons, err := repositories.ListCupons()
	if err != nil {
		return nil, err
	}
	return cupons, nil
}

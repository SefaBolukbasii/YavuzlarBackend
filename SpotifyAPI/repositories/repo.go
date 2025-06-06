package repositories

import (
	"errors"
	db "spotifyAPI/config"
	"spotifyAPI/models"
)

func Login(username string, password string) (models.User, error) {
	row := db.Db.QueryRow("SELECT t_users.id,t_users.username,t_users.accountType,t_users.cash,t_roles.name FROM t_users  INNER JOIN t_roles ON t_users.roleId=t_roles.id WHERE t_users.username = $1 AND t_users.password = $2", username, password)
	var User models.User
	err := row.Scan(&User.Id, &User.Username, &User.AccountType, &User.Cash, &User.Roles)
	if err != nil {
		return models.User{}, err
	}
	return User, nil
}
func Register(user models.User) error {
	_, err := db.Db.Exec("INSERT INTO t_users (username,password,accountType,cash,roleId) VALUES ($1,$2,$3,$4,(Select id from t_roles where name=$5))", user.Username, user.Password, "free", user.Cash, user.Roles)
	if err != nil {
		return err
	}
	return nil
}
func GetUser(id string) (models.User, error) {
	row := db.Db.QueryRow("SELECT t_users.id,t_users.username,t_users.accountType,t_users.cash,t_roles.name FROM t_users  INNER JOIN t_roles ON t_users.roleId=t_roles.id WHERE t_users.id = $1", id)
	var User models.User
	err := row.Scan(&User.Id, &User.Username, &User.AccountType, &User.Cash, &User.Roles)
	if err != nil {
		return models.User{}, err
	}
	return User, nil
}
func UpdateUser(id string, username string, password string) error {
	_, err := db.Db.Exec("UPDATE t_users SET username = $1, password = $2 WHERE id = $3", username, password, id)
	if err != nil {
		return err
	}
	return nil
}
func DeleteUser(id string) error {
	_, err := db.Db.Exec("DELETE FROM t_users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func GetAllUsers() ([]models.User, error) {
	rows, err := db.Db.Query("SELECT t_users.id,t_users.username,t_users.accountType,t_users.cash,t_roles.name FROM t_users  INNER JOIN t_roles ON t_users.roleId=t_roles.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var User models.User
		err := rows.Scan(&User.Id, &User.Username, &User.AccountType, &User.Cash, &User.Roles)
		if err != nil {
			return nil, err
		}
		users = append(users, User)
	}
	return users, nil
}
func GetPremium(id string, cuponId string) error {
	if cuponId != "nul" {
		_, err := db.Db.Exec("UPDATE t_users SET accountType = 'premium',cash=cash-50 WHERE id = $1", id)
		if err != nil {
			return err
		}
		return nil
	} else {
		row := db.Db.QueryRow("SELECT discount FROM t_cupons WHERE id = $1 and userId=$2", cuponId, id)
		var discount int
		err := row.Scan(&discount)
		if err != nil {
			return err
		}
		kalanCash := 50 - discount
		_, err = db.Db.Exec("DELETE FROM t_cupons WHERE id = $1", cuponId)
		if err != nil {
			return err
		}
		_, err = db.Db.Exec("UPDATE t_users SET accountType = 'premium',cash=cash-$1 WHERE id = $2", kalanCash, id)
		if err != nil {
			return err
		}
		return nil
	}

}
func AddSong(song models.Song) error {
	_, err := db.Db.Exec("INSERT INTO t_songs (name,songerName) VALUES ($1,$2)", song.Name, song.SongerName)
	if err != nil {
		return err
	}
	return nil
}
func DeleteSong(id string) error {
	_, err := db.Db.Exec("DELETE FROM t_songs WHERE id = $1 ", id)
	if err != nil {
		return err

	}
	return nil
}
func UpdateSong(song models.Song, id string) error {
	_, err := db.Db.Exec("UPDATE t_songs SET name=$1,songerName=$2 WHERE id=$3", song.Name, song.SongerName, id)
	if err != nil {
		return err

	}
	return nil
}
func GetAllSongs(page int) ([]models.Song, error) {
	offset := page * 10
	rows, err := db.Db.Query("SELECT id,name,songerName from t_songs ORDER BY clickCount DESC LIMIT 10 OFFSET $1", offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.Id, &song.Name, &song.SongerName)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}
func GetSong(id string) (models.Song, error) {
	row := db.Db.QueryRow("SELECT name,songerName from t_songs WHERE id=$1", id)
	var song models.Song
	err := row.Scan(&song.Name, &song.SongerName)
	if err != nil {
		return models.Song{}, err
	}
	return song, nil
}
func GetMyPlaylists(id string) ([]models.Playlist, error) {
	rows, err := db.Db.Query("SELECT id,name FROM t_playlists WHERE userId=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playlists []models.Playlist
	for rows.Next() {
		var playlist models.Playlist
		err := rows.Scan(&playlist.Id, &playlist.Name)
		if err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}
func GetMyPlaylistID(id string, userId string) (models.Playlist, error) {
	row := db.Db.QueryRow("SELECT id,name FROM t_playlists WHERE id=$1 AND userId=$2", id, userId)
	var playlist models.Playlist
	err := row.Scan(&playlist.Id, &playlist.Name)
	if err != nil {
		return models.Playlist{}, err
	}
	return playlist, nil
}
func AddPlaylist(playlist models.Playlist, userId string) error {
	_, err := db.Db.Exec("INSERT INTO t_playlists (name,userId) VALUES ($1,$2)", playlist.Name, userId)
	if err != nil {
		return err
	}
	return nil
}
func DeletePlaylist(id string, userId string) error {
	_, err := db.Db.Exec("DELETE FROM t_playlists WHERE id = $1 AND userId=$2", id, userId)
	if err != nil {
		return err
	}
	return nil
}
func PlaylistAddSong(playlistId string, songId string, userId string, accountType string) error {
	var KontrolId string
	err := db.Db.QueryRow("SELECT userId FROM t_playlists WHERE id=$1", playlistId).Scan(&KontrolId)
	if err != nil {
		return err
	}
	if KontrolId == userId {
		if accountType == "premium" {
			_, err := db.Db.Exec("INSERT INTO t_playlist_songs (playlistId,songId) VALUES ($1,$2) ", playlistId, songId)
			if err != nil {
				return err
			}
			return nil
		} else if accountType == "free" {
			var songCount int
			err := db.Db.QueryRow("SELECT COUNT(*) FROM t_playlist_songs WHERE playlistId=$1", playlistId).Scan(&songCount)
			if err != nil {
				return err
			}
			if songCount < 5 {
				_, err := db.Db.Exec("INSERT INTO t_playlist_songs (playlistId,songId) VALUES ($1,$2) ", playlistId, songId)
				if err != nil {
					return err
				}
				return nil
			} else {
				return errors.New("you have reached the limit of 5 songs in your playlist")
			}
		} else {
			return errors.New("invalid account type")
		}
	}
	return errors.New("you are not the owner of this playlist")
}
func GetPlaylistUser(userId string) ([]models.Playlist, error) {
	rows, err := db.Db.Query(`
    SELECT 
        t_playlists.id,
        t_playlists.name,
        t_songs.id,
        t_songs.name,
        t_songs.songerName
    FROM 
        t_playlists
    INNER JOIN 
        t_playlist_songs ON t_playlist_songs.playlistId = t_playlists.id
    INNER JOIN 
        t_songs ON t_songs.id = t_playlist_songs.songId
    WHERE 
        t_playlists.userId = $1
`, userId)
	if err != nil {
		return nil, err
	}
	playlistMap := make(map[string]*models.Playlist)
	for rows.Next() {
		var (
			playlistId   string
			playlistName string
			songId       string
			songName     string
			songerName   string
		)
		err := rows.Scan(&playlistId, &playlistName, &songId, &songName, &songerName)
		if err != nil {
			return nil, err
		}
		pl, ok := playlistMap[playlistId]
		if !ok {
			pl = &models.Playlist{
				Id:   playlistId,
				Name: playlistName,
			}
			playlistMap[playlistId] = pl
		}
		pl.Songs = append(pl.Songs, models.Song{
			Id:         songId,
			Name:       songName,
			SongerName: songerName,
		})
	}
	var playlists []models.Playlist
	for _, p := range playlistMap {
		playlists = append(playlists, *p)
	}
	return playlists, nil
}
func GetPlaylistSongUser(songId string) (models.Song, error) {
	row := db.Db.QueryRow("SELECT t_songs.id,t_songs.name,t_songs.songerName from t_songs where t_songs.id=$1", songId)
	var song models.Song
	err := row.Scan(&song.Id, &song.Name, &song.SongerName)
	if err != nil {
		return models.Song{}, err
	}
	return song, nil
}
func CreateCupon(cupon models.Cupon) error {
	if cupon.UserId != "" {
		_, err := db.Db.Exec("INSERT INTO t_cupons (name,discount,userId) VALUES ($1,$2,$3)", cupon.Name, cupon.Discount, cupon.UserId)
		if err != nil {
			return err
		}
	} else {
		_, err := db.Db.Exec("INSERT INTO t_cupons (name,discount) VALUES ($1,$2)", cupon.Name, cupon.Discount)
		if err != nil {
			return err
		}
	}
	return nil
}
func AssignCupon(cuponId string, userId string) error {
	_, err := db.Db.Exec("UPDATE t_cupons SET userId=$1 WHERE id=$2", userId, cuponId)
	if err != nil {
		return err
	}
	return nil
}
func GetCupon(userId string) ([]models.Cupon, error) {
	rows, err := db.Db.Query("SELECT id,name,discount FROM t_cupons WHERE userId=$1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cupons []models.Cupon
	for rows.Next() {
		var cupon models.Cupon
		err := rows.Scan(&cupon.Id, &cupon.Name, &cupon.Discount)
		if err != nil {
			return nil, err
		}
		cupons = append(cupons, cupon)
	}
	return cupons, nil
}
func ListCupons() ([]models.Cupon, error) {
	rows, err := db.Db.Query("SELECT id,name,discount FROM t_cupons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cupons []models.Cupon
	for rows.Next() {
		var cupon models.Cupon
		err := rows.Scan(&cupon.Id, &cupon.Name, &cupon.Discount)
		if err != nil {
			return nil, err
		}
		cupons = append(cupons, cupon)
	}
	return cupons, nil
}

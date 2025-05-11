package repositories

import (
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
func GetPremium(id string) error {
	_, err := db.Db.Exec("UPDATE t_users SET accountType = 'premium',cash=cash-50 WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

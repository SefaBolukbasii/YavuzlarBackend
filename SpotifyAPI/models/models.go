package models

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccountType string `json:"account_type"`
	Cash        int    `json:"cash"`
	Roles       string `json:"roles"`
}
type Song struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	SongerName string `json:"songer_name"`
}
type Playlist struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	UserId string `json:"user_id"`
	Songs  []Song `json:"songs"`
}
type Cupon struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Discount int    `json:"discount"`
	UserId   string `json:"user_id"`
}

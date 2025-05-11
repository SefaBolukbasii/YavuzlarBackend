package models

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccountType string `json:"account_type"`
	Cash        int    `json:"cash"`
	Roles       string `json:"roles"`
}

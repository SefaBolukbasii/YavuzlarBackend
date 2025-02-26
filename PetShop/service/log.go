package service

import (
	"petshop/Database"
)

type LogService struct{}

func (ls *LogService) LogAdd(UserId int, Transaction string) error {
	Veritabani, err := Database.Connect()
	if err != nil {
		return err
	}
	if err := Veritabani.Db.Insert("Log", map[string]any{
		"UserId":      UserId,
		"Transaction": Transaction,
	}); err != nil {
		return err
	}
	return nil
}

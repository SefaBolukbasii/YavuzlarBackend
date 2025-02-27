package service

import (
	"petshop/domain"

	jsondb "github.com/SefaBolukbasii/JsonDB"
)

type LogService struct {
	db *jsondb.Database
}

func CreateLogService(db *jsondb.Database) domain.ILog {
	return &LogService{db: db}
}
func (ls *LogService) LogAdd(UserId int, Transaction string) error {

	if err := ls.db.Insert("Log", map[string]any{
		"UserId":      UserId,
		"Transaction": Transaction,
	}); err != nil {
		return err
	}
	return nil
}

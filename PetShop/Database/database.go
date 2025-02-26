package Database

import jsondb "github.com/SefaBolukbasii/JsonDB"

type Database struct {
	Db *jsondb.Database // jsondb.Database nesnesi
}

func Connect() (*Database, error) {

	db, err := jsondb.DatabaseOlustur("PetShopDB")
	if err != nil {
		return nil, err
	}
	return &Database{Db: db}, nil
}

package Database

import jsondb "github.com/SefaBolukbasii/JsonDB"

func Connect() (*jsondb.Database, error) {

	db, err := jsondb.DatabaseOlustur("PetShopDB")
	if err != nil {
		return nil, err
	}
	return db, nil
}

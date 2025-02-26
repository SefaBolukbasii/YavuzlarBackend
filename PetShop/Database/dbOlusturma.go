package Database

import (
	"fmt"

	jsondb "github.com/SefaBolukbasii/JsonDB"
)

func PetShopDbOlustur() {
	db, err := jsondb.DatabaseOlustur("PetShopDB")
	if err != nil {
		fmt.Print(err)
	}
	kolon0 := []jsondb.Column{
		{Name: "id", Type: jsondb.KendiINT, PrimaryKey: true, AutoIncrement: true},
		{Name: "username", Type: jsondb.KendiSTRING, PrimaryKey: false, AutoIncrement: false},
		{Name: "password", Type: jsondb.KendiSTRING, PrimaryKey: false, AutoIncrement: false},
		{Name: "role", Type: jsondb.KendiSTRING, PrimaryKey: false, AutoIncrement: false},
		{Name: "balance", Type: jsondb.KendiINT, PrimaryKey: false, AutoIncrement: false},
	}
	db.TabloOlustur("Users", kolon0)

	kolon1 := []jsondb.Column{
		{Name: "id", Type: jsondb.KendiINT, PrimaryKey: true, AutoIncrement: true},
		{Name: "ItemName", Type: jsondb.KendiSTRING, PrimaryKey: false, AutoIncrement: false},
		{Name: "Price", Type: jsondb.KendiINT, PrimaryKey: false, AutoIncrement: false},
	}
	db.TabloOlustur("Items", kolon1)

	kolon2 := []jsondb.Column{
		{Name: "id", Type: jsondb.KendiINT, PrimaryKey: true, AutoIncrement: true},
		{Name: "Name", Type: jsondb.KendiSTRING, PrimaryKey: false, AutoIncrement: false},
		{Name: "Species", Type: jsondb.KendiSTRING, PrimaryKey: false, AutoIncrement: false},
		{Name: "OwnerId", Type: jsondb.KendiINT, PrimaryKey: false, AutoIncrement: false},
	}
	db.TabloOlustur("Animals", kolon2)

	kolon3 := []jsondb.Column{
		{Name: "id", Type: jsondb.KendiINT, PrimaryKey: true, AutoIncrement: true},
		{Name: "ItemId", Type: jsondb.KendiINT, PrimaryKey: false, AutoIncrement: false},
		{Name: "UserId", Type: jsondb.KendiINT, PrimaryKey: false, AutoIncrement: false},
		{Name: "TotalPrice", Type: jsondb.KendiINT, PrimaryKey: false, AutoIncrement: false},
	}
	db.TabloOlustur("ItemUser", kolon3)
	kolon4 := []jsondb.Column{
		{Name: "id", Type: jsondb.KendiINT, PrimaryKey: true, AutoIncrement: true},
		{Name: "UserId", Type: jsondb.KendiINT, PrimaryKey: false, AutoIncrement: false},
		{Name: "Transaction", Type: jsondb.KendiSTRING, PrimaryKey: false, AutoIncrement: false},
	}
	db.TabloOlustur("Log", kolon4)
}

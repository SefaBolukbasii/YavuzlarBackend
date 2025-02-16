package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
)

type DataType string

const (
	KendiINT    DataType = "INT"
	KendiSTRING DataType = "STRING"
	KendiBOOL   DataType = "BOOL"
)

type Column struct {
	Name          string   `json:"name"`
	Type          DataType `json:"type"`
	PrimaryKey    bool     `json:"primary_key,omitempty"`
	AutoIncrement bool     `json:"auto_increment,omitempty"`
	DefaultValue  any      `json:"default_value,omitempty"`
}

type Table struct {
	Name    string           `json:"name"`
	Columns []Column         `json:"columns"`
	Data    []map[string]any `json:"data"`
	mutex   sync.Mutex
}

type Database struct {
	Name   string
	Path   string
	Tables map[string]*Table
}

func DatabaseOlustur(name string) (*Database, error) {
	path := filepath.Join("./", name)
	info, err := os.Stat(path)
	if info.IsDir() {
		db := &Database{Name: name, Path: path, Tables: make(map[string]*Table)}
		db.DbGetir()
		return db, nil
	} else {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return nil, err
		}
		return &Database{Name: name, Path: path, Tables: make(map[string]*Table)}, nil
	}

}
func (db *Database) TabloOlustur(tAdi string, columns []Column) error {
	_, varMi := db.Tables[tAdi]
	if varMi {
		return errors.New("tablo zaten var")
	}
	table := &Table{Name: tAdi, Columns: columns, Data: []map[string]any{}}
	db.Tables[tAdi] = table
	return table.save(db.Path)
}
func (db *Database) DbGetir() error {
	dosya, err := os.ReadDir(db.Path)
	if err != nil {
		return err
	}
	for _, a := range dosya {
		if filepath.Ext(a.Name()) == ".json" {
			tAdi := a.Name()[:len(a.Name())-5]
			tablo := &Table{Name: tAdi}
			err = tablo.TabloGetir(db.Path)
			if err != nil {
				return err
			}
			db.Tables[tAdi] = tablo
		}
	}
	return nil
}
func (t *Table) TabloGetir(dbPath string) error {
	dosyaYolu := filepath.Join(dbPath, t.Name+".json")
	dosya, err := os.Open(dosyaYolu)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer dosya.Close()
	return json.NewDecoder(dosya).Decode(&t)

}

func (t *Table) save(dbPath string) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	filePath := filepath.Join(dbPath, t.Name+".json")
	file, err := os.Create(filePath)
	if err != nil {

		return err
	}
	err = json.NewEncoder(file).Encode(t)
	file.Close()
	return err
}
func (db *Database) Insert(tabloAdi string, degerler map[string]any) error {
	t, varMi := db.Tables[tabloAdi]
	if !varMi {
		return errors.New("boyle bir tabbblo yok")
	}
	t.mutex.Lock() // nesneyi kilitler çakışma olmasın die

	for _, col := range t.Columns {
		if col.PrimaryKey { //primary key kontrol kısmı
			if col.Type == KendiINT && col.AutoIncrement {
				degerler[col.Name] = len(t.Data) + 1
				break
			} else if col.Type == KendiSTRING {
				degerler[col.Name] = uuid.NewString()
				break
			}
		}
	}
	t.Data = append(t.Data, degerler)
	t.mutex.Unlock()
	return t.save(db.Path)
}

func (db *Database) Select(tAdi string) ([]map[string]any, error) {
	t, varMi := db.Tables[tAdi]
	if !varMi {
		return nil, errors.New("Tablo yok")
	}
	return t.Data, nil
}

func (db *Database) Delete(tAdi string, key string, value any) error {
	t, varMi := db.Tables[tAdi]
	if !varMi {
		return errors.New("Tablo Yok")
	}
	t.mutex.Lock()

	for i, row := range t.Data {
		if row[key] == value {
			t.Data = append(t.Data[:i], t.Data[i+1:]...)
		}
	}
	t.mutex.Unlock()
	return t.save(db.Path)
}
func (db *Database) Update(tAdi string, key string, eskiValue any, yeniValue any) error {
	t, varMi := db.Tables[tAdi]
	if !varMi {
		return errors.New("Tablo Yok")
	}
	t.mutex.Lock()

	for i, row := range t.Data {
		if row[key] == eskiValue {
			t.Data[i][key] = yeniValue
		}
	}
	t.mutex.Unlock()
	return t.save(db.Path)
}
func main() {
	db, err := DatabaseOlustur("db1")
	if err != nil {
		fmt.Println("olmadı")
	}
	db.Update("user", "name", "Mehmet", "Mahmut")

}

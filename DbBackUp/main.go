package main

import (
	"fmt"
	"os/exec"

	argParser "github.com/SefaBolukbasii/ArgumentParser"
)

func MysqlBackUp(user, password, dbName, filepath string) error {
	cmd := exec.Command("mysqldump", "-u", user, fmt.Sprintf("-p%s", password), "--databases", dbName, "-r", filepath)
	return cmd.Run()
}
func MysqlLoad(user, password, dbName, dumpFile string) error {
	cmd := exec.Command("mysql", "-u", user, fmt.Sprintf("-p%s", password), dbName, "-e", fmt.Sprintf("source %s", dumpFile))
	return cmd.Run()
}
func PostgreBackUp(user, password, dbName, filepath string) error {
	cmd := exec.Command("pg_dump", "-U", user, "-F", "c", "-f", filepath, dbName)
	cmd.Env = append(cmd.Env, fmt.Sprintf("PGPASSWORD=%s", password))
	return cmd.Run()
}
func PostgreLoad(user, password, dbName, filepath string) error {
	cmd := exec.Command("pg_restore", "-U", user, "-d", dbName, filepath)
	cmd.Env = append(cmd.Env, fmt.Sprintf("PGPASSWORD=%s", password))
	return cmd.Run()
}
func SqlLiteBackup(dbName, filepath string) error {
	cmd := exec.Command("sqlite3", dbName, fmt.Sprintf(".backup '%s'", filepath))
	return cmd.Run()
}
func SqlLiteLoad(dbName, filepath string) error {
	cmd := exec.Command("cp", filepath, dbName)
	return cmd.Run()
}
func main() {
	argParser.AddArgument(&argParser.Argument{
		Name:         "user",
		Description:  "kullanıcı adını belirtir",
		Example:      "--user Sefa",
		Forced:       true,
		ArgumentType: argParser.CommandType,
	})
	argParser.AddArgument(&argParser.Argument{
		Name:         "password",
		Description:  "şifre belirtir",
		Example:      "--password 123Asd",
		Forced:       true,
		ArgumentType: argParser.CommandType,
	})
	argParser.AddArgument(&argParser.Argument{
		Name:         "db-type",
		Description:  "hedef veritabanını belirtir (mysql,postgresql,sqlite) olarak 3 seçenek mevcuttur",
		Example:      "--db-type <dbType>",
		Forced:       true,
		ArgumentType: argParser.CommandType,
	})
	argParser.AddArgument(&argParser.Argument{
		Name:         "backup",
		Description:  "yedekleme işlemi yapılacağını belirtir",
		Example:      "--backup",
		Forced:       false,
		ArgumentType: argParser.OptionType,
	})
	argParser.AddArgument(&argParser.Argument{
		Name:         "backup-dir",
		Description:  "yedekleme dosyasının saklanacağı ana yedekleme dizinini belirtir",
		Example:      "--backup-dir <path>",
		Forced:       false,
		ArgumentType: argParser.CommandType,
	})
	argParser.AddArgument(&argParser.Argument{
		Name:         "load",
		Description:  "yedek kullanılarak geri yükleme yapılacağı belirtilir",
		Example:      "--load",
		Forced:       false,
		ArgumentType: argParser.OptionType,
	})
	argParser.AddArgument(&argParser.Argument{
		Name:         "dump-file",
		Description:  "geri yükleme işlemi için kullanılıcak dump dosyasının dosya yolunu belirtir",
		Example:      "--dump-file <path>",
		Forced:       false,
		ArgumentType: argParser.CommandType,
	})
	argParser.AddArgument(&argParser.Argument{
		Name:         "db-name",
		Description:  "geri yükleme işlemi için kullanılıcak dump dosyasının dosya yolunu belirtir",
		Example:      "--db-name <db_name>",
		Forced:       true,
		ArgumentType: argParser.CommandType,
	})
	Arguments, err := argParser.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	UserName := Arguments["user"]
	password := Arguments["password"]
	db_type := Arguments["db-type"]
	db_name := Arguments["db-name"]
	backup := Arguments["backup"]
	load := Arguments["load"]
	if backup == true {
		path := Arguments["backup-dir"]
		if db_type == "mysql" {
			if err := MysqlBackUp(UserName.(string), password.(string), db_name.(string), path.(string)); err != nil {
				fmt.Println(err)
				return
			}
		} else if db_type == "postgresql" {
			if err := PostgreBackUp(UserName.(string), password.(string), db_name.(string), path.(string)); err != nil {
				fmt.Println(err)
				return
			}
		} else if db_type == "sqlite" {
			if err := SqlLiteBackup(db_name.(string), path.(string)); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("Geçerli bir veri tabanı türü giriniz")
		}
	} else if load == true {
		path := Arguments["dump-file"]
		if db_type == "mysql" {
			if err := MysqlLoad(UserName.(string), password.(string), db_name.(string), path.(string)); err != nil {
				fmt.Println(err)
				return
			}
		} else if db_type == "postgresql" {
			if err := PostgreLoad(UserName.(string), password.(string), db_name.(string), path.(string)); err != nil {
				fmt.Println(err)
				return
			}
		} else if db_type == "sqlite" {
			if err := SqlLiteLoad(db_name.(string), path.(string)); err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("Geçerli bir veri tabanı türü giriniz")
		}
	}
}

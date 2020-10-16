package internal

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	v "github.com/spf13/viper"
)

func MakeDatabase(remote bool, db_name string) {
	if remote {
		err := makeRemote(db_name)
		if err != nil {
			throwError(err)
		}
	} else {
		err := makeLocal(db_name)
		if err != nil {
			throwError(err)
		}
	}
}

func makeLocal(db_name string) error {
	name := v.GetString("local.user")
	password := v.GetString("local.password")

	db, err := sql.Open("mysql", name+":"+password+"@tcp(localhost)/")
	command := "CREATE DATABASE " + db_name

	if err != nil {
		return err
	} else {
		_, err := db.Exec(command)

		if err != nil {
			return err
		}
	}

	db.Close()
	return nil
}

func makeRemote(db_name string) error {
	// host := v.GetString("remote.host")
	name := v.GetString("remote.user")
	password := v.GetString("remote.password")

	db, err := sql.Open("mysql", name+":"+password+"@tcp(localhost)/")
	command := "CREATE DATABASE " + db_name

	if err != nil {
		return err
	} else {
		_, err := db.Exec(command)

		if err != nil {
			return err
		}
	}

	db.Close()
	return nil
}

func throwError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

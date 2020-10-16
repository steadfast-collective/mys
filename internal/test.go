package internal

import (
	"database/sql"
	"fmt"

	v "github.com/spf13/viper"
)

func TestRemote() error {
	host := v.GetString("remote.host")
	name := v.GetString("remote.user")
	password := v.GetString("remote.password")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", name, password, host, "3306"))
	if err != nil {
		return err
	} else {
		fmt.Println("connected!!")
	}
	db.Close()
	return nil
}

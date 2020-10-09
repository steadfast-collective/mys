package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DropDatabase(db_name string, name string, password string) {
	db, err := sql.Open("mysql", name+":"+password+"@tcp(localhost)/")
	command := "DROP DATABASE " + db_name

	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err := db.Exec(command)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("Successfully dropped database '%s'\n", db_name)
		}
	}
}

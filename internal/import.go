package internal

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
)

func ImportDatabase(db_name string, file_name string, name string, password string) {
	cmdString := "mysql -u" + name
	pwString := "-p " + password
	if len(password) == 0 {
		pwString = ""
	}
	cmd := exec.Command("bash", "-c", cmdString+pwString+" "+db_name+" < "+file_name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

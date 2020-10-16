package internal

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
	v "github.com/spf13/viper"
)

var (
	configCMD        *flag.FlagSet
	makeCMD          *flag.FlagSet
	dropCMD          *flag.FlagSet
	importCMD        *flag.FlagSet
	dumpCMD          *flag.FlagSet
	testCMD          *flag.FlagSet
	is_remote        bool
	db_name          string
	file_name        string
	destination_name string
)

func init() {

	testCMD = flag.NewFlagSet("test", flag.ExitOnError)

	configCMD = flag.NewFlagSet("config", flag.ExitOnError)

	makeCMD = flag.NewFlagSet("make", flag.ExitOnError)
	makeCMD.BoolVarP(&is_remote, "remote", "r", false, "Should mys connect to the remote mysql")
	makeCMD.StringVarP(&db_name, "database", "d", "", "The name of the database to create / drop / import to depending on the command")

	dropCMD = flag.NewFlagSet("drop", flag.ExitOnError)
	dropCMD.BoolVarP(&is_remote, "remote", "r", false, "Should mys connect to the remote mysql")
	dropCMD.StringVarP(&db_name, "database", "d", "", "The name of the database to drop")

	importCMD = flag.NewFlagSet("import", flag.ExitOnError)
	importCMD.StringVarP(&db_name, "database", "d", "", "The name of the database to import to")
	importCMD.StringVarP(&file_name, "file", "f", "", "The name of the SQL file to import")

	dumpCMD = flag.NewFlagSet("dump", flag.ExitOnError)
	dumpCMD.StringVarP(&db_name, "database", "d", "", "The name of the database to dump")
	dumpCMD.StringVarP(&destination_name, "output", "o", "", "The location to output the SQL dump to")
}

func RunCmd(command string) error {
	name := v.GetString("local.user")
	pass := v.GetString("local.password")

	switch command {
	case "test":
		err := TestRemote()
		if err != nil {
			fmt.Println(err.Error())
		}
	case "config":
		WriteConfig()
	case "make":
		makeCMD.Parse(os.Args[2:])
		MakeDatabase(is_remote, db_name)
	case "drop":
		dropCMD.Parse(os.Args[2:])
		DropDatabase(db_name, name, pass)
	case "import":
		importCMD.Parse(os.Args[2:])
		fn, _ := filepath.Abs(file_name)
		ImportDatabase(db_name, fn, name, pass)
	case "dump":
		dumpCMD.Parse(os.Args[2:])
		DumpDatabase(db_name, destination_name, name, pass)
	default:
		fmt.Println("Please supply an arg")
	}
	return nil
}

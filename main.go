package main

import (
	"errors"
	"fmt"
	"mys/internal"
	"os"

	"github.com/manifoldco/promptui"

	flag "github.com/spf13/pflag"
	v "github.com/spf13/viper"
)

var (
	makeCMD          *flag.FlagSet
	dropCMD          *flag.FlagSet
	importCMD        *flag.FlagSet
	dumpCMD          *flag.FlagSet
	db_name          string
	file_name        string
	destination_name string
)

func main() {
	parseCommand()
}

func init() {
	err := scaffoldConfig()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	validateUsername := func(input string) error {
		if len(input) == 0 {
			return errors.New("Invalid MySQL username")
		}
		return nil
	}
	config_prompt := promptui.Prompt{
		Label:    "Local MySQL username",
		Validate: validateUsername,
	}

	result, err := config_prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	fmt.Printf("You chose %q\n", result)

	makeCMD = flag.NewFlagSet("make", flag.ExitOnError)
	makeCMD.StringVarP(&db_name, "database", "d", "", "The name of the database to create / drop / import to depending on the command")

	dropCMD = flag.NewFlagSet("drop", flag.ExitOnError)
	dropCMD.StringVarP(&db_name, "database", "d", "", "The name of the database to drop")

	importCMD = flag.NewFlagSet("import", flag.ExitOnError)
	importCMD.StringVarP(&db_name, "database", "d", "", "The name of the database to import to")
	importCMD.StringVarP(&file_name, "file", "f", "", "The name of the SQL file to import")

	dumpCMD = flag.NewFlagSet("dump", flag.ExitOnError)
	dumpCMD.StringVarP(&db_name, "database", "d", "", "The name of the database to dump")
	dumpCMD.StringVarP(&destination_name, "output", "o", "", "The location to output the SQL dump to")

	if len(os.Args) < 2 {
		fmt.Println("Expected a valid command, run -h | --help for usage")
		os.Exit(1)
	}
}

func parseCommand() {
	name := v.GetString("name")
	pass := v.GetString("password")

	switch os.Args[1] {
	case "make":
		makeCMD.Parse(os.Args[2:])
		internal.MakeDatabase(db_name, name, pass)
	case "drop":
		dropCMD.Parse(os.Args[2:])
		internal.DropDatabase(db_name, name, pass)
	}
}

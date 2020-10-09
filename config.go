package main

import (
	"fmt"

	v "github.com/spf13/viper"
)

func scaffoldConfig() error {
	v.SetConfigName(".mysconfig")
	v.SetConfigType("toml")
	v.AddConfigPath("$HOME")

	v.SetDefault("name", "root")
	v.SetDefault("password", "")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("%s\n Lets create one now!\n", err)
		err := v.SafeWriteConfig()
		if err != nil {
			return err
		}
		fmt.Println("Config created at $HOME/.mysconfig.toml")
	}
	return nil
}

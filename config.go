package main

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
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

func writeConfig() {
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
	v.Set("name", result)
}

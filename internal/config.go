package internal

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	v "github.com/spf13/viper"
)

func ScaffoldConfig() error {
	v.SetConfigName(".mysconfig")
	v.SetConfigType("yaml")
	v.AddConfigPath("$HOME")

	v.SetDefault("local.user", "root")
	v.SetDefault("local.password", "")
	v.SetDefault("remote.host", "")
	v.SetDefault("remote.user", "")
	v.SetDefault("remote.password", "")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("%s\nLets create one now!\n", err)
		err := v.SafeWriteConfig()
		if err != nil {
			return err
		}
		fmt.Println("Config created at $HOME/.mysconfig.yaml")
	}
	return nil
}

func WriteConfig() {
	setUsername()
	setPassword()
	prompt := promptui.Prompt{
		Label:     "Configure remote connection?",
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("you chose %q\n", result)
	v.WriteConfig()
}

func setUsername() {
	validateUsername := func(input string) error {
		if len(input) == 0 {
			return errors.New("Invalid MySQL username")
		}
		return nil
	}
	template := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	config_prompt := promptui.Prompt{
		Label:     "Local MySQL username: ",
		Templates: template,
		Validate:  validateUsername,
	}

	result, err := config_prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	v.Set("local.user", result)
}

func setPassword() {
	config_prompt := promptui.Prompt{
		Label: "Local MySQL password",
		Mask:  '*',
	}

	result, err := config_prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}
	v.Set("local.password", result)
}

func setRemote() {
	fmt.Println("here goes the remote logic")
}

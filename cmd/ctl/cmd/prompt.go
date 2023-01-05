package cmd

import "github.com/manifoldco/promptui"

func Select(title string, items []string) (string, error) {
	prompt := promptui.Select{
		Label: title,
		Items: items,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

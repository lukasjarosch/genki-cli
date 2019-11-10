package cli

import "github.com/AlecAivazis/survey/v2"

func Confirm(message string) bool {
	var confirm bool
	prompt := &survey.Confirm{
		Message: message,
		Default: false,
	}
	_ = survey.AskOne(prompt, &confirm)
	return confirm
}

func Select(message string, options ...string) (string, error) {
	var selected string
	prompt := &survey.Select{
		Message: message,
		Options:options,
	}
	err := survey.AskOne(prompt, &selected)
	if err != nil {
		return "", err
	}
	return selected, nil
}

func Prompt(message string) (string, error) {
	var input string
	prompt := &survey.Input{
		Message:  message,
	}
	err := survey.AskOne(prompt, &input)
	if err != nil {
	    return "", err
	}
	return input, nil
}


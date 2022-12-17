package app

import (
	"fmt"

	"jlpt/src/utils"

	"github.com/manifoldco/promptui"
)

func (app *AppImpl) SelectLevel() error {
	prompt := promptui.Select{
		Label: "Select Level",
		Items: []string{"JLPT N-5", "JLPT N-4", "JLPT N-3", "JLPT N-2", "JLPT N-1"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	level, err := utils.MapLevelToEnum(result)
	if err != nil {
		return err
	}

	app.Level = *level

	return nil
}

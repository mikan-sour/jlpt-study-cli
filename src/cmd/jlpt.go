package cmd

import (
	"context"
	"fmt"
	"os"

	"jlpt/src/app"
	"jlpt/src/config"
	"jlpt/src/game"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "It's an app to study JLPT words",
	Long:  `It's an app to study JLPT words but this is a longer explanation`,
	Run: func(cmd *cobra.Command, args []string) {

		application := app.New(*config.New())

		mgs := application.GetMatchingGameService()

		game := game.NewGame(mgs)

		err := game.Run(context.Background(), nil)
		if err != nil {
			fmt.Println(errors.Wrapf(err, "game.Run failed: %s", err.Error()))
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}

// func Execute() {
// 	if err := rootCmd.Execute(); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
// }

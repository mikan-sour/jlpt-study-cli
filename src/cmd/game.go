/*
Copyright © 2022 cli-jlpt jedzeins@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// gameCmd represents the game command
var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "Play a game!",
	Long:  `Play a game to study JLPT!`,
}

func init() {
	rootCmd.AddCommand(gameCmd)
}

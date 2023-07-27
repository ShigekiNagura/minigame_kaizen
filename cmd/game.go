/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ShigekiNagura/minigame_kaizen/game"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// gachaCmd represents the gacha command
var gachaCmd = &cobra.Command{
	Use:   "game",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "どのゲームで遊びますか？",
			Items: []string{"○×ゲーム", "ガチャ", "ガチャ(サンプル)", "クイズ"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "○×ゲーム":
			game.Marubatsu()
		case "ガチャ":
			game.Gacha()
		case "ガチャ(サンプル)":
			game.GachaSample()
		case "クイズ":
			game.Quiz()
		}
	},
}

func init() {
	rootCmd.AddCommand(gachaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gachaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gachaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

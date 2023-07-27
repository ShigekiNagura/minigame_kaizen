package game

import (
	"fmt"
	"time"

	"github.com/manifoldco/promptui"
)

func Quiz() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("問題です！")
	sound("quiz_shutsudai")

	prompt := promptui.Select{
		Label: "ラクスの設立年はいつ？",
		Items: []string{"1. 1999年", "2. 2000年", "3. 2001年", "4. 2002年"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "2. 2000年" {
		fmt.Println("正解！")
		sound("quiz_seikai")
	} else {
		fmt.Println("不正解！")
		sound("quiz_huseikai")
	}
}

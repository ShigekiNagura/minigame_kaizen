package game

import (
	"fmt"
	"time"

	"github.com/manifoldco/promptui"
)

func GachaSample() {
	prompt := promptui.Select{
		Label: "どのレアリティのキャラクターを出力しますか？",
		Items: []string{"3（60%）", "4（25%）", "5（10%）", "6（4%）", "7（1%）"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var charactor Charactor

	switch result {
	case "3（60%）":
		charactor = charactorLevel3[0]
	case "4（25%）":
		charactor = charactorLevel4[0]
	case "5（10%）":
		charactor = charactorLevel5[0]
	case "6（4%）":
		charactor = charactorLevel6[0]
	case "7（1%）":
		charactor = charactorLevel7[0]
	}

	fmt.Println("ガチャン")
	// sound("door_lock")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("ガラガラ...")
	// sound("gatyagatya")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("ポン！")
	// sound("level_up")
	time.Sleep(1500 * time.Millisecond)

	fmt.Println(charactor.Image)
	fmt.Println("なまえ : ", charactor.Name)
	fmt.Println("レアリティ : ", charactor.Rarity)
	sound(charactor.SoundFile)

}

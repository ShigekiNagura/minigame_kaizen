package game

import (
	"fmt"
	"time"
)

func Gacha() {

	fmt.Println("排出率：⭐︎3（60%）・⭐︎4（25%）・⭐︎5（10%）・⭐︎6（4%）・⭐︎7（1%）")
	fmt.Println("ガチャを回す。Enterを押してください。")
	fmt.Scanln()
	fmt.Println("ガチャン")
	// sound("door_lock")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("ガラガラ...")
	// sound("gatyagatya")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("ポン！")
	// sound("level_up")
	time.Sleep(1500 * time.Millisecond)

	charactor := generateCharactor()

	fmt.Println(charactor.Image)
	fmt.Println("なまえ : ", charactor.Name)
	fmt.Println("レアリティ : ", charactor.Rarity)
	sound(charactor.SoundFile)

}

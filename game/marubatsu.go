package game

import "fmt"

func Marubatsu() {
	a := [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}

	fmt.Println("以下は盤面です。")
	fmt.Println("0|1|2")
	fmt.Println("3|4|5")
	fmt.Println("6|7|8")
	fmt.Println("-----------------")
	fmt.Println()

	for i := 0; ; i++ {

		// ターン表示
		if i%2 == 0 {
			fmt.Println("oのターンです")
		} else {
			fmt.Println("xのターンです")
		}
		fmt.Println("どこに置きますか？0~8の中から選んでください。")
		var n int
		fmt.Scan(&n)

		// 0~8以外の数字が入力された場合はやり直し
		if n < 0 || n > 8 {
			fmt.Println("0~8の数字を入力してください")
			i--
			continue
		}

		// すでに置かれている場合はやり直し
		if a[n] != " " {
			fmt.Println("既に置かれています")
			i--
			continue
		}

		// oとxの交互に置く
		if i%2 == 0 {
			a[n] = "o"
		} else {
			a[n] = "x"
		}

		// 盤面表示
		for i := 0; i < 9; i++ {
			if i%3 == 0 && i != 0 {
				fmt.Print("|")
				fmt.Println()
			}
			fmt.Print("|")
			fmt.Print(a[i])
			if i == 8 {
				fmt.Print("|")
			}
		}
		fmt.Println()

		// 勝敗判定
		if a[0] == a[1] && a[1] == a[2] && a[0] != " " {
			fmt.Println(a[0] + "の勝ちです")
			break
		}
		if a[3] == a[4] && a[4] == a[5] && a[3] != " " {
			fmt.Println(a[3] + "の勝ちです")
			break
		}
		if a[6] == a[7] && a[7] == a[8] && a[6] != " " {
			fmt.Println(a[6] + "の勝ちです")
			break
		}
		if a[0] == a[3] && a[3] == a[6] && a[0] != " " {
			fmt.Println(a[0] + "の勝ちです")
			break
		}
		if a[1] == a[4] && a[4] == a[7] && a[1] != " " {
			fmt.Println(a[1] + "の勝ちです")
			break
		}
		if a[2] == a[5] && a[5] == a[8] && a[2] != " " {
			fmt.Println(a[2] + "の勝ちです")
			break
		}
		if a[0] == a[4] && a[4] == a[8] && a[0] != " " {
			fmt.Println(a[0] + "の勝ちです")
			break
		}
		if a[2] == a[4] && a[4] == a[6] && a[2] != " " {
			fmt.Println(a[2] + "の勝ちです")
			break
		}

		// 9回目までに勝敗が決まらなかったら引き分け
		if i == 8 {
			fmt.Println("引き分けです")
			break
		}
	}
}

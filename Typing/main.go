package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	//"time"
)

func main() {
	fmt.Println("タイピングゲームが始まった！頑張って！")
	var num int
	for i := 0; i < 5; i++ {
		word := randomWord()
		fmt.Println("以下の英単語をタイピングしてください：", word)
		right := compare(word)
		num = num + right
	}
	fmt.Println("ゲーム終了！　正確単語数：", num)
}

//単語を比較して、スコアを計算する
func compare(word string) (right int) {
	score := 0
	num := 0
	inputword := imp()
	if len(inputword) == len(word) {
		for i := 0; i < len(inputword); i++ {
			wd1 := inputword[i]
			wd2 := word[i]
			if wd1 == wd2 {
				num++
			}
		}
	}
	if num == len(word) {
		score = score + 1
	}
	right = score
	return right
}

//英単語をランダムに取り出す
func randomWord() (word string) {
	words := []string{"banana", "apple", "milk", "fruits", "cat", "car", "elephant", "unbralla", "interface", "tissues", "format"}
	num := rand.Intn(10)
	return words[num]
}

//入力単語を取得する
func imp() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput = scanner.Text()
	stringInput = strings.TrimSpace(stringInput)
	return stringInput
}

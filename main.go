package main

import (
	"fmt"
	"strings"
)

func decryptRailFence(cipher string, key int) string {
	rail := [][]string{}
	splString := strings.Split(cipher, "")
	for i := 0; i < key; i++ {
		rail = append(rail, []string{})
		for j := 0; j < len(splString); j++ {
			rail[i] = append(rail[i], "/n")
		}
	}

	dir_down := true
	row, col := 0, 0

	for i := 0; i < len(splString); i++ {
		if row == 0 {
			dir_down = true
		}
		if row == key-1 {
			dir_down = false
		}
		rail[row][col] = "*"
		col += 1
		if dir_down {
			row += 1
		} else {
			row -= 1
		}
	}

	index := 0
	for i := 0; i < key; i++ {
		for j := 0; j < len(splString); j++ {
			if (rail[i][j] == "*") && (index < len(cipher)) {
				rail[i][j] = splString[index]
				index += 1
			}
		}
	}

	result := []string{}
	row, col = 0, 0

	for i := 0; i < len(splString); i++ {
		if row == 0 {
			dir_down = true
		}
		if row == key-1 {
			dir_down = false
		}

		if rail[row][col] != "*" {
			result = append(result, rail[row][col])
			col += 1
		}

		if dir_down {
			row += 1
		} else {
			row -= 1
		}
	}

	return strings.Join(result, "")
}

func main() {
	listOfCountChar := make(map[string]int)
	text := "CYtZBsWZaZliYZocWLZlXuZZYWYeYXZsXeZXtXWpXeRYYYd!ZnYeWXoYXasnX,WXWrWPoAdWesnciGenWr"
	secret := strings.Split(text, "")

	mostV := 0
	for i := 0; i < len(secret); i++ {
		listOfCountChar[secret[i]] += 1
		if mostV < listOfCountChar[secret[i]] {
			mostV = listOfCountChar[secret[i]]
		}
	}

	for ch, i := range listOfCountChar {
		if mostV == i {
			text = strings.ReplaceAll(text, ch, "")
		}
	}

	fmt.Println(decryptRailFence(text, 4))
}

package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

//ref. https://www.geeksforgeeks.org/how-to-reverse-a-string-in-golang/
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)

	var temp = strings.Split(string(sd), ":")

	for _, s := range temp {
		whatIsIt = reverse(s) + " " + whatIsIt
	}
	fmt.Println(whatIsIt)
}

//iangnoW:NAM:ENIL:ta:su:nioJ
//Wongnai:MAN:LINE:at:us:Join
//Join us at LINE MAN Wongnai

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var L int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L)

	var H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &H)

	scanner.Scan()
	T := scanner.Text()
	_ = T // to avoid unused error

	ASCII := []string{}
	answer := []string{}
	for i := 0; i < H; i++ { // get the ascii art
		scanner.Scan()
		ROW := scanner.Text()
		ASCII = append(ASCII, ROW)
	}

	alphabets := "ABCDEFGHIJKLMNOPQRSTUVWXYZ?"
	var indexes []int

	for _, t := range T { // find the letters' indexes
		index := -1
		for i, a := range alphabets {
			ut := unicode.ToUpper(t)
			if ut == a {
				index = i
			}
		}

		if index == -1 {
			indexes = append(indexes, len(alphabets)-1)
		} else {
			indexes = append(indexes, index)
		}
	}

	for i := 0; i < H; i++ { // initialize string slice
		answer = append(answer, "")
	}

	for _, i := range indexes { //find the ascii art for the given letters
		for j, ascii := range ASCII {
			answer[j] += ascii[i*L : i*L+L]
		}
	}

	for _, str := range answer {
		fmt.Println(str)
	}

}

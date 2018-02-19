package main

// Project rand-string
// file: main.go
// Author: Matt Weidner <matt.weidner@gmail.com>
// Description: random string generator

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const RANDOM_STR_LEN = 20

func shuffleDict(dict string) []rune {
	shuffledDict := []rune(dict)
	for i := 0; i < len(dict); i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(len(dict))))
		randn := r.Int64()
		char := shuffledDict[i]
		shuffledDict[i] = shuffledDict[randn]
		shuffledDict[randn] = char
	}
	return shuffledDict
}

func main() {
	fmt.Println("Here's 10 random strings for your enjoyment:")
	var dictionary string = "abcdefghiklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRZTUVWXYZ0123456789"
	for k := 0; k < 10; k++ {
		shuffledDict := shuffleDict(dictionary)
		for j := 0; j < RANDOM_STR_LEN; j++ {
			i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(dictionary))))
			fmt.Printf("%c", shuffledDict[i.Int64()])
		}
		fmt.Printf("\n")
	}
}

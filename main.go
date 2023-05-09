package main

import (
	"fmt"
	"strings"
)

func main() {
	kata := "Selamat Malam"
	kata = strings.ToLower(kata)
	jumlahHuruf := make(map[int32]int)
	for _, huruf := range kata {
		fmt.Printf("%c\n", huruf)
		jumlahHuruf[huruf]++
	}
	result := "Map["
	for huruf, jumlah := range jumlahHuruf {
		result += fmt.Sprintf("%c :%d ", huruf, jumlah)
	}
	result += "]"
	fmt.Println(result)
}

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
func main() {
	r1 :=[26]byte{'E','K','M','F','L','G','D','Q','V','Z','N','T','O','W','Y','H','X','U','S','P','A','I','B','R','C','J'}
	r2 :=[26]byte{'A','J','D','K','S','I','R','U','X','B','L','H','W','T','M','C','Q','G','Z','N','P','Y','F','V','O','E'}
	r3 :=[26]byte{'B','D','F','H','J','L','C','P','R','T','X','V','Z','N','Y','E','I','W','G','A','K','M','U','S','Q','O'}
	fmt.Println("Hello world")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ToUpper(text)
	fmt.Println(text)
	
	var slice []byte
	for _,c := range text {
		if c == '\n' {break}
		var b byte = r1[c-'A']
		slice = append(slice, b)

	}
	fmt.Println(string(slice[:]))
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
}

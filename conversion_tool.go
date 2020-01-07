package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("====================================")
	fmt.Println("Welcome to my Go Conversion Tool")
	fmt.Println("====================================")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(`
Make a selection of the type of conversion you want by entering the number.
1. Length
2. Area
3. Volume
4. Temperature
5. Speed
6. Time
7. Mass
8. Numeral System
`)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

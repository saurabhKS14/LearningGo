package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Tell us your rankings : \n")
	input, _ := reader.ReadString('\n')
	fmt.Println("Rankings : ", input)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give the pizza rating from 1 to 5")

	input, _ := reader.ReadString('\n')

	inputIntegerValue, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("After adding 1 to your rating : ", inputIntegerValue+1)
	}

}

package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(strings.ToLower(input))

	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

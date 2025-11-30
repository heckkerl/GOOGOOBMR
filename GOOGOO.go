package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your goo goo mummy name: ")
	name , _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("Enter your goo goo meter: ")
	meter := 0; fmt.Scanln(&meter)
	for i := 0; i <= meter; i++{
		fmt.Printf("GREAT MOTHER %s has been GOO GOO you %d times 🍼\n", name, i)
	}
}

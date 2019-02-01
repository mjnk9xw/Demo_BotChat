package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aichaos/rivescript-go"
)

func main() {
	bot := rivescript.New(rivescript.WithUTF8())
	// err := bot.LoadDirectory("eg/brain")
	// if err != nil {
	// 	fmt.Printf("Error loading from directory: %s", err)
	// }
	// bot.SortReplies()

	err := bot.LoadFile("./demo.rive")
	if err != nil {
		fmt.Printf("Error loading from file: %s", err)
	}
	bot.SortReplies()

	for {
		fmt.Print("You> ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")
		if len(input) == 0 {
			break
		}

		reply, err := bot.Reply("Minh: ", input)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("\nBot> %s\n\n", reply)
		}
	}
	return
}

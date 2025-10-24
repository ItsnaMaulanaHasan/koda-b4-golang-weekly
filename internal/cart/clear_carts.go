package cart

import (
	"bufio"
	"fmt"
	"golang-weekly/internal/models"
	"os"
	"strings"
)

func ClearCart() {
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	for loop {
		fmt.Println("\x1bc")
		fmt.Print("Are you sure you want to clear carts (y/n)? ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		if strings.ToLower(choiceStr) == "y" {
			models.Carts = []models.CartItem{}
			fmt.Print("Carts successfully cleared! Press enter to continue... ")
			scanner.Scan()
			loop = false
		} else if strings.ToLower(choiceStr) == "n" {
			loop = false
		} else {
			fmt.Print("Invalid input, please enter 'y' or 'n'... ")
			scanner.Scan()
		}
	}
}

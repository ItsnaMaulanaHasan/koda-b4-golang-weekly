package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

func ClearCache(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer) {
	loop := true
	for loop {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Print(r)
					scanner.Scan()
				}
			}()
			fmt.Println("\x1bc")
			fmt.Print("Are you sure you want to clear cache (y/n)? ")
			choiceStr := InputString(reader)
			if strings.ToLower(choiceStr) == "y" {
				tempPath := os.TempDir()
				tempDirPath := filepath.Join(tempPath, "mixue-pos")

				err := os.RemoveAll(tempDirPath)
				if err != nil {
					panic(err)
				}

				fmt.Print("Cache cleared successfully! Press enter to continue... ")
				scanner.Scan()
				loop = false
				return
			} else if strings.ToLower(choiceStr) == "n" {
				loop = false
				return
			} else {
				panic("Invalid input, please enter 'y' or 'n'... ")
			}
		}()
	}
}

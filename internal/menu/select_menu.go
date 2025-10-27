package menu

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang-weekly/internal/models"
	"golang-weekly/internal/utils"
	"os"
	"path/filepath"
	"text/tabwriter"
	"time"
)

func getDataMenu(tempFilePath *string) {
	// get data menu
	dataMenu := utils.GetData("https://raw.githubusercontent.com/ItsnaMaulanaHasan/koda-b4-golang-weekly-data/refs/heads/main/data.json")

	// unmarshall dataMenu ke dalam bentu slice of struct
	err := json.Unmarshal(dataMenu, &models.Menus)
	if err != nil {
		panic(err)
	}

	// membuat file temp dan mengisinya dengan dataMenu
	err = os.WriteFile(*tempFilePath, dataMenu, 0666)
	if err != nil {
		panic(err)
	}
}

func cachingDataMenu() {
	tempPath := os.TempDir()

	// membuat directory file temp
	mixuePosDir := filepath.Join(tempPath, "mixue-pos")
	err := os.MkdirAll(mixuePosDir, 0755)
	if err != nil {
		panic(err)
	}

	tempFilePath := filepath.Join(mixuePosDir, "menu.json")

	// cek apakah file Temp ada
	fileTemp, err := os.Stat(tempFilePath)
	if os.IsNotExist(err) {
		// jika file tidak ada maka:
		getDataMenu(&tempFilePath)
	} else if err != nil {
		// jika ada error yang lain
		panic(err)
	} else {
		// jika file temp sudah ada
		modTimeFile := fileTemp.ModTime()
		currentTime := time.Now()
		duration := currentTime.Sub(modTimeFile)
		targetDuration := 5 * time.Second

		if duration >= targetDuration {
			// jika waktu sudah melebihi 5 detik maka:
			getDataMenu(&tempFilePath)
		} else {
			// membaca isi file temp
			fileData, err := os.ReadFile(tempFilePath)
			if err != nil {
				panic(err)
			}

			// unmarshall isi file temp ke dalam bentuk slice of struct
			err = json.Unmarshal(fileData, &models.Menus)
			if err != nil {
				panic(err)
			}
		}
	}
}

func SelectMenu(reader *bufio.Reader, scanner *bufio.Scanner, w *tabwriter.Writer) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\x1bc")
			fmt.Println(r)
			fmt.Print("Press Enter to go back...")
			scanner.Scan()
		}
	}()
	cachingDataMenu()
	loop := true
	for loop {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Print(r)
					scanner.Scan()
				}
			}()
			fmt.Printf("\x1bc")
			fmt.Print("----------------- Select Menu -----------------\n\n")

			fmt.Println("-----------------------------------------------")
			fmt.Fprintln(w, "No\tName\tPrice")
			fmt.Fprintln(w, "---\t----------------------------\t------------")

			models.PrintRows(&models.Menu{}, w)

			w.Flush()

			fmt.Print("-----------------------------------------------\n")

			fmt.Print("\n0. Back to Home\n")
			fmt.Print("\nChoose a menu: ")

			choice, err := utils.InputInt(reader)

			if err != nil {
				panic("Invalid input, please enter a number... ")
			}

			if choice == 0 {
				loop = false
				return
			}

			found := false
			for _, item := range models.Menus {
				if item.ID == choice {
					fmt.Println("\nYou selected:", item.Name)
					found = true
					itemExists := false
					for i := range models.Carts {
						if models.Carts[i].ID == item.ID {
							models.Carts[i].Quantity++
							itemExists = true
							break
						}
					}

					if !itemExists {
						models.Carts = append(models.Carts, models.Cart{
							ID:       item.ID,
							Name:     item.Name,
							Quantity: 1,
							Price:    item.Price,
						})
					}
					fmt.Printf("%s has been added to your cart.\n", item.Name)
				}
			}

			if !found {
				panic("Menu not found, please try again... ")
			}

			fmt.Print("\nPress Enter to continue... ")
			scanner.Scan()
		}()
	}
}

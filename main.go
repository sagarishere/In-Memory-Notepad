package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Capacity int
var Notepad = make([]string, 0, 5)

func processInput(input string) {
	words := strings.Split(strings.TrimSpace(input), " ")
	if len(words) == 0 {
		return
	}
	switch words[0] {
	case "create":
		if len(words) == 1 {
			fmt.Println("[ERROR] Missing note argument")
		} else if len(Notepad) < Capacity {
			Notepad = append(Notepad, strings.Join(words[1:], " "))
			fmt.Println("[OK] The note was successfully created")
		} else {
			fmt.Println("[ERROR] Notepad is full")
		}
	case "list":
		if len(Notepad) == 0 {
			fmt.Println("[Info] Notepad is empty")
		} else {
			for i, note := range Notepad {
				fmt.Printf("[Info] %d: %s\n", i+1, note)
			}
		}
	case "update":
		if len(words) == 1 {
			fmt.Println("[Error] Missing position argument")
		} else if len(words) == 2 {
			fmt.Println("[Error] Missing note argument")
		} else if len(Notepad) == 0 {
			fmt.Println("[ERROR] Notepad is empty")
		} else {
			noteIndex, err := strconv.Atoi(words[1])
			if err != nil {
				fmt.Printf("[ERROR] Invalid position: %s\n", words[1])
			} else if noteIndex > len(Notepad) {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", noteIndex, len(Notepad))
			} else {
				Notepad[noteIndex-1] = strings.Join(words[2:], " ")
				fmt.Printf("[OK] The note at position %d was successfully updated\n", noteIndex)
			}
		}
	case "delete":
		if len(words) == 1 {
			fmt.Println("[Error] Missing position argument")
		} else if len(Notepad) == 0 {
			fmt.Println("[ERROR] Notepad is empty")
		} else {
			noteIndex, err := strconv.Atoi(words[1])
			if err != nil {
				fmt.Printf("[ERROR] Invalid position: %s\n", words[1])
			} else if noteIndex > len(Notepad) {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", noteIndex, len(Notepad))
			} else {
				Notepad = append(Notepad[:noteIndex-1], Notepad[noteIndex:]...)
				fmt.Printf("[OK] The note at position %d was successfully deleted\n", noteIndex)
			}
		}
	case "clear":
		Notepad = make([]string, 0, 5)
		fmt.Println("[OK] All notes were successfully deleted")
	case "exit":
		fmt.Println("[Info] Bye!")
		os.Exit(0)
	default:
		fmt.Println("[Error] Unknown command")
	}
}

func main() {
	fmt.Printf("Enter the maximum number of notes:")

	for {
		if Capacity != 0 {
			fmt.Printf("Enter command and data:")
		}

		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			input := scanner.Text()
			if Capacity == 0 {
				Capacity, _ = strconv.Atoi(input)
				continue
			}
			processInput(input)
		}
	}
}

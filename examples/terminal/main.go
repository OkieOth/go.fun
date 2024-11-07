package main

import (
	"fmt"
	"time"

	"github.com/k0kubun/go-ansi"
)

func printInOneLine() {
	lines := []string{
		"Line 1: Hello World!",
		"Line 2: Welcome to Go",
		"Line 3: Learning is fun",
		"Line 4: Keep coding!",
		"Line 5: Halfway there...",
		"Line 6: Go is awesome!",
		"Line 7: Almost done...",
		"Line 8: Just two more...",
		"Line 9: Almost finished!",
		"Line 10: Completed!",
	}

	for _, line := range lines {
		fmt.Printf("\r%s\r", "                          ") // Clear the line by overwriting with spaces
		fmt.Printf("\r%s", line)                           // Print the new line text
		time.Sleep(1 * time.Second)                        // Wait for 1 second before updating again
	}
	fmt.Println() // Move to the next line after the loop finishes
}

func printInTwoLines() {
	lines := []string{
		"Line 1: Hello World!",
		"Line 2: Welcome to Go",
		"Line 3: Learning is fun",
		"Line 4: Keep coding!",
		"Line 5: Halfway there...",
		"Line 6: Go is awesome!",
		"Line 7: Almost done...",
		"Line 8: Just two more...",
		"Line 9: Almost finished!",
		"Line 10: Completed!",
	}

	for i, line := range lines {
		if i%2 == 0 {
			// Odd lines (0-based index, so i%2 == 0 is actually line 1, 3, 5, ...)
			fmt.Printf("\033[1A\033[K\r%s", line) // Move cursor up, clear line, print new content
		} else {
			// Even lines
			fmt.Printf("\033[1B\033[K\r%s", line) // Move cursor down, clear line, print new content
		}
		time.Sleep(1 * time.Second) // Wait for 1 second before updating
	}
}

func printInTwoLinesWindowsReady() {
	lines := []string{
		"Line 1: Hello World!",
		"Line 2: Welcome to Go",
		"Line 3: Learning is fun",
		"Line 4: Keep coding!",
		"Line 5: Halfway there...",
		"Line 6: Go is awesome!",
		"Line 7: Almost done...",
		"Line 8: Just two more...",
		"Line 9: Almost finished!",
		"Line 10: Completed!",
	}
	for i, line := range lines {
		if i%2 == 0 {
			// Odd lines (0-based index, so i%2 == 0 is actually line 1, 3, 5, ...)
			ansi.Printf("\033[1A\033[K\r%s", line) // Move cursor up, clear line, print new content
		} else {
			// Even lines
			ansi.Printf("\033[1B\033[K\r%s", line) // Move cursor down, clear line, print new content
		}
		time.Sleep(1 * time.Second) // Wait for 1 second before updating
	}

}

func main() {
	printInTwoLinesWindowsReady()
}

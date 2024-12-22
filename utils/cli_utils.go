package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eiannone/keyboard"
)

func SelectItems(items []string, message string) ([]string, error) {
	if message != "" {
		fmt.Println(message)
	}

	selected := make(map[int]bool) // Track selected items
	currentIndex := 0              // Current highlighted index

	// Channel to listen for interrupt signals
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt, syscall.SIGTERM)

	// Open the keyboard for capturing input
	if err := keyboard.Open(); err != nil {
		return nil, fmt.Errorf("error opening keyboard: %w", err)
	}
	defer keyboard.Close()

	go func() {
		<-exitSignal // Wait for Ctrl+C
		keyboard.Close()
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	fmt.Println("Use ↑/↓ to navigate, Space to select, Enter to confirm, ESC to cancel.")

	for {
		// Clear the screen
		fmt.Print("\033[H\033[2J")
		fmt.Println("Use ↑/↓ to navigate, Space to select, Enter to confirm, ESC to cancel.")
		fmt.Println()

		// Display the menu
		for i, item := range items {
			if i == currentIndex {
				// Highlight the current row (e.g., inverted colors or background color)
				if selected[i] {
					fmt.Printf("\033[7m[*] %s\033[0m\n", item) // \033[7m = reverse video mode, \033[0m = reset formatting
				} else {
					fmt.Printf("\033[7m   %s\033[0m\n", item) // \033[7m = reverse video mode, \033[0m = reset formatting
				}
			} else if selected[i] {
				// Mark selected items
				fmt.Printf("[*] %s\n", item)
			} else {
				// Default formatting
				fmt.Printf("   %s\n", item)
			}
		}

		// Capture key press
		_, key, err := keyboard.GetKey()
		if err != nil {
			return nil, fmt.Errorf("error reading key: %w", err)
		}

		switch key {
		case keyboard.KeyArrowUp:
			if currentIndex > 0 {
				currentIndex--
			}
		case keyboard.KeyArrowDown:
			if currentIndex < len(items)-1 {
				currentIndex++
			}
		case keyboard.KeySpace:
			// Toggle selection
			selected[currentIndex] = !selected[currentIndex]
		case keyboard.KeyEnter:
			// Collect selected items and return
			var result []string
			for i, item := range items {
				if selected[i] {
					result = append(result, item)
				}
			}
			return result, nil
		case keyboard.KeyEsc:
			// Exit on ESC
			fmt.Println("\nOperation canceled.")
			return nil, nil
		}
	}
}

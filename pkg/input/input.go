// Made for isolate input logic for UI from UI itself
package input

import (
	"strings"
	"strconv"
	"bufio"
	"fmt"
	"os"
)

// GetInt: return int from Stdin input
func GetInt() (int, error) {
	// Creating new buffer-reader
	reader := bufio.NewReader(os.Stdin)

	// Reading text from Stdin
	input, err := reader.ReadString('\n')
	if err != nil { 
		return 0, fmt.Errorf("error while reading number: %v", err)
	}

	// Preparing and converting into int
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil { 
		return 0, fmt.Errorf("error while reading number: %v", err)
	}

	return number, nil
}

// GetFloat: return float from Stdin input
func GetFloat() (float64, error) {
	// Creating new buffer-reader
	reader := bufio.NewReader(os.Stdin)

	// Reading text from Stdin
	input, err := reader.ReadString('\n')
	if err != nil { 
		return 0, fmt.Errorf("error while reading number: %v", err)
	}

	// Preparing and converting into float64
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil { 
		return 0, fmt.Errorf("error while reading number: %v", err)
	}

	return number, nil
}

// GetString: return string from Stdin input
func GetString() (string, error) {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print(">> ")
    input, err := reader.ReadString('\n')
    if err != nil {
        return "", fmt.Errorf("error while reading string: %v", err)
    }
    
    return strings.TrimSpace(input), nil
}

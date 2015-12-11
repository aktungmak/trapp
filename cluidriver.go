package trapp

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// implements the UiDriver interface to provide command line ui
type ClUiDriver struct {
	Last   string
	Reader *bufio.Reader
}

func NewClUiDriver() *ClUiDriver {
	return &ClUiDriver{
		Reader: bufio.NewReader(os.Stdin),
	}
}

// read a single character from the input without waiting for enter
func (d *ClUiDriver) Prompt(prompt string) string {
	fmt.Print(prompt)
	text, err := d.Reader.ReadString('\n')
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return ""
	}
	text = strings.TrimSpace(text)

	// check if special
	switch text {
	case "up":
		return UP
	case "home":
		return HOME
	default:
		return text
	}
}

func (d *ClUiDriver) DisplayPath(path []string) {
	fmt.Println(strings.Join(path, " > "))
}

func (d *ClUiDriver) DisplayContent(content string) {
	fmt.Printf("%s\n\n", content)
}

func (d *ClUiDriver) ClearContent() {
	// doesn't apply to this type of ui
}

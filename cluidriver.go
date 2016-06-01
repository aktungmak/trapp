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

// constructor
func NewClUiDriver() *ClUiDriver {
	return &ClUiDriver{
		Reader: bufio.NewReader(os.Stdin),
	}
}

//// UiDriver methods ////

func (d *ClUiDriver) Prompt(prompt string) string {
	fmt.Print(prompt)
	text, err := d.Reader.ReadString('\n')
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return ""
	}
	text = strings.TrimSpace(text)

	if len(text) == 0 {
		text = d.Last
	} else {
		d.Last = text
	}

	// check if special
	switch text {
	case "up":
		return UP
	case "home":
		return HOME
	case "quit":
		return QUIT
	default:
		return text
	}
}

func (d *ClUiDriver) DisplayOpts(opts map[string]string) {
    fmt.Print("[ ")
    for opt, name := range opts {
        fmt.Printf("%s:%s ", opt, name)
    }
    fmt.Print("]")
}

func (d *ClUiDriver) DisplayPath(path []string) {
	fmt.Println("")
	fmt.Println(strings.Join(path, " > "))
}

func (d *ClUiDriver) DisplayContent(content string) {
	fmt.Printf("%s\n\n", content)
}

// doesn't apply to this type of ui
func (d *ClUiDriver) ClearContent() {
}

// doesn't apply to this type of ui
func (d *ClUiDriver) CleanUp() {
}

//// end UiDriver methods ////

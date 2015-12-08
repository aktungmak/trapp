package trapp

import (
	"fmt"
)

type UiDriver interface {
	promptUnbuffered(string) string
	promptBuffered(string) string
	displayPath([]string)
	displayContent(string)
	clearContent()
}

type clUiDriver struct {
}

// read a single character from the input without waiting for enter
func (d *clUiDriver) promptUnbuffered(prompt string) string {
	var c string
	fmt.Scan(&c)
	return c
}

func (d *clUiDriver) promptBuffered(prompt string) string {

}

func (d *clUiDriver) displayPath(path []string) {

}

func (d *clUiDriver) displayContent(content string) {
	fmt
}

func (d *clUiDriver) clearContent() {

}

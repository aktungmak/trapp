package trapp

import (
	"github.com/nsf/termbox-go"
	"strings"
)

// this driver uses terbox to display the ui
type TbUiDriver struct {
	Last string
}

func NewTbUiDriver() *TbUiDriver {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlue)
	return &TbUiDriver{}
}

func (d *TbUiDriver) Prompt(prompt string) string {
	return "s"
}
func (d *TbUiDriver) DisplayPath(path []string) {
	printLine(strings.Join(path, " > "), 0, 10)
}
func (d *TbUiDriver) DisplayContent(content string) {
}
func (d *TbUiDriver) ClearContent() {
}

func printLine(str string, x int, y int) {
	for i := range str {
		termbox.SetCell(x+i, y, rune(str[i]), termbox.ColorDefault, termbox.ColorDefault)
	}
}

package trapp

import (
	"fmt"
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

//// UiDriver methods ////

func (d *TbUiDriver) Prompt(prompt string) string {
	// handle resize here also
	rbuf := make([]rune, 0)
promptloop:
	for {
		printLine(string(rbuf), 0, 12)
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				rbuf = rbuf[:0]
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				rbuf = rbuf[:len(rbuf)-1]
			case termbox.KeyTab:
				// implement autocomplete??
			case termbox.KeySpace:
				rbuf = append(rbuf, ' ')
			case termbox.KeyEnter:
				break promptloop
			default:
				if ev.Ch != 0 {
					rbuf = append(rbuf, ev.Ch)
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
	return string(rbuf)
}

func (d *TbUiDriver) DisplayOpts(opts map[string]string) {
	printLine(fmt.Sprintf("%v", opts), 0, 11)
}

func (d *TbUiDriver) DisplayPath(path []string) {
	printLine(strings.Join(path, " > "), 0, 0)
}

func (d *TbUiDriver) DisplayContent(content string) {
}

func (d *TbUiDriver) ClearContent() {
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlue)
}

func (d *TbUiDriver) CleanUp() {
	termbox.Close()
}

//// end UiDriver methods ////

func printLine(msg string, x, y int) {
	for i, c := range msg {
		termbox.SetCell(x+i, y, c, termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.Flush()
}

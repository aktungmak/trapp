package trapp

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"strings"
)

const (
	TB_BG = termbox.ColorDefault
	TB_FG = termbox.ColorDefault
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
	termbox.Clear(TB_FG, TB_BG)
	return &TbUiDriver{}
}

//// UiDriver methods ////

func (d *TbUiDriver) Prompt(prompt string) string {
	// handle resize here also
	rbuf := make([]rune, 0)
promptloop:
	for {
		termbox.Flush()
		_, h := termbox.Size()
		printLine(string(rbuf), 0, h-1)
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
		case termbox.EventResize:
			d.Redraw()
		}
	}

	text := string(rbuf)
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

func (d *TbUiDriver) DisplayOpts(opts map[string]string) {
	_, h := termbox.Size()
	printLine(fmt.Sprintf("%v", opts), 0, h-2)
}

func (d *TbUiDriver) DisplayPath(path []string) {
	printLine(strings.Join(path, " > "), 0, 0)
}

func (d *TbUiDriver) DisplayContent(content string) {
	printLine(content, 2, 15)
	lines := strings.Split(content, "\n")
	_, h := termbox.Size()
	for i, line := range lines {
		// todo handle /r/n
		printLine(line, 0, i+10)
		if i+1 > h-3 {
			break
		}
	}
}

func (d *TbUiDriver) ClearContent() {
	termbox.Clear(TB_FG, TB_BG)
}

func (d *TbUiDriver) CleanUp() {
	termbox.Close()
	fmt.Println("cleaned up")
}

//// end UiDriver methods ////

func (d *TbUiDriver) Redraw() {
	termbox.Clear(TB_FG, TB_BG)

}

func printLine(msg string, x, y int) {
	for i, c := range msg {
		termbox.SetCell(x+i, y, c, TB_FG, TB_BG)
	}
	termbox.Flush()
}

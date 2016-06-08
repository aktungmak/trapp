package trapp

import (
	"github.com/nsf/termbox-go"
	"strings"
)

const (
	TB_BG = termbox.ColorBlue
	TB_FG = termbox.ColorWhite
    WRAP = true
)

// this driver uses terbox to display the ui
type TbUiDriver struct {
	Last string
	w    int
	h    int
}

// constructor
func NewTbUiDriver() *TbUiDriver {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.Clear(TB_FG, TB_BG)
	t := &TbUiDriver{}
	t.w, t.h = termbox.Size()
	return t
}

//// UiDriver methods ////

func (d *TbUiDriver) Prompt(prompt string) string {
	// handle resize here also
	rbuf := make([]rune, 0)
promptloop:
	for {
		d.printLine(prompt+string(rbuf), 0, d.h-1)
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
			d.w, d.h = termbox.Size()
			termbox.Flush()
		}
	}

	text := string(rbuf)

	if len(text) == 0 {
		text = d.Last
	} else {
		d.Last = text
	}

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
	s := make([]string, 0)
	for k, v := range opts {
		s = append(s, k+": "+v)
	}
	d.printLineInvert(strings.Join(s, " "), 0, d.h-2)
}

func (d *TbUiDriver) DisplayPath(path []string) {
	d.printLineInvert(strings.Join(path, " > "), 0, 0)
}

func (d *TbUiDriver) DisplayContent(content string) {
    content = strings.Replace(content, "\r", "", -1)
	lines := strings.Split(content, "\n")

    // TODO wrap lines that are too long

	for i, line := range lines {
		d.printLine(line, 0, i+1)
        // trim off content that won't fit
		if i+1 > d.h-3 {
			break
		}
	}
}

func (d *TbUiDriver) ClearContent() {
	termbox.Clear(TB_FG, TB_BG)
    termbox.Flush()
}

func (d *TbUiDriver) CleanUp() {
	termbox.Close()
}

//// end UiDriver methods ////

func (d *TbUiDriver) clearLine(x, y int) {
	for i := x; i < d.w; i++ {
		termbox.SetCell(i, y, ' ', TB_FG, TB_BG)
	}
	termbox.Flush()
}
func (d *TbUiDriver) clearLineInvert(x, y int) {
	for i := x; i < d.w; i++ {
		termbox.SetCell(i, y, ' ', TB_BG, TB_FG)
	}
	termbox.Flush()
}

func (d *TbUiDriver) printLine(msg string, x, y int) {
	d.clearLine(x, y)
	for i, c := range msg {
		termbox.SetCell(x+i, y, c, TB_FG, TB_BG)
	}
	termbox.Flush()
}

func (d *TbUiDriver) printLineInvert(msg string, x, y int) {
	d.clearLineInvert(x, y)
	for i, c := range msg {
		termbox.SetCell(x+i, y, c, TB_BG, TB_FG)
	}
	termbox.Flush()
}

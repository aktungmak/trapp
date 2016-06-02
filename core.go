package trapp

import (
	"errors"
)

const (
	UP   string = "_TRAPP_GO_UP_"
	HOME string = "_TRAPP_GO_HOME_"
	QUIT string = "_TRAPP_QUIT_"
)

// a mapping of strings to nodes, part of the
// tree's basic structure
type OptMap map[string]*Node

// represents a single node of the tree
// can have children nodes, or a function or both
type Node struct {
	Name string
	Func func(AppState)string
	Opts OptMap

	parent *Node
}

// tree application core.
type Trapp struct {
	// the base of the tree
	Root *Node
	// our current location
	Current *Node
	// our link to the outside world
	Ui UiDriver
	// current-continuation, represents the app data
	AppState AppState
}

func NewTrapp(tree *Node, ui UiDriver, as AppState) *Trapp {
	t := &Trapp{
		Root:    tree,
		Current: tree,
		Ui:      ui,
		AppState:      as,
	}

	return t
}

func (t *Trapp) Select(opt string) (string, error) {
	// first check for special
	switch opt {
	case UP:
		t.Up()
		return "", nil
	case HOME:
		t.Home()
		return "", nil
	case QUIT:
		return "", errors.New(QUIT)
	}

	// now check for regular options
	next, ok := t.Current.Opts[opt]
	if !ok {
		return "", errors.New("not a valid option")
	}

	// execute the func if specified and get its output
    var output string
	if next.Func != nil {
		output = next.Func(t.AppState)
	}

	// if it has options, change to that
	if len(next.Opts) > 0 {
		// set the parent field of the child so we can go back
		next.parent = t.Current
		t.Current = next
	}
	return output, nil
}

// go up one level, setting t.Current to its parent
func (t *Trapp) Up() *Node {
	if t.Current.parent != nil {
		t.Current = t.Current.parent
		return t.Current
	} else {
		return nil
	}
}

//jump to the root of the tree
func (t *Trapp) Home() {
	//keep calling Up() until it returns nil
	for t.Up() != nil {
	}
}

// the main loop of the application, that processes each stage
func (t *Trapp) EventLoop() {
	defer t.Ui.CleanUp()
	for {

		// print current position in tree
		t.Ui.DisplayPath(t.GetCurrentPath())

		//show the available options
		opts := make(map[string]string)
		for k, v := range t.Current.Opts {
			opts[k] = v.Name
		}
		t.Ui.DisplayOpts(opts)

		// wait for input
		opt := t.Ui.Prompt(": ")

		t.Ui.ClearContent()
		// if we got something, try selecting
		if len(opt) > 0 {
			output, err := t.Select(opt)
			if err != nil {
				if err.Error() == QUIT {
					break
				} else {
					t.Ui.DisplayContent(err.Error())
				}
			} else {
                t.Ui.DisplayContent(output)
            }
		}
	}
}

// get a string slice showing the current position in the tree
func (t *Trapp) GetCurrentPath() []string {
	p := make([]string, 0)
	n := t.Current
	for n != nil {
		// note that this is prepending!
		p = append([]string{n.Name}, p...)
		n = n.parent
	}
	return p
}

type UiDriver interface {
	Prompt(string) string
	DisplayPath([]string)
	DisplayOpts(map[string]string)
	DisplayContent(string)
	ClearContent()
	CleanUp()
}

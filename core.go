package trapp

import (
	"errors"
)

const (
	UP   string = "_TRAPP_GO_UP_"
	HOME string = "_TRAPP_GO_HOME_"
)

// a mapping of strings to nodes, part of the
// tree's basic structure
type OptMap map[string]*Node

// represents a single node of the tree
// can have children nodes, or a function
// or both
type Node struct {
	Name string
	Func func(*Trapp)
	Opts OptMap

	parent *Node
}

func NewNode(name string, f func(*Trapp), opts OptMap) *Node {
	return &Node{
		Name: name,
		Func: f,
		Opts: opts,
	}
}
func NewNodeBlank() *Node {
	return &Node{
		Name: "",
		Func: func(*Trapp) {},
		Opts: make(OptMap),
	}
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
	Cc interface{}
}

func NewTrapp(tree *Node, ui UiDriver, cc interface{}) *Trapp {
	t := &Trapp{
		Root:    tree,
		Current: tree,
		Ui:      ui,
		Cc:      cc,
	}

	return t
}

func (t *Trapp) Select(opt string) error {
	// first check for special
	if opt == UP {
		t.Up()
		return nil
	} else if opt == HOME {
		t.Home()
		return nil
	}

	// now check for regular options
	next, ok := t.Current.Opts[opt]
	if !ok {
		return errors.New("not a valid option")
	}

	// execute the func
	next.Func(t)

	// if it has options, change to that
	if len(next.Opts) > 0 {
		// set the parent field of the child so we can go back
		next.parent = t.Current
		t.Current = next
	}
	return nil
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
	for {
		// print current position in tree
		t.Ui.DisplayPath(t.GetCurrentPath())

		// print options

		// wait for input
		opt := t.Ui.Prompt(":")

		// if we got something, try selecting
		if len(opt) > 0 {
			err := t.Select(opt)
			if err != nil {
				t.Ui.DisplayContent(err.Error())
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
	DisplayContent(string)
	ClearContent()
}

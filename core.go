package trapp

import (
	"errors"
)

// a mapping of strings to nodes, part of the
// tree's basic structure
type OptMap map[string]*Node

// represents a single node of the tree
// can have children nodes, or a function
// or both
type Node struct {
	Name string
	Func func(interface{})
	Opts OptMap

	parent *Node
}

func NewNode(name string, f func(interface{}), opts OptMap) *Node {
	return &Node{
		Name: name,
		Func: f,
		Opts: opts,
	}
}
func NewNodeBlank() *Node {
	return &Node{
		Name: "",
		Func: func(interface{}) {},
		Opts: make(OptMap),
	}
}

// tree application core.
type Trapp struct {
	Root    *Node       // the base of the tree
	Current *Node       // our current location
	Ui      UiDriver    // our link to the outside world
	Cc      interface{} // current-continuation, represents the app data
	// gets passed to every action
}

func NewTrapp(tree *Node, ui UiDriver, cc interface{}) *Trapp {
	t := &Trapp{}

	t.Root = tree
	t.Current = tree
	t.Ui = ui
	t.Cc = cc

	return t
}

func (t *Trapp) Select(opt string) error {
	next, ok := t.Current.Opts[opt]
	if !ok {
		return errors.New("not a valid option")
	}
	// execute the func with t.Cc
	next.Func(t.Cc)
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
		p = append(p, n.Name)
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

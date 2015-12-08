package trapp

import (
	"errors"
	"log"
)

// represents a single node of the tree
// can have children nodes, or a function
// or both
type Node struct {
	Name string
	Func func(interface{})
	Opts map[string]*Node

	parent *Node
}

func NewNode(name string, f func(interface{}), opts map[string]*Node) *Node {
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
		Opts: make(map[string]*Node),
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

func NewTrapp(tree *Node, ui UiDriver, cc interface{}) (*Trapp, error) {
	t := &Trapp{}

	err := validateTree(tree)
	if err != nil {
		return t, err
	}
	t.Root = tree
	t.Current = tree
	t.Ui = ui
	t.Cc = cc
}

func (t *Trapp) Select(opt string) error {
	next, ok := t.Current[opt]
	if !ok {
		return Errors.New("not a valid option")
	}
	// if has children, change to that
	// set the parent field of the child so we can go back
	// else has func, execute func with t.Cc
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
}

// the main loop of the application, that processes each stage
func (t *Trapp) EventLoop() {
	for {
		// print current position in tree

		// print options

		// wait for input

		// if special, do that
		// else check if it exists

		// update last
	}
}

// get a string slice showing the current position in the tree
func (t *Trapp) GetCurrentPath() []string {

}

// walk the supplied tree and ensure it is well-formed
func ValidateTree(tree map[string]*Node) error {
	// thanks to types, this is actually not needed I think!

}

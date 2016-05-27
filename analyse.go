package trapp

import (
    "reflect"
    "errors"
    "encoding/json"
)

// Cc is a struct with a number of custom methods
// which specify what the application can do.
// The fields of the struct are the variables 
// defining the application's state.
type Cc interface{}

// represents a node in the config, where all fields are string
type CfgNode struct {
    Name string
    Func string
    Opts map[string]CfgNode
}
func NewNodeFromCfgNode(cn CfgNode, parent *Node, cct reflect.Type) (*Node, error) {
	n := &Node{}
	n.Name = cn.Name
    n.parent = parent
	
    // for the func field, look up the method on cc
    // if it is not there, bail out with an error
    meth, ok := cct.MethodByName(cn.Func)
    if !ok {
        return nil, errors.Errorf("method %s used in config but not defined", fname)
    }

    // now iterate through the opts, making a new Node for each
    n.Opts = make(map[string]Node)
	for k, v := range cn.Opts {
		n.Opts[k], err := NewNodeFromCfgNode(v, n, cct)
        if err != nil {
            return nil, err
        }
	}
    
	return n, nil
}

// given a json string, parse it into a tree structure
// of nodes, representing the structure of the application.
// the structure cc represents the state of the app, 
// and should have methods with the same name as the 
// "func" fields in the json.
func ProcessJsonConfig(json string, cc Cc) (*Node, error) {
    // parse the json as a tree of CfgNodes first
    cn := CfgNode{}
	err := json.Unmarshal([]byte(data), &cn)
    if err != nil {
        return nil, err
    }

    // get the reflect representation of cc's type
    cct := reflect.TypeOf(cc)

    // recursively create the real Node from CfgNode    
    n, err := NewNodeFromCfgNode(cn, nil, cct)
    if err != nil {
        return nil, err
    } else {
        return n, nil
}

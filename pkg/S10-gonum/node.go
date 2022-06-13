package gonum_examples

import (
	"fmt"

	"gonum.org/v1/gonum/graph"
)

type node struct {
	id   int64
	name string
	graph.Node
	dotID string
	dims
	left, right *node
}

func (n node) ID() int64           { return n.id }
func (n node) negated() node       { return node{-n.id, n.name} }
func (n *node) SetDOTID(id string) { n.name = id }
func (n *node) String() string {
	if n.left == nil || n.right == nil {
		rows, cols := n.shape()
		return fmt.Sprintf("[%d×%d]", rows, cols)
	}
	rows, cols := n.shape()
	return fmt.Sprintf("(%s * %s):[%d×%d]", n.left, n.right, rows, cols)
}

// SetDOTID sets the DOT ID of the node.
func (n *node) SetDOTID(id string) { n.dotID = id }

func (n *node) String() string { return n.dotID }

// Package basenode implements the BaseNode type and methods.
package basenode

import (
	"github.com/stvmln86/galah/galah/elems/icon"
	"github.com/stvmln86/galah/galah/elems/node"
)

// BaseNode is a basic implementation of a Node.
type BaseNode struct {
	icon icon.Icon
	name string
	open bool
}

// New returns a new BaseNode.
func New(icon icon.Icon, name string, open bool) *BaseNode {
	return &BaseNode{icon, name, open}
}

// NewFunc returns a new BaseNode generator function.
func NewFunc(icon icon.Icon, name string, open bool) node.Function {
	return func() node.Node {
		return New(icon, name, open)
	}
}

// Icon returns the Node's Icon.
func (n *BaseNode) Icon() icon.Icon {
	return n.icon
}

// Name returns the Node's name.
func (n *BaseNode) Name() string {
	return n.name
}

// Open returns true if the Node can be moved through.
func (n *BaseNode) Open() bool {
	return n.open
}

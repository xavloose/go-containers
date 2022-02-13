package list

import "reflect"

type node struct {
	n *node
	p *node

	t reflect.Type
	v interface{}
}

func newnode(v interface{}) *node {
	return &node{t: reflect.TypeOf(v), v: v}
}

func (n *node) Type() reflect.Type {
	return n.t
}

func (n *node) Val() interface{} {
	return n.v
}

package list

import (
	"reflect"
	"sync"
)

// List is a superlist that combines the functionality of all the other types of lists (singly,doubly,circularly)
type List = list

// NewList will create a superlist that supports the given type
func NewList(t reflect.Type) *List {
	return newlist(t)
}

// list is meant to provide a generic universal object that other things can use
// While the functionality of all list could be combined for the purposes of this
// the intention of the different implementations is to demonstrate their purposes
// This module is meant to be more educational as opposed to being something pratically
// used except maybe as basis for developing templates
type list struct {
	h *node // marks the first node in the list
	e *node // marks the last node in the list
	c *node // cursor for iteration

	t reflect.Type // data type

	size int // It is up to the implementation to take care of this

	// we offer lock so that implementations can take advantage of
	// creating thread safe Cursors to iterate through lists
	lock sync.RWMutex
}

func newlist(t reflect.Type) *list {
	return &list{t: t}
}

func (l *list) listType() reflect.Type {
	return l.t
}

// nodeAt will traverse the list from the beginning to find the node at
// the given index. It will return a nil node if there is nothing
func (l *list) nodeAt(idx int) *node {
	// Node to return
	n := l.h

	for i := 1; i < idx && n != nil; i++ {
		n = n.n
	}

	return n
}

// checkType is a simple function that will check whether or not a given value is
// the correct type for the list
func (l *list) checkType(v interface{}) bool {
	return reflect.TypeOf(v).ConvertibleTo(l.t)
}

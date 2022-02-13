package list

import "reflect"

// SinglyLinkedList
type SinglyLinkedList = list

// NewSinglyLinkedList will create a new container singly linked list that
// supports the given data type
func NewSinglyLinkedList(t reflect.Type) *SinglyLinkedList {
	return newlist(t)
}

// Type will return the data type supported by this list
func (sll *SinglyLinkedList) Type() reflect.Type {
	return sll.listType()
}

// Len will return the size of the object
func (sll *SinglyLinkedList) Len() int {
	return sll.size
}

// Clear will empty the list and reset everything
func (sll *SinglyLinkedList) Clear() {
	sll.h = nil
	sll.e = nil
	sll.c = nil
	sll.size = nil
}

func (sll *SinglyLinkedList) Add(v interface{}) error {
	// Check the type
	if !sll.checkType(v) {
		return ErrWrongDataType
	}

	// create the new node
	nn := newnode(v)

	// Case: empty
	if sll.h == nil {
		// The new node is the head of the list
		sll.h = nn
		// The new node is also the end of the list
		sll.e = nn
		// Set the cursor to the beginning of the list
		sll.c = nn
		return nil
	}

	// Append the element to the end of the list
	sll.e.n = nn
	sll.e = nn
	sll.size += 1

	return nil
}

func (sll *SinglyLinkedList) InsertAt(idx int, v interface{}) error {
	// Check the type
	if !sll.checkType(v) {
		return ErrWrongDataType
	}

	// create the new node
	nn := newnode(v)

	// We are inserting it at the beginning of the list
	if idx == 0 {
		nn.n = sll.h
		sll.h = nn
		sll.size += 1
		return nil
	}

	// Insertion at all other ends
	n := sll.nodeAt(idx - 1)
	if n == nil {
		return ErrOutOfBounds
	}

	// Set new nodes next to previous nodes next
	nn.n = n.n
	// Set previous nodes next to the new node
	n.n = nn
	sll.size += 1

	return nil
}

func (sll *SinglyLinkedList) SetIterator(idx int) error {
	if idx >= sll.size {
		sll.c = nil
		return ErrOutOfBounds
	}

	// Cursor Node
	n := sll.h
	defer func(){
		sll.c = n
	}()

	for i := 0; i < 
}

func (sll *SinglyLinkedList) At() interface{} {
	return nil
}

func (sll *SinglyLinkedList) Next() interface{} {
	return nil
}

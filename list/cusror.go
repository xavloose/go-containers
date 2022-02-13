package list

type cursorStyle int

type Cursor struct {
	l  *list
	n  *node
	fn *node
}

// newCursor creates a Cursor that lists can provide for list
// style
func newCursor(l *list, n *node) *Cursor {
	return &Cursor{l: l, n: n, fn: n}
}

// At will return the curro
func (c *Cursor) At() (interface{}, error) {
	// A nodes value cannot be changed after it
	// has been set so this is threadsafe
	if c.n != nil {
		return c.n.Val(), nil
	}
	c.fn = nil
	return nil, ErrCursorNotOnNode
}

// Next will
func (c *Cursor) Next() (interface{}, error) {
	c.l.lock.RLock()
	defer c.l.lock.RUnlock()
	if c.n != nil {
		return nil, ErrCursorNotOnNode
	}
	c.n = c.n.n
	return c.At()
}

func (c *Cursor) Prev() (interface{}, error) {

}

package dull

const (
	// Render the cursor as a line at the bottom of the cell.
	CursorTypeUnder CursorType = 0 + iota
	// Render the cursor as a block,
	// by inverting the cell's background and foreground colors.
	CursorTypeBlock CursorType = iota
)

type CursorType int

// Cursor defines a cursor within a cell.
type Cursor struct {
	// The column of the cell to show the cursor in.
	Column int
	// The row of the cell to show the cursor in.
	Row int
	// The color to use to draw the cursor.
	Color Color
	// How the cursor is to be rendered.
	Type CursorType
	// Whether the cursor is renderer or not.
	// Should always be set to true, unless it is used
	// to perform cursor flashing.
	Visible bool
}

// CursorId is an identifier provided when adding a cursor.
// It may later be used to remove a cursor.
type CursorId int

// Cursors represents a collection of cursors that a window may render.
//
// An instance is provided by a Window.
type Cursors struct {
	nextId  CursorId
	cursors map[CursorId]*Cursor
}

func newCursors() *Cursors {
	return &Cursors{
		nextId:  0,
		cursors: make(map[CursorId]*Cursor),
	}
}

// Add adds a cursor.
//
// The returned CursorId may be later used to remove the cursor.
func (c *Cursors) Add(cursor *Cursor) CursorId {
	c.nextId++
	id := c.nextId
	c.cursors[id] = cursor

	return id
}

// Removes a cursor.
//
// The cursor is identified by an id returned from the Add function.
func (c *Cursors) Remove(id CursorId) {
	delete(c.cursors, id)
}

// RemoveAll removes all cursors.
func (c *Cursors) RemoveAll() {
	c.cursors = make(map[CursorId]*Cursor)
}

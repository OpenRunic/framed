package framed

type Table struct {

	// internal state to mark data types are resolved
	resolved bool

	// slice of table rows
	Rows []*Row

	// current state of table
	State *State

	// current resolved options for table
	Options *Options
}

// New creates new [Table] instance with [OptionCallback]
func New(ocbs ...OptionCallback) *Table {
	return (&Table{
		resolved: false,
		Rows:     make([]*Row, 0),
		Options:  NewOptions(ocbs...),
		State: &State{
			Indexes:     make(IndexCache),
			Columns:     make([]string, 0),
			Definitions: make(map[string]*Definition, 0),
		},
	}).Initialize()
}

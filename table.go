package framed

type IndexCache = map[string]int

type Table struct {
	resolved bool
	Rows     []*Row
	State    *State
	Options  *Options
}

func New(ocbs ...OptionCallback) *Table {
	return (&Table{
		resolved: false,
		Rows:     make([]*Row, 0),
		Options:  NewOptions(ocbs...),
		State: &State{
			Indexes:     make(IndexCache),
			Columns:     make([]string, 0),
			Definitions: make(map[string]*ColumnDefinition, 0),
		},
	}).Initialize()
}

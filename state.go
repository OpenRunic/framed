package framed

import (
	"maps"
	"reflect"
	"slices"
)

// columns index cache type
type IndexCache = map[string]int

// current state of table
type State struct {

	// list of table columns
	Columns []string

	// indexes of detected columns
	Indexes IndexCache

	// resolved column definitions
	Definitions map[string]*Definition
}

// IsEmpty checks if columns are available
func (s *State) IsEmpty() bool {
	return len(s.Columns) < 1
}

// Index retrieves the index of column
func (s *State) Index(name string) int {
	idx, ok := s.Indexes[name]
	if ok {
		return idx
	}
	return -1
}

// HasColumn checks if column is available
func (s *State) HasColumn(name string) bool {
	return s.Index(name) > -1
}

// ColumnName retrieves the name of column from index
func (s *State) ColumnName(idx int) string {
	if s.IsEmpty() || idx >= len(s.Columns) {
		return ""
	}

	return s.Columns[idx]
}

// HasDefinition checks if definition is available
func (s *State) HasDefinition(name string) bool {
	_, ok := s.Definitions[name]
	return ok
}

// Definition retrieves the value definition
func (s *State) Definition(name string) *Definition {
	return s.Definitions[name]
}

// DefinitionAt retrieves the value definition via index
func (s *State) DefinitionAt(idx int) *Definition {
	if idx >= len(s.Columns) {
		return nil
	}

	return s.Definitions[s.Columns[idx]]
}

// DataTypes returns the saved data types in definitions
func (s *State) DataTypes() map[string]reflect.Type {
	return SliceKeyMap(s.Columns, func(col string, _ int) (string, reflect.Type) {
		return col, s.DataType(col)
	})
}

// DataType returns data type for single value
func (s *State) DataType(name string) reflect.Type {
	def := s.Definition(name)
	if def != nil {
		return def.Type
	}

	return nil
}

// Clone duplicates state to new instance
func (s *State) Clone() *State {
	return &State{
		Columns:     slices.Clone(s.Columns),
		Indexes:     maps.Clone(s.Indexes),
		Definitions: maps.Clone(s.Definitions),
	}
}

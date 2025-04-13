package framed

import (
	"maps"
	"reflect"
	"slices"
)

// current state of table
type State struct {

	// list of table columns
	Columns []string

	// indexes of detected columns
	Indexes IndexCache

	// resolved column definitions
	Definitions map[string]*ColumnDefinition
}

// check is columns exist
func (s *State) IsEmpty() bool {
	return len(s.Columns) < 1
}

// get the current index of column
func (s *State) Index(name string) int {
	idx, ok := s.Indexes[name]
	if ok {
		return idx
	}
	return -1
}

func (s *State) HasColumn(name string) bool {
	return s.Index(name) > -1
}

// get the name of column from index
func (s *State) ColumnName(idx int) string {
	if s.IsEmpty() || idx >= len(s.Columns) {
		return ""
	}

	return s.Columns[idx]
}

func (s *State) HasDefinition(name string) bool {
	_, ok := s.Definitions[name]
	return ok
}

func (s *State) Definition(name string) *ColumnDefinition {
	return s.Definitions[name]
}

func (s *State) DefinitionAt(idx int) *ColumnDefinition {
	if idx >= len(s.Columns) {
		return nil
	}

	return s.Definitions[s.Columns[idx]]
}

func (s *State) DataTypes() map[string]reflect.Type {
	return SliceKeyMap(s.Columns, func(col string, _ int) (string, reflect.Type) {
		return col, s.DataType(col)
	})
}

func (s *State) DataType(name string) reflect.Type {
	def := s.Definition(name)
	if def != nil {
		return def.Type
	}

	return nil
}

func (s *State) Clone() *State {
	return &State{
		Columns:     slices.Clone(s.Columns),
		Indexes:     maps.Clone(s.Indexes),
		Definitions: maps.Clone(s.Definitions),
	}
}

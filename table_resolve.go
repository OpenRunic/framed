package framed

import (
	"fmt"
	"reflect"
	"slices"
)

// Initialize executes the options on table instance
func (t *Table) Initialize() *Table {
	if t.Options.Columns != nil {
		t.UseColumns(t.Options.Columns)
	}

	for n, def := range t.Options.Definitions {
		_, ok := t.State.Definitions[n]
		if !ok {
			t.SetDefinition(n, def)
		}
	}

	return t
}

// SetState overrides the [State] of table
func (t *Table) SetState(s *State) *Table {
	t.State = s
	return t
}

// SetOptions overrides the [Options] of table
func (t *Table) SetOptions(opts *Options) *Table {
	t.Options = opts
	return t
}

// SetColumns updates the header columns for table
func (t *Table) SetColumns(cols []string) {
	t.State.Columns = cols
}

// SetIndexes updates the column indexes for table
func (t *Table) SetIndexes(cache IndexCache) {
	t.State.Indexes = cache
}

// SetDefinition assigns [Definition] for the column
func (t *Table) SetDefinition(name string, def *Definition) *Definition {
	t.State.Definitions[name] = def
	return def
}

// ResolveDefinition stores the column [Definition] if it doesn't exist
func (t *Table) ResolveDefinition(name string, tp reflect.Type) *Definition {
	def := t.State.Definition(name)
	if def != nil {
		return def
	}

	t.State.Definitions[name] = NewDefinition(tp)

	return t.State.Definition(name)
}

// ResolveValueDefinition detects data type of column value and creates [Definition]
func (t *Table) ResolveValueDefinition(idx int, name string, value string) *Definition {
	def := t.State.Definition(name)
	if def != nil {
		return def
	}

	tp := ToType("")
	dContinue := true
	if t.Options.TypeReader != nil {
		rTp := t.Options.TypeReader(idx, value)
		if rTp != nil {
			tp = rTp
			dContinue = false
		}
	}

	if dContinue {
		tp = DetectValueType(name, value)
	}

	t.State.Definitions[name] = NewDefinition(tp)

	return t.State.Definition(name)
}

// ResolveTypes resolves the data types from the column values
func (t *Table) ResolveTypes(names []string, values []string) error {
	if !t.resolved {
		t.resolved = true

		if len(names) != len(values) {
			return fmt.Errorf("invalid size of column names and values; %d != %d", len(names), len(values))
		}

		for idx, value := range values {
			t.ResolveValueDefinition(idx, names[idx], value)
		}
	}

	return nil
}

// UseColumns updates header columns for table
func (t *Table) UseColumns(values []string) {
	cache := make(IndexCache)
	for idx, col := range values {
		cache[col] = idx
	}

	t.SetColumns(slices.Clone(values))
	t.SetIndexes(cache)
}

// MarkUnresolved marks table as unresolved
func (t *Table) MarkUnresolved() *Table {
	t.resolved = false
	return t
}

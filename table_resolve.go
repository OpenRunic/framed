package framed

import (
	"fmt"
	"reflect"
	"slices"
)

// initialize the options on table instance
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

func (t *Table) SetState(s *State) *Table {
	t.State = s
	return t
}

func (t *Table) SetOptions(opts *Options) *Table {
	t.Options = opts
	return t
}

func (t *Table) SetColumns(cols []string) {
	t.State.Columns = cols
}

func (t *Table) SetIndexes(cache IndexCache) {
	t.State.Indexes = cache
}

func (t *Table) SetDefinitions(defs map[string]*ColumnDefinition) *Table {
	t.State.Definitions = defs
	return t
}

func (t *Table) SetDefinition(name string, def *ColumnDefinition) *ColumnDefinition {
	t.State.Definitions[name] = def
	return def
}

// resolve the type of column value if not defined
func (t *Table) ResolveDefinition(idx int, name string, tp reflect.Type) *ColumnDefinition {
	def := t.State.Definition(name)
	if def != nil {
		return def
	}

	t.State.Definitions[name] = NewDefinition(tp)

	return t.State.Definition(name)
}

// detect and resolve the type of column value
func (t *Table) ResolveValueDefinition(idx int, name string, value string) *ColumnDefinition {
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

// resolve the data types from the column values
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

// save and resolve columns from provided values
func (t *Table) UseColumns(values []string) {
	cache := make(IndexCache)
	for idx, col := range values {
		cache[col] = idx
	}

	t.SetColumns(slices.Clone(values))
	t.SetIndexes(cache)
}

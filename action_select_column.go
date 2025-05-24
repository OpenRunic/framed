package framed

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
)

// ColumnStatValue defines struct for key value
// stat for selected column
type ColumnStatValue struct {
	Key   any
	Value any
}

func (v ColumnStatValue) String() string {
	return fmt.Sprintf("%v, %v", v.Key, v.Value)
}

// ActionSelectConfig defines struct for column
// settings to be applied on action
type ActionSelectConfig struct {
	reverse  bool
	matches  []any
	resolver func(string) bool

	alias  string
	prefix string
	suffix string

	values   []any
	builder  func(*Table) ([]any, error)
	modifier func(*State, *Row) (any, error)
}

//

// Alias sets new name for column
func (c *ActionSelectConfig) Alias(n string) *ActionSelectConfig {
	c.alias = n
	return c
}

// Prefix allows to add prefix for column name
func (c *ActionSelectConfig) Prefix(s string) *ActionSelectConfig {
	c.prefix = s
	return c
}

// Suffix allows to add suffix  for column name
func (c *ActionSelectConfig) Suffix(s string) *ActionSelectConfig {
	c.suffix = s
	return c
}

// WithValues allows to assign raw values to column
func (c *ActionSelectConfig) WithValues(values []any) *ActionSelectConfig {
	c.values = values
	return c
}

// Modify allows you to create/update column value using callback
func (c *ActionSelectConfig) Modify(cb func(*State, *Row) (any, error)) *ActionSelectConfig {
	c.modifier = cb
	return c
}

// Build allows you to collect column values using callback
func (c *ActionSelectConfig) Build(cb func(*Table) ([]any, error)) *ActionSelectConfig {
	c.builder = cb
	return c
}

// ResolveName adds any prefix of suffix to column
func (c *ActionSelectConfig) ResolveName(column string) string {
	if len(c.alias) > 0 {
		return c.alias
	} else if len(c.prefix) > 0 {
		return fmt.Sprintf("%s%s", c.prefix, column)
	} else if len(c.suffix) > 0 {
		return fmt.Sprintf("%s%s", column, c.suffix)
	}

	return column
}

// ResolveColumns evaluates the config to get column names
func (c *ActionSelectConfig) ResolveColumns(columns []string) ([]string, error) {
	nColumns := make([]string, 0)

	if c.matches != nil {
		for _, m := range c.matches {
			switch t := m.(type) {
			case *regexp.Regexp:
				nColumns = append(nColumns, SliceFilter(columns, func(s string) bool {
					return t.MatchString(s)
				})...)
			case string:
				if slices.Contains(columns, t) || c.builder != nil ||
					c.values != nil || c.modifier != nil {
					nColumns = append(nColumns, t)
				}
			default:
				return nil, errors.New("only string and regexp.Regexp are supported for column selection")
			}
		}
	} else if c.resolver != nil {
		nColumns = SliceFilter(columns, c.resolver)
	}

	if c.reverse {
		return SliceFilter(columns, func(s string) bool {
			return !slices.Contains(nColumns, s)
		}), nil
	}

	return nColumns, nil
}

// Col creates new instance of column selection
// where accepted values for matches should be string or *[regexp.Regexp]
func Col(matches ...any) *ActionSelectConfig {
	return &ActionSelectConfig{matches: matches}
}

// ColAll allows you to pick all columns except defined ones
// where accepted values for matches should be string or *[regexp.Regexp]
func ColAll(matches ...any) *ActionSelectConfig {
	return &ActionSelectConfig{matches: matches, reverse: true}
}

// ColFunc allows you to pick from existing columns using func
func ColFunc(resolver func(string) bool) *ActionSelectConfig {
	return &ActionSelectConfig{resolver: resolver}
}

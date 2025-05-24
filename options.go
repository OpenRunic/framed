package framed

import (
	"maps"
	"reflect"
	"slices"
)

// Options defines settings for table data
type Options struct {

	// Name of table
	Name string

	// sample rows count
	SampleSize int

	// printable sample text length
	SampleMaxLength int

	// stop import at max defined count
	MaxRows int

	// ignore the header line
	IgnoreHeader bool

	// column separater character
	Separator byte

	// pre-defined columns for table
	Columns []string

	// pre-defined column definitions
	Definitions map[string]*Definition

	// helper to read column type
	TypeReader func(int, string) reflect.Type

	// text to display for empty values
	EmptyText string

	// defines characters to use for table pretty print
	PrettyFormatCharacters []rune
}

// OptionCallback defines function signature for option builder
type OptionCallback = func(*Options)

// Clone duplicates the options as new instance
func (o *Options) Clone() *Options {
	return &Options{
		Name:                   o.Name,
		MaxRows:                o.MaxRows,
		Separator:              o.Separator,
		SampleSize:             o.SampleSize,
		SampleMaxLength:        o.SampleMaxLength,
		IgnoreHeader:           o.IgnoreHeader,
		EmptyText:              o.EmptyText,
		Columns:                slices.Clone(o.Columns),
		Definitions:            maps.Clone(o.Definitions),
		TypeReader:             o.TypeReader,
		PrettyFormatCharacters: o.PrettyFormatCharacters,
	}
}

// NewOptions creates option's instance using [OptionCallback]
func NewOptions(ocbs ...OptionCallback) *Options {
	options := &Options{
		Name:            "",
		SampleSize:      5,
		SampleMaxLength: 15,
		MaxRows:         -1,
		Separator:       ',',
		IgnoreHeader:    false,
		Definitions:     make(map[string]*Definition, 0),
		EmptyText:       "<empty>",
		PrettyFormatCharacters: []rune{
			'┌', // table head start
			'┬', // table head col separator
			'┐', // table head end
			'│', // row start and end
			'┆', // column separator
			'╞', // table body section start
			'╪', // table body section col separator
			'═', // table body section repeater
			'╡', // table body section end
			'└', // table foot start
			'┴', // table foot col separator
			'┘', // table foot end
			'─', // column filler repeater
			'-', // column type separator repeater
			' ', // space around content
		},
	}

	for _, cb := range ocbs {
		cb(options)
	}

	return options
}

func WithName(name string) OptionCallback {
	return func(o *Options) {
		o.Name = name
	}
}

func WithEmptyText(et string) OptionCallback {
	return func(o *Options) {
		o.EmptyText = et
	}
}

func WithIgnoreHeader(ih bool) OptionCallback {
	return func(o *Options) {
		o.IgnoreHeader = ih
	}
}

func WithMaxRows(s int) OptionCallback {
	return func(o *Options) {
		o.MaxRows = s
	}
}

func WithSampleSize(s int) OptionCallback {
	return func(o *Options) {
		o.SampleSize = s
	}
}

func WithSampleMaxLength(s int) OptionCallback {
	return func(o *Options) {
		o.SampleMaxLength = s
	}
}

func WithSeparator(sep byte) OptionCallback {
	return func(o *Options) {
		o.Separator = sep
	}
}

func WithColumns(cols ...string) OptionCallback {
	return func(o *Options) {
		o.Columns = cols
	}
}

func WithTypeReader(cb func(int, string) reflect.Type) OptionCallback {
	return func(o *Options) {
		o.TypeReader = cb
	}
}

func WithDefinition(name string, def *Definition) OptionCallback {
	return func(o *Options) {
		o.Definitions[name] = def
	}
}

func WithDefinitionType(name string, tp reflect.Type) OptionCallback {
	return func(o *Options) {
		o.Definitions[name] = NewDefinition(tp)
	}
}

func WithPrettyFormatCharacters(chars []rune) OptionCallback {
	return func(o *Options) {
		if len(chars) >= 15 {
			o.PrettyFormatCharacters = chars
		}
	}
}

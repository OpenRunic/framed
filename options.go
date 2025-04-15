package framed

import (
	"maps"
	"reflect"
	"slices"
)

// Options defines settings for table data
type Options struct {

	// sample rows count
	Sampling int

	// printable sample size
	SampleSize int

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
}

// OptionCallback defines function signature for option builder
type OptionCallback = func(*Options)

// Clone duplicates the options as new instance
func (o *Options) Clone() *Options {
	return &Options{
		MaxRows:      o.MaxRows,
		Separator:    o.Separator,
		Sampling:     o.Sampling,
		SampleSize:   o.SampleSize,
		IgnoreHeader: o.IgnoreHeader,
		Columns:      slices.Clone(o.Columns),
		Definitions:  maps.Clone(o.Definitions),
		TypeReader:   o.TypeReader,
	}
}

// NewOptions creates option's instance using [OptionCallback]
func NewOptions(ocbs ...OptionCallback) *Options {
	options := &Options{
		Sampling:     2,
		SampleSize:   10,
		MaxRows:      -1,
		Separator:    ',',
		IgnoreHeader: false,
		Definitions:  make(map[string]*Definition, 0),
	}

	for _, cb := range ocbs {
		cb(options)
	}

	return options
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

func WithSampling(s int) OptionCallback {
	return func(o *Options) {
		o.Sampling = s
	}
}

func WithSampleSize(s int) OptionCallback {
	return func(o *Options) {
		o.SampleSize = s
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

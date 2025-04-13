package framed

import (
	"reflect"
)

var StringTranslatableKinds = []reflect.Kind{
	reflect.String,
	reflect.Int, reflect.Int32, reflect.Int64,
	reflect.Float32, reflect.Float64,
	reflect.Bool,
}

// ColumnDefinition defines details and encoder/decoder for columns
type ColumnDefinition struct {
	Label   string
	Type    reflect.Type
	Encoder func(any) (string, error)
	Decoder func(string) (any, error)
}

func (d *ColumnDefinition) Kind() reflect.Kind {
	return d.Type.Kind()
}

func (d *ColumnDefinition) WithLabel(val string) *ColumnDefinition {
	d.Label = val
	return d
}

func (d *ColumnDefinition) WithEncoder(cb func(any) (string, error)) *ColumnDefinition {
	d.Encoder = cb
	return d
}

func (d *ColumnDefinition) WithDecoder(cb func(string) (any, error)) *ColumnDefinition {
	d.Decoder = cb
	return d
}

func NewDefinition(tp reflect.Type) *ColumnDefinition {
	return &ColumnDefinition{
		Type: tp,
	}
}

func ToType[T any](v T) reflect.Type {
	tp := reflect.TypeOf(v)
	if tp.Kind() == reflect.Slice {
		return tp.Elem()
	}
	return tp
}

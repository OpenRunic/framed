package framed

import (
	"reflect"
)

// Definition defines details and encoder/decoder for values
type Definition struct {
	Label   string
	Type    reflect.Type
	Encoder func(any) (string, error)
	Decoder func(string) (any, error)
}

// Kind returns the reflect.Kind from stored reflect.Type
func (d *Definition) Kind() reflect.Kind {
	return d.Type.Kind()
}

// WithLabel changes the label for current definition
func (d *Definition) WithLabel(val string) *Definition {
	d.Label = val
	return d
}

// WithEncoder updates the encoder to use for data
func (d *Definition) WithEncoder(cb func(any) (string, error)) *Definition {
	d.Encoder = cb
	return d
}

// WithEncoder updates the decoder to use for data
func (d *Definition) WithDecoder(cb func(string) (any, error)) *Definition {
	d.Decoder = cb
	return d
}

// NewDefinition creates new definition instance with reflect.Type
func NewDefinition(tp reflect.Type) *Definition {
	return &Definition{
		Type: tp,
	}
}

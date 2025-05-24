package framed

import (
	"reflect"
)

// Definition defines details and encoder/decoder for values
type Definition struct {
	Type         reflect.Type
	EmptyChecker func(any) bool
	Encoder      func(any) (string, error)
	Decoder      func(string) (any, error)
}

// Kind returns the reflect.Kind from stored reflect.Type
func (d *Definition) Kind() reflect.Kind {
	return d.Type.Kind()
}

// Copy copies extra details of definition if type matches
func (d *Definition) Copy(rd *Definition) *Definition {
	if rd != nil && d.Type == rd.Type {
		if rd.Encoder != nil {
			d.UseEncoder(rd.Encoder)
		}
		if rd.Decoder != nil {
			d.UseDecoder(rd.Decoder)
		}
		if rd.EmptyChecker != nil {
			d.UseEmptyChecker(rd.EmptyChecker)
		}
	}
	return d
}

// UseEncoder updates the encoder to use for data
func (d *Definition) UseEncoder(cb func(any) (string, error)) *Definition {
	d.Encoder = cb
	return d
}

// UseDecoder updates the decoder to use for data
func (d *Definition) UseDecoder(cb func(string) (any, error)) *Definition {
	d.Decoder = cb
	return d
}

// UseEncoder updates the encoder to use for data
func (d *Definition) UseEmptyChecker(cb func(any) bool) *Definition {
	d.EmptyChecker = cb
	return d
}

// NewDefinition creates new definition instance with reflect.Type
func NewDefinition(tp reflect.Type) *Definition {
	return &Definition{
		Type: tp,
	}
}

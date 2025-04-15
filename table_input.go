package framed

import (
	"io"
	"net/http"
	"os"
)

// File opens file and creates table
func File(path string, cbs ...OptionCallback) (*Table, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return Reader(file, cbs...)
}

// URL sends http request to url and creates table
func URL(uri string, cbs ...OptionCallback) (*Table, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	return Reader(response.Body, cbs...)
}

// Reader iterates through [io.Reader] and creates table
func Reader(r io.Reader, cbs ...OptionCallback) (*Table, error) {
	c, ok := r.(io.ReadCloser)
	if ok {
		defer c.Close()
	}

	df := New(cbs...)

	err := df.InsertGenBytes(ReaderToLines(r))
	if err != nil {
		return nil, err
	}

	return df, nil
}

// Lines imports slices of string to create table
func Lines(lines []string, cbs ...OptionCallback) (*Table, error) {
	df := New(cbs...)

	err := df.InsertLines(lines)
	if err != nil {
		return nil, err
	}

	return df, nil
}

// Raw imports slices of raw data to create table
func Raw(ss [][]string, cbs ...OptionCallback) (*Table, error) {
	df := New(cbs...)

	err := df.InsertSlices(ss)
	if err != nil {
		return nil, err
	}

	return df, nil
}

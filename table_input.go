package framed

import (
	"io"
	"net/http"
	"os"
)

// create table from file path
func File(path string, cbs ...OptionCallback) (*Table, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return Reader(file, cbs...)
}

// create table from external uri
func URL(uri string, cbs ...OptionCallback) (*Table, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	return Reader(response.Body, cbs...)
}

// create table from io.Reader or io.ReadCloser
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

// create table from slice of strings
func Lines(lines []string, cbs ...OptionCallback) (*Table, error) {
	df := New(cbs...)

	err := df.InsertLines(lines)
	if err != nil {
		return nil, err
	}

	return df, nil
}

// create table from slice of raw data
func Raw(ss [][]string, cbs ...OptionCallback) (*Table, error) {
	df := New(cbs...)

	err := df.InsertSlices(ss)
	if err != nil {
		return nil, err
	}

	return df, nil
}

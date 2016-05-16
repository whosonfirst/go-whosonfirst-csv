package csv

import (
	gocsv "encoding/csv"
	"io"
	"os"
)

type DictReader struct {
	Reader     *gocsv.Reader
	Fieldnames []string
}

type DictWriter struct {
	Writer     *gocsv.Writer
	Fieldnames []string
}

// to do: update to take io.Reader and not a string
// https://golang.org/pkg/encoding/csv/#NewReader
// (20160516/thisisaaronland)

func NewDictReader(path string) (*DictReader, error) {

	body, open_err := os.Open(path)

	if open_err != nil {
		return nil, open_err
	}

	reader := gocsv.NewReader(body)

	row, read_err := reader.Read()

	if read_err != nil {
		return nil, read_err
	}

	dr := DictReader{Reader: reader, Fieldnames: row}
	return &dr, nil
}

func (dr DictReader) Read() (map[string]string, error) {

	row, err := dr.Reader.Read()

	if err != nil {
		return nil, err
	}

	dict := make(map[string]string)

	for i, value := range row {
		key := dr.Fieldnames[i]
		dict[key] = value
	}

	return dict, nil
}

func NewDictWriter(fh io.Writer, fieldnames []string) (*DictWriter, error) {

	writer := gocsv.NewWriter(fh)

	dw := DictWriter{Writer: writer, Fieldnames: fieldnames}
	return &dw, nil
}

func (dw DictWriter) WriteHeader() {
	dw.Writer.Write(dw.Fieldnames)
	dw.Writer.Flush()
}

// to do - check flags for whether or not to be liberal when missing keys
// (20160516/thisisaaronland)

func (dw DictWriter) WriteRow(row map[string]string) {

	out := make([]string, 0)

	for _, k := range dw.Fieldnames {

		v, ok := row[k]

		if !ok {
			v = ""
		}

		out = append(out, v)
	}

	dw.Writer.Write(out)
	dw.Writer.Flush() // move me somewhere more sensible
}

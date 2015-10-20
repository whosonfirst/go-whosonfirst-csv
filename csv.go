package csv

import (
	gocsv "encoding/csv"
	"os"
)

type DictReader struct {
	Reader     *gocsv.Reader
	Fieldnames []string
}

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

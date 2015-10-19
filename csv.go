package csv

import (
       gocsv "encoding/csv"
       "os"
       "io"
       "flag"
       "fmt"
)

type DictReader struct {
     Reader *csv.Reader
     Fieldnames []string
}

func NewDictReader(path string) *DictReader {
     
	body, _ := os.Open(path)
	reader := gocsv.NewReader(body)

	row, _ := reader.Read()

	dr := DictReader{Reader: reader, Fieldnames: row}
	return &dr
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

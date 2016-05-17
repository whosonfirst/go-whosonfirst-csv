# go-whosonfirst-csv

A simple Go package to implement a "dict reader" style CSV parser (on top of the default `encoding/csv` package) to return rows a key-value dictionaries rather than lists.

## Usage

### Simple

```
import (
	csv "github.com/whosonfirst/go-whosonfirst-csv"
)

reader, reader_err := csv.NewDictReaderFromFile(csv_file)

if reader_err != nil {
	return reader_err
}

for {
	row, err := reader.Read()

	if err == io.EOF {
		break
	}

	if err != nil {
		return err
	}

	value, ok := row["some-key"]

	// and so on...
}
```

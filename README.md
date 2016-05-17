# go-whosonfirst-csv

A simple Go package to implement a "dict reader" style CSV parser (on top of the default `encoding/csv` package) to return rows a key-value dictionaries rather than lists.

## Setup

Run the handy `build` target in the [Makefile](Makefile).

```
make build
```

## Usage

### Reading files

```
import (
	csv "github.com/whosonfirst/go-whosonfirst-csv"
	"os"
)

reader, reader_err := csv.NewDictReaderFromFile(csv_file)

// or maybe you might do
// reader, err := csv.NewDictReader(os.Stdin)

if err != nil {
	panic(err)
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

### Writing files

```
import (
	csv "github.com/whosonfirst/go-whosonfirst-csv"
	"os"
)

fieldnames := []string{"foo", "bar"}

writer, err := csv.NewDictWriter(os.Stdout, fieldnames)

// or maybe you might do
// writer, err := csv.NewDictWriterFromPath("new.csv", fieldnames)

if err != nil {
	panic(err)
}

writer.WriteHeader()

row := make(map[string]string)
row["foo"] = "hello"
row["bar"] = "world"

writer.WriteRow(row)
```

## Tools

### 

## See also

* https://golang.org/pkg/encoding/csv/
* https://docs.python.org/2/library/csv.html
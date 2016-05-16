package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-csv"
	"io"
	"os"
	"strings"
	"sync"
)

func main() {

	var str_cols = flag.String("columns", "", "Columns to filter on")

	flag.Parse()

	cols := strings.Split(*str_cols, ",")

	if len(cols) == 0 {
		panic("NO COLUMNS")
	}

	files := flag.Args()

	writer, err := csv.NewDictWriter(os.Stdout, cols)

	if err != nil {
		panic(err)
	}

	writer.WriteHeader()

	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for _, path := range files {

		wg.Add(1)

		go func(path string, cols []string) {

			defer wg.Done()

			reader, err := csv.NewDictReader(path)

			if err != nil {
				fmt.Println(path, err)
				return
			}

			for {

				row, err := reader.Read()

				if err == io.EOF {
					break
				}

				if err != nil {
					continue
				}

				mu.Lock()

				writer.WriteRow(row)
				mu.Unlock()

			}

		}(path, cols)

	}

	wg.Wait()

}

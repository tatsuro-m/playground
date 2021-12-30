package seed

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io"
	"os"
	"path/filepath"
)

func Exec() error {
	p, _ := filepath.Abs("../../KEN_ALL_ROME.CSV")
	utf8F, _ := os.OpenFile(p, os.O_RDONLY, 0666)
	defer utf8F.Close()
	r := csv.NewReader(transform.NewReader(utf8F, japanese.ShiftJIS.NewDecoder()))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Printf("%#v\n", record)
	}

	return nil
}

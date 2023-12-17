package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	data := "name, age, networth\n" + "Robert, 45, $.25B\n" + "Axe, 30, $.230M"

	reader := csv.NewReader(strings.NewReader(data));

	writer := csv.NewWriter(os.Stdout)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)

		newRecord := append(record, "verified")
		writer.Write(newRecord)
	}

	writer.Flush()
}
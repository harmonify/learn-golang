package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	csvString := "eko,kurniawan,khannedy\n" +
		"budi,pratama,nugraha\n" +
		"joko,morro,diah"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	fmt.Println("===============================")

	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"eko", "kurniawan", "khannedy"})
	_ = writer.Write([]string{"budi", "pratama", "nugraha"})
	_ = writer.Write([]string{"joko", "morro", "diah"})

	writer.Flush()
}

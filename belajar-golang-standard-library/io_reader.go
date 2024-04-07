// The `io` library only provide basic interfaces, i.e:
// `Reader`
// `Writer`
// To create readers or writers, use the `bufio` library
package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("this is long string\nwith new line\n")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}
}

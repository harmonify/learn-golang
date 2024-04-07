// The `io` library only provide basic interfaces, i.e:
// `Reader`
// `Writer`
// To create readers or writers, use the `bufio` library
package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	_, _ = writer.WriteString("Hello World!\n")
	_, _ = writer.WriteString("Selamat belajar\n")

	writer.Flush()
}

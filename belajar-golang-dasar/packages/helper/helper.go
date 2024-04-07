package helper

var Application = "Golang"
var version = "1.0.0"

// The first word of the function name should be uppercased
// to be able to access it from another package
// To learn more, search about Golang's access modifier
func SayHello(name string) string {
	return "Hello " + name
}

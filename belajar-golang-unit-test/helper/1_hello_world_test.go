package helper

import (
	"testing"
)

func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// unit test failed
		// t.Log to log test messages
		// t.Fail to fail an unit test, but will continue its execution
		// t.Error is like t.Log followed by t.Fail
		t.Error("Result is not 'Hello Eko'")
	}
	t.Log("TestHelloWorldEko done")
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")
	if result != "Hello Khannedy" {
		// t.FailNow to fail an unit test and exit immediately
		// t.Fatal is like t.Log followed by t.FailNow
		t.Fatal("Result is not 'Hello Khannedy'")
	}
	t.Log("TestHelloWorldKhannedy done") // will not be executed if this test is failed
}

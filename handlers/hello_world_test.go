package handlers_test

import (
	"testing"

	"github.com/keaaa/hello-world-tutorial/handlers"
)

func Test_hello_world_response(t *testing.T) {
	msg := handlers.GetHelloWorldMsg()

	if msg != "Hello world!" {
		t.Error()
	}
}

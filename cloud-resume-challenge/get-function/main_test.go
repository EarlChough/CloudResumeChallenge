package main

import (
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	if os.Getenv("Handler") != ""{
		t.Skip("skip testing for this function")
	}
}

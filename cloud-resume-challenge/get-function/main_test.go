package main

import (
	"os"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	if os.Getenv("Handler") != ""{
		t.Skip("skip testing for this function")
	}
}

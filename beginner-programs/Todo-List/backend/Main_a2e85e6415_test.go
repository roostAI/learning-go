// Test generated by RoostGPT for test GoUnitTest using AI Type Open AI and AI Model gpt-4

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/router"
)

func main() {
	r := router.Router()

	fmt.Println("Starting server on the port 3000...")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func TestMain_a2e85e6415(t *testing.T) {
	r := router.Router()

	ts := httptest.NewServer(r)
	defer ts.Close()

	// Test case 1: Check if server is up and running
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("Expected server to start, but got error:", err)
	} else {
		t.Log("Success: Server started successfully")
	}

	// Test case 2: Check if server returns 200 status code
	if res.StatusCode != http.StatusOK {
		t.Error("Expected status code 200, but got:", res.StatusCode)
	} else {
		t.Log("Success: Server returned status code 200")
	}
}

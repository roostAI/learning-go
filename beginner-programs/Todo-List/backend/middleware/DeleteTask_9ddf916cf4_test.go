package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDeleteTask_9ddf916cf4(t *testing.T) {
	// Mock http.ResponseWriter
	w := httptest.NewRecorder()

	// Test case 1: Valid id
	r, _ := http.NewRequest("DELETE", "/task/valid_id", nil)
	DeleteTask(w, r)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got '%v'", resp.Status)
	}
	if string(body) != "valid_id" {
		t.Errorf("Expected 'valid_id', got '%v'", string(body))
	}

	// Test case 2: Invalid id
	r, _ = http.NewRequest("DELETE", "/task/invalid_id", nil)
	DeleteTask(w, r)

	resp = w.Result()
	body, _ = ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status NotFound, got '%v'", resp.Status)
	}
	if string(body) != "Task not found" {
		t.Errorf("Expected 'Task not found', got '%v'", string(body))
	}
}

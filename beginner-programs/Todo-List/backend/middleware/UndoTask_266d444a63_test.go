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
	"net/http/httptest"
)

func TestUndoTask_266d444a63(t *testing.T) {
	req, err := http.NewRequest("GET", "/undoTask", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UndoTask)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{"id":"123"}`
	assert.Equal(t, expected, rr.Body.String())
}

func TestUndoTask_266d444a63_Failure(t *testing.T) {
	req, err := http.NewRequest("GET", "/undoTask", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UndoTask)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

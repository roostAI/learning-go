package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/stretchr/testify/assert"
)





type DeleteResult struct {
	DeletedCount int64 `bson:"n"` // The number of documents deleted.
}

type DeleteOptions struct {
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation

	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment interface{}

	// The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.4 will return an error
	// if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option
	// is specified. The driver will return an error if this option is specified during an unacknowledged write
	// operation. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint interface{}

	// Specifies parameters for the delete expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let interface{}
}

type File struct {
	*file // os specific
}

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func (m *mockErrorCollection) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return nil, fmt.Errorf("mock delete error")
}
func TestdeleteOneTask(t *testing.T) {

	mockClient := mongomock.NewMockClient()
	collection = mockClient.Database(dbName).Collection(collName)

	tests := []struct {
		name          string
		taskID        string
		mockSetup     func()
		expectedCount int64
		expectError   bool
	}{
		{
			name:   "Successfully delete a task with a valid ID",
			taskID: "507f1f77bcf86cd799439011",
			mockSetup: func() {

				_, _ = collection.InsertOne(context.Background(), bson.M{"_id": primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")})
			},
			expectedCount: 1,
			expectError:   false,
		},
		{
			name:          "Attempt to delete a task with an invalid ID format",
			taskID:        "invalid-id",
			mockSetup:     func() {},
			expectedCount: 0,
			expectError:   true,
		},
		{
			name:   "Attempt to delete a task that does not exist",
			taskID: "507f1f77bcf86cd799439012",
			mockSetup: func() {

			},
			expectedCount: 0,
			expectError:   false,
		},
		{
			name:   "Handling an error from the database during deletion",
			taskID: "507f1f77bcf86cd799439011",
			mockSetup: func() {

				_, _ = collection.InsertOne(context.Background(), bson.M{"_id": primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")})
				collection = &mockErrorCollection{}
			},
			expectedCount: 0,
			expectError:   true,
		},
		{
			name:   "Verify logging of deleted document count",
			taskID: "507f1f77bcf86cd799439011",
			mockSetup: func() {
				_, _ = collection.InsertOne(context.Background(), bson.M{"_id": primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")})
			},
			expectedCount: 1,
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			var output string
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			defer func() {
				os.Stdout = old
			}()

			deleteOneTask(tt.taskID)

			w.Close()
			fmt.Fscanf(r, "%s", &output)

			if !tt.expectError {
				assert.Equal(t, fmt.Sprintf("Deleted Document %d", tt.expectedCount), output)
			}

			if tt.expectError {
				assert.NotEqual(t, fmt.Sprintf("Deleted Document %d", tt.expectedCount), output)
			}
		})
	}
}


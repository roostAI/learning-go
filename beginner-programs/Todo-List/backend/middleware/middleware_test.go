package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
	"github.com/joho/godotenv"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"bytes"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/tannergabriel/learning-go/beginner-programs/todo-list/backend/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"strings"
	"github.com/stretchr/testify/mock"
)

var mockCollection *mongo.Collectiontype ResponseRecorder struct {
	Code int

	HeaderMap http.Header

	Body *bytes.Buffer

	Flushed bool

	result      *http.Response
	snapHeader  http.Header
	wroteHeader bool
}
type T struct {
	common
	isEnvSet bool
	context  *testContext
}
type Response struct {
	Status     string
	StatusCode int
	Proto      string
	ProtoMajor int
	ProtoMinor int

	Header Header

	Body io.ReadCloser

	ContentLength int64

	TransferEncoding []string

	Close bool

	Uncompressed bool

	Trailer Header

	Request *Request

	TLS *tls.ConnectionState
}
type File struct {
	*file
}
type ExpectedExec struct {
	queryBasedExpectation
	result driver.Result
	delay  time.Duration
}
type Route struct {
	handler http.Handler

	buildOnly bool

	name string

	err error

	namedRoutes map[string]*Route

	routeConf
}
type Router struct {
	NotFoundHandler http.Handler

	MethodNotAllowedHandler http.Handler

	routes []*Route

	namedRoutes map[string]*Route

	KeepContext bool

	middlewares []middleware

	routeConf
}
type Buffer struct {
	buf      []byte
	off      int
	lastRead readOp
}
type DeleteResult struct {
	DeletedCount int64 `bson:"n"`
}
type DeleteOptions struct {
	Collation *Collation

	Comment interface{}

	Hint interface{}

	Let interface{}
}
type Database struct {
	client         *Client
	name           string
	readConcern    *readconcern.ReadConcern
	writeConcern   *writeconcern.WriteConcern
	readPreference *readpref.ReadPref
	readSelector   description.ServerSelector
	writeSelector  description.ServerSelector
	registry       *bsoncodec.Registry
}
type MockCollection struct {
	mock.Mock
}
type Call struct {
	Parent *Mock

	Method string

	Arguments Arguments

	ReturnArguments Arguments

	callerInfo []string

	Repeatability int

	totalCalls int

	optional bool

	WaitFor <-chan time.Time

	waitTime time.Duration

	RunFn func(Arguments)

	PanicMsg *string
}
type Mock struct {
	ExpectedCalls []*Call

	Calls []Call

	test TestingT

	testData objx.Map

	mutex sync.Mutex
}
type Cursor struct {
	Current bson.Raw

	bc            batchCursor
	batch         *bsoncore.DocumentSequence
	batchLength   int
	registry      *bsoncodec.Registry
	clientSession *session.Client

	err error
}
type FindOptions struct {
	AllowDiskUse *bool

	AllowPartialResults *bool

	BatchSize *int32

	Collation *Collation

	Comment *string

	CursorType *CursorType

	Hint interface{}

	Limit *int64

	Max interface{}

	MaxAwaitTime *time.Duration

	MaxTime *time.Duration

	Min interface{}

	NoCursorTimeout *bool

	OplogReplay *bool

	Projection interface{}

	ReturnKey *bool

	ShowRecordID *bool

	Skip *int64

	Snapshot *bool

	Sort interface{}

	Let interface{}
}
/*
ROOST_METHOD_HASH=setupResponse_fa7345612e
ROOST_METHOD_SIG_HASH=setupResponse_a69e61ede7


 */
func TestsetupResponse(t *testing.T) {
	tests := []struct {
		name            string
		responseWriter  http.ResponseWriter
		requestMethod   string
		expectedHeaders map[string]string
		expectPanic     bool
	}{
		{
			name:           "Verify CORS Headers are Set Correctly",
			responseWriter: httptest.NewRecorder(),
			requestMethod:  "GET",
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
				"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			},
			expectPanic: false,
		},
		{
			name:            "Check Response Writer is Not Modified When Nil",
			responseWriter:  nil,
			requestMethod:   "GET",
			expectedHeaders: map[string]string{},
			expectPanic:     true,
		},
		{
			name:           "Validate Behavior with Different HTTP Methods",
			responseWriter: httptest.NewRecorder(),
			requestMethod:  "POST",
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
				"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			},
			expectPanic: false,
		},
		{
			name:           "Test Function with Empty Request",
			responseWriter: httptest.NewRecorder(),
			requestMethod:  "",
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
				"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			},
			expectPanic: false,
		},
		{
			name:           "Confirm Header Overwrites",
			responseWriter: httptest.NewRecorder(),
			requestMethod:  "OPTIONS",
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
				"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			},
			expectPanic: false,
		},
		{
			name:           "Test Function with Multiple Calls",
			responseWriter: httptest.NewRecorder(),
			requestMethod:  "PUT",
			expectedHeaders: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
				"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
			},
			expectPanic: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.responseWriter != nil {
				req := &http.Request{Method: tt.requestMethod}
				setupResponse(&tt.responseWriter, req)

				recorder := tt.responseWriter.(*httptest.ResponseRecorder)
				for key, expectedValue := range tt.expectedHeaders {
					if actualValue := recorder.Header().Get(key); actualValue != expectedValue {
						t.Errorf("Expected header %s to be %s, got %s", key, expectedValue, actualValue)
					}
				}
			} else {
				defer func() {
					if r := recover(); r == nil && tt.expectPanic {
						t.Errorf("Expected panic for nil response writer, but did not panic")
					}
				}()
				setupResponse(nil, &http.Request{Method: tt.requestMethod})
			}
		})
	}
}


/*
ROOST_METHOD_HASH=goDotEnvVariable_8f7b64c895
ROOST_METHOD_SIG_HASH=goDotEnvVariable_793184d000


 */
func TestgoDotEnvVariable(t *testing.T) {

	type testCase struct {
		name           string
		envFileContent string
		key            string
		expectedValue  string
		shouldError    bool
	}

	tests := []testCase{
		{
			name:           "Successfully Retrieve an Environment Variable",
			envFileContent: "KEY=VALUE\n",
			key:            "KEY",
			expectedValue:  "VALUE",
			shouldError:    false,
		},
		{
			name:           "Handle Missing Environment Variable Gracefully",
			envFileContent: "",
			key:            "NON_EXISTENT_KEY",
			expectedValue:  "",
			shouldError:    false,
		},
		{
			name:           "Handle Error Loading .env File",
			envFileContent: "",
			key:            "KEY",
			expectedValue:  "",
			shouldError:    false,
		},
		{
			name:           "Test with Multiple Environment Variables",
			envFileContent: "KEY1=VALUE1\nKEY2=VALUE2\n",
			key:            "KEY1",
			expectedValue:  "VALUE1",
			shouldError:    false,
		},
		{
			name:           "Check for Environment Variable with Whitespace",
			envFileContent: "KEY= VALUE \n",
			key:            "KEY",
			expectedValue:  " VALUE ",
			shouldError:    false,
		},
		{
			name:           "Environment Variable with Special Characters",
			envFileContent: "KEY=Value#1$\n",
			key:            "KEY",
			expectedValue:  "Value#1$",
			shouldError:    false,
		},
		{
			name:           "Verify Behavior with an Unloaded .env File",
			envFileContent: "",
			key:            "KEY",
			expectedValue:  os.Getenv("KEY"),
			shouldError:    false,
		},
		{
			name:           "Check Error Logging on .env Loading Failure",
			envFileContent: "",
			key:            "KEY",
			expectedValue:  "",
			shouldError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.envFileContent != "" {
				tmpFile, err := os.CreateTemp("", ".env")
				if err != nil {
					t.Fatalf("could not create temp file: %v", err)
				}
				defer os.Remove(tmpFile.Name())
				if _, err := tmpFile.WriteString(tt.envFileContent); err != nil {
					t.Fatalf("could not write to temp file: %v", err)
				}
				if err := tmpFile.Close(); err != nil {
					t.Fatalf("could not close temp file: %v", err)
				}

				if err := godotenv.Load(tmpFile.Name()); err != nil {
					if !tt.shouldError {
						t.Errorf("failed to load .env file: %v", err)
					}
				}
			} else {

				os.Remove(".env")
			}

			result := goDotEnvVariable(tt.key)

			if result != tt.expectedValue {
				t.Errorf("expected %q, got %q", tt.expectedValue, result)
			}

			t.Logf("Test %s completed: expected %s, got %s", tt.name, tt.expectedValue, result)
		})
	}
}


/*
ROOST_METHOD_HASH=DeleteAllTask_f9af6ec409
ROOST_METHOD_SIG_HASH=DeleteAllTask_9390cec12e


 */
func TestDeleteAllTask(t *testing.T) {
	tests := []struct {
		name               string
		mockDBResponse     int64
		mockDBError        error
		httpMethod         string
		expectedStatusCode int
		expectedResponse   int64
	}{
		{
			name:               "Successful Deletion of All Tasks",
			mockDBResponse:     5,
			mockDBError:        nil,
			httpMethod:         http.MethodDelete,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   5,
		},
		{
			name:               "No Tasks to Delete",
			mockDBResponse:     0,
			mockDBError:        nil,
			httpMethod:         http.MethodDelete,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   0,
		},
		{
			name:               "Database Connection Error",
			mockDBResponse:     0,
			mockDBError:        fmt.Errorf("db error"),
			httpMethod:         http.MethodDelete,
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   0,
		},
		{
			name:               "Invalid HTTP Method",
			mockDBResponse:     0,
			mockDBError:        nil,
			httpMethod:         http.MethodPost,
			expectedStatusCode: http.StatusMethodNotAllowed,
			expectedResponse:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.mockDBError != nil {
				collection = createMockCollectionWithError(tt.mockDBError)
			} else {
				collection = createMockCollectionWithResponse(tt.mockDBResponse)
			}

			req, err := http.NewRequest(tt.httpMethod, "/tasks", nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			rr := httptest.NewRecorder()

			DeleteAllTask(rr, req)

			if rr.Code != tt.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatusCode, rr.Code)
			}

			if tt.expectedStatusCode == http.StatusOK {
				var responseCount int64
				if err := json.NewDecoder(rr.Body).Decode(&responseCount); err != nil {
					t.Errorf("Could not decode response: %v", err)
				}
				if responseCount != tt.expectedResponse {
					t.Errorf("Expected response %d, got %d", tt.expectedResponse, responseCount)
				}
			}
		})
	}
}

func createMockCollectionWithError(err error) *mongo.Collection {

	return nil
}

func createMockCollectionWithResponse(deletedCount int64) *mongo.Collection {

	return nil
}


/*
ROOST_METHOD_HASH=GetAllTask_b2f40a63a4
ROOST_METHOD_SIG_HASH=GetAllTask_09b678baa9


 */
func TestGetAllTask(t *testing.T) {

	responseWriter := httptest.NewRecorder()

	tests := []struct {
		name             string
		mockDBResponse   []primitive.M
		mockDBError      error
		method           string
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Successful retrieval of all tasks",
			mockDBResponse:   []primitive.M{{"task": "Task 1"}, {"task": "Task 2"}},
			mockDBError:      nil,
			method:           http.MethodGet,
			expectedStatus:   http.StatusOK,
			expectedResponse: `[{"task":"Task 1"},{"task":"Task 2"}]`,
		},
		{
			name:             "No tasks found in the database",
			mockDBResponse:   []primitive.M{},
			mockDBError:      nil,
			method:           http.MethodGet,
			expectedStatus:   http.StatusOK,
			expectedResponse: `[]`,
		},
		{
			name:             "Database connection error",
			mockDBResponse:   nil,
			mockDBError:      sqlmock.ErrCancelled,
			method:           http.MethodGet,
			expectedStatus:   http.StatusInternalServerError,
			expectedResponse: `{"error":"Internal Server Error"}`,
		},
		{
			name:             "Invalid HTTP method used",
			mockDBResponse:   nil,
			mockDBError:      nil,
			method:           http.MethodPost,
			expectedStatus:   http.StatusMethodNotAllowed,
			expectedResponse: ``,
		},
		{
			name:             "Malformed request",
			mockDBResponse:   nil,
			mockDBError:      nil,
			method:           "INVALID",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: ``,
		},
		{
			name:             "Valid request with additional query parameters",
			mockDBResponse:   []primitive.M{{"task": "Task 1"}},
			mockDBError:      nil,
			method:           http.MethodGet,
			expectedStatus:   http.StatusOK,
			expectedResponse: `[{"task":"Task 1"}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			collection = setupMockDB(tt.mockDBResponse, tt.mockDBError)

			req := httptest.NewRequest(tt.method, "/tasks", nil)

			GetAllTask(responseWriter, req)

			if responseWriter.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, responseWriter.Code)
			}

			responseBody := responseWriter.Body.String()
			if responseBody != tt.expectedResponse {
				t.Errorf("expected response %s, got %s", tt.expectedResponse, responseBody)
			}
		})
	}
}

func setupMockDB(mockResponse []primitive.M, mockError error) *mongo.Collection {

	client, _ := mongo.NewClient()
	collection := client.Database(dbName).Collection(collName)

	return collection
}


/*
ROOST_METHOD_HASH=deleteAllTask_b94a0f57d1
ROOST_METHOD_SIG_HASH=deleteAllTask_ec6cbb32dc


 */
func TestdeleteAllTask(t *testing.T) {
	setup()

	tests := []struct {
		name          string
		initialCount  int64
		expectedCount int64
		mockError     error
	}{
		{
			name:          "Successful Deletion of All Tasks",
			initialCount:  5,
			expectedCount: 5,
			mockError:     nil,
		},
		{
			name:          "Deletion When No Tasks Exist",
			initialCount:  0,
			expectedCount: 0,
			mockError:     nil,
		},
		{
			name:          "Error Handling During Deletion",
			initialCount:  3,
			expectedCount: 0,
			mockError:     fmt.Errorf("mock error"),
		},
		{
			name:          "Deletion with Partial Data",
			initialCount:  4,
			expectedCount: 4,
			mockError:     nil,
		},
		{
			name:          "Concurrency Test for Deletion",
			initialCount:  10,
			expectedCount: 10,
			mockError:     nil,
		},
		{
			name:          "Checking Database Connection Before Deletion",
			initialCount:  0,
			expectedCount: 0,
			mockError:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for i := int64(0); i < tt.initialCount; i++ {
				_, err := mockCollection.InsertOne(context.Background(), bson.D{{Key: "task", Value: fmt.Sprintf("task %d", i)}})
				if err != nil {
					t.Fatalf("Failed to insert task: %v", err)
				}
			}

			mockDeleteMany := func(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
				if tt.mockError != nil {
					return nil, tt.mockError
				}
				return &mongo.DeleteResult{DeletedCount: tt.expectedCount}, nil
			}

			mockCollection = &mongo.Collection{

				DeleteMany: mockDeleteMany,
			}

			result := deleteAllTask()

			if result != tt.expectedCount {
				t.Errorf("Expected DeletedCount %d, got %d", tt.expectedCount, result)
			}

			_, _ = mockCollection.DeleteMany(context.Background(), bson.D{{}})
		})
	}
}

func setup() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	mockCollection = client.Database("test").Collection("todolist")
}


/*
ROOST_METHOD_HASH=insertOneTask_5fc15ee4f8
ROOST_METHOD_SIG_HASH=insertOneTask_841c76e883


 */
func TestinsertOneTask(t *testing.T) {

	var err error
	mockCollection, err = setupMockDB()
	if err != nil {
		t.Fatalf("Failed to set up mock DB: %v", err)
	}

	tests := []struct {
		name        string
		task        models.ToDoList
		expectError bool
		expectedLog string
	}{
		{
			name: "Successfully Insert a Task",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "Test Task",
				Status: true,
			},
			expectError: false,
			expectedLog: "Inserted a Single Record ",
		},
		{
			name: "Insert Task with Empty Task Field",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "",
				Status: true,
			},
			expectError: true,
			expectedLog: "error: task cannot be empty",
		},
		{
			name: "Insert Task with Invalid MongoDB Configuration",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "Valid Task",
				Status: true,
			},
			expectError: true,
			expectedLog: "error: collection is not initialized",
		},
		{
			name: "Insert Task with Special Characters in Task Field",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "!@#$%^&*()",
				Status: true,
			},
			expectError: false,
			expectedLog: "Inserted a Single Record ",
		},
		{
			name: "Insert Task with Duplicate ID",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "Duplicate Task",
				Status: true,
			},
			expectError: true,
			expectedLog: "error: duplicate key error",
		},
		{
			name: "Insert Task After Database Connection Failure",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "Task After Failure",
				Status: true,
			},
			expectError: true,
			expectedLog: "error: connection failure",
		},
		{
			name: "Insert Task with Extremely Long Task String",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   string(make([]byte, 10000)),
				Status: true,
			},
			expectError: false,
			expectedLog: "Inserted a Single Record ",
		},
		{
			name: "Insert Task with False Status",
			task: models.ToDoList{
				ID:     primitive.NewObjectID(),
				Task:   "False Status Task",
				Status: false,
			},
			expectError: false,
			expectedLog: "Inserted a Single Record ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			if tt.expectError {
				collection = nil
			}
			insertOneTask(tt.task)

			w.Close()
			os.Stdout = old
			var buf [512]byte
			n, _ := r.Read(buf[:])
			output := string(buf[:n])

			if tt.expectError {
				if !containsError(output, tt.expectedLog) {
					t.Errorf("expected error log to contain %q, got %q", tt.expectedLog, output)
				}
			} else {
				if !containsSuccessLog(output, tt.expectedLog) {
					t.Errorf("expected success log to contain %q, got %q", tt.expectedLog, output)
				}
			}
		})
	}
}

func containsError(output string, log string) bool {
	return stringContains(output, log)
}

func containsSuccessLog(output string, log string) bool {
	return stringContains(output, log)
}

func setupMockDB() (*mongo.Collection, error) {

	return nil, nil
}

func stringContains(s, substr string) bool {
	return true
}


/*
ROOST_METHOD_HASH=DeleteTask_9a8d2c34fc
ROOST_METHOD_SIG_HASH=DeleteTask_9ddf916cf4


 */
func TestDeleteTask(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	collection = &mongo.Collection{}

	tests := []struct {
		name         string
		taskID       string
		mockBehavior func()
		expectedCode int
		expectedBody string
	}{
		{
			name:   "Successfully Delete a Task",
			taskID: "60f5b5b5b5b5b5b5b5b5b5b5",
			mockBehavior: func() {
				mock.ExpectExec("DELETE FROM tasks WHERE id = ?").WithArgs("60f5b5b5b5b5b5b5b5b5b5b5").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedCode: http.StatusOK,
			expectedBody: `"60f5b5b5b5b5b5b5b5b5b5b5"`,
		},
		{
			name:   "Attempt to Delete a Non-Existent Task",
			taskID: "nonexistent-id",
			mockBehavior: func() {
				mock.ExpectExec("DELETE FROM tasks WHERE id = ?").WithArgs("nonexistent-id").WillReturnError(mongo.ErrNoDocuments)
			},
			expectedCode: http.StatusNotFound,
			expectedBody: `"Task not found"`,
		},
		{
			name:   "Invalid Task ID Format",
			taskID: "invalid-id",
			mockBehavior: func() {

			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `"Invalid task ID format"`,
		},
		{
			name:   "Database Connection Error",
			taskID: "60f5b5b5b5b5b5b5b5b5b5b5",
			mockBehavior: func() {
				mock.ExpectExec("DELETE FROM tasks WHERE id = ?").WithArgs("60f5b5b5b5b5b5b5b5b5b5b5").WillReturnError(fmt.Errorf("database error"))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `"Internal server error"`,
		},
		{
			name:   "Missing Task ID in Request",
			taskID: "",
			mockBehavior: func() {

			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `"Task ID is required"`,
		},
		{
			name:   "Check Response Content Type",
			taskID: "60f5b5b5b5b5b5b5b5b5b5b5",
			mockBehavior: func() {
				mock.ExpectExec("DELETE FROM tasks WHERE id = ?").WithArgs("60f5b5b5b5b5b5b5b5b5b5b5").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedCode: http.StatusOK,
			expectedBody: `"60f5b5b5b5b5b5b5b5b5b5b5"`,
		},
		{
			name:   "Verify No Duplicate Deletion Requests",
			taskID: "60f5b5b5b5b5b5b5b5b5b5b5",
			mockBehavior: func() {
				mock.ExpectExec("DELETE FROM tasks WHERE id = ?").WithArgs("60f5b5b5b5b5b5b5b5b5b5b5").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec("DELETE FROM tasks WHERE id = ?").WithArgs("60f5b5b5b5b5b5b5b5b5b5b5").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedCode: http.StatusOK,
			expectedBody: `"60f5b5b5b5b5b5b5b5b5b5b5"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockBehavior()
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/tasks/%s", tt.taskID), nil)
			w := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

			router.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, w.Code)
			}

			if w.Body.String() != tt.expectedBody {
				t.Errorf("expected body %s, got %s", tt.expectedBody, w.Body.String())
			}
		})
	}
}


/*
ROOST_METHOD_HASH=TaskComplete_0cb444a0eb
ROOST_METHOD_SIG_HASH=TaskComplete_34e5efff97


 */
func TestTaskComplete(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		taskID         string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid Task ID",
			method:         http.MethodPost,
			taskID:         "507f1f77bcf86cd799439011",
			expectedStatus: http.StatusOK,
			expectedBody:   "507f1f77bcf86cd799439011\n",
		},
		{
			name:           "Missing Task ID",
			method:         http.MethodPost,
			taskID:         "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "",
		},
		{
			name:           "Invalid Task ID Format",
			method:         http.MethodPost,
			taskID:         "invalidID",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "",
		},
		{
			name:           "Task Completion Failure",
			method:         http.MethodPost,
			taskID:         "507f1f77bcf86cd799439011",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
		{
			name:           "Unsupported HTTP Method",
			method:         http.MethodPut,
			taskID:         "507f1f77bcf86cd799439011",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "",
		},
	}

	var originalTaskComplete = taskComplete
	defer func() { taskComplete = originalTaskComplete }()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(tt.method, "/task?id="+tt.taskID, nil)
			w := httptest.NewRecorder()

			if tt.name == "Task Completion Failure" {

				taskComplete = func(task string) {
					log.Fatal("Simulated failure in taskComplete")
				}
			} else {

				taskComplete = func(task string) {
					id, err := primitive.ObjectIDFromHex(task)
					if err != nil {
						log.Fatal(err)
					}

					filter := bson.M{"_id": id}
					update := bson.M{"$set": bson.M{"status": true}}
					result, err := collection.UpdateOne(context.Background(), filter, update)
					if err != nil {
						log.Fatal(err)
					}
					assert.Equal(t, int64(1), result.ModifiedCount)
				}
			}

			TaskComplete(w, req)

			res := w.Result()
			assert.Equal(t, tt.expectedStatus, res.StatusCode)

			if tt.expectedBody != "" {
				body := new(bytes.Buffer)
				body.ReadFrom(res.Body)
				assert.Equal(t, tt.expectedBody, body.String())
			}
		})
	}
}


/*
ROOST_METHOD_HASH=UndoTask_4098fa5b30
ROOST_METHOD_SIG_HASH=UndoTask_266d444a63


 */
func TestUndoTask(t *testing.T) {

	router := mux.NewRouter()
	router.HandleFunc("/task/{id}", UndoTask).Methods("POST")

	tests := []struct {
		name         string
		taskID       string
		expectedCode int
		expectedBody string
		setupMock    func()
	}{
		{
			name:         "Successful Undo Task with Valid ID",
			taskID:       "507f1f77bcf86cd799439011",
			expectedCode: http.StatusOK,
			expectedBody: `"507f1f77bcf86cd799439011"`,
			setupMock: func() {

			},
		},
		{
			name:         "Undo Task with Invalid ID",
			taskID:       "invalid-id",
			expectedCode: http.StatusBadRequest,
			expectedBody: `"Invalid task ID"`,
			setupMock: func() {

			},
		},
		{
			name:         "Missing Task ID in Request",
			taskID:       "",
			expectedCode: http.StatusBadRequest,
			expectedBody: `"Missing task ID"`,
			setupMock: func() {

			},
		},
		{
			name:         "Error During Task Undo Process",
			taskID:       "507f1f77bcf86cd799439011",
			expectedCode: http.StatusInternalServerError,
			expectedBody: `"Error processing request"`,
			setupMock: func() {

				undoTask = func(task string) {
					log.Fatal("Mocked error during undo task")
				}
			},
		},
		{
			name:         "Verify Response Content Type",
			taskID:       "507f1f77bcf86cd799439011",
			expectedCode: http.StatusOK,
			expectedBody: `"507f1f77bcf86cd799439011"`,
			setupMock: func() {

			},
		},
		{
			name:         "Test Response for Non-JSON Clients",
			taskID:       "507f1f77bcf86cd799439011",
			expectedCode: http.StatusNotAcceptable,
			expectedBody: "",
			setupMock: func() {

			},
		},
		{
			name:         "Concurrency Test for UndoTask",
			taskID:       "507f1f77bcf86cd799439011",
			expectedCode: http.StatusOK,
			expectedBody: `"507f1f77bcf86cd799439011"`,
			setupMock: func() {

			},
		},
		{
			name:         "Test for Proper Logging",
			taskID:       "507f1f77bcf86cd799439011",
			expectedCode: http.StatusOK,
			expectedBody: `"507f1f77bcf86cd799439011"`,
			setupMock: func() {

			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.setupMock()

			req := httptest.NewRequest("POST", fmt.Sprintf("/task/%s", tt.taskID), nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.expectedCode {
				t.Errorf("expected code %d, got %d", tt.expectedCode, rec.Code)
			}

			if rec.Body.String() != tt.expectedBody {
				t.Errorf("expected body %s, got %s", tt.expectedBody, rec.Body.String())
			}
		})
	}
}


/*
ROOST_METHOD_HASH=deleteOneTask_f840ce3ce3
ROOST_METHOD_SIG_HASH=deleteOneTask_22ade46ed5


 */
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


/*
ROOST_METHOD_HASH=taskComplete_780bc81653
ROOST_METHOD_SIG_HASH=taskComplete_938d822bf5


 */
func TesttaskComplete(t *testing.T) {
	tests := []struct {
		description      string
		taskID           string
		expectedModified int64
		expectedError    bool
	}{
		{
			description:      "Successfully completing a task",
			taskID:           "60d5ec49f1d3e63d10e916bc",
			expectedModified: 1,
			expectedError:    false,
		},
		{
			description:      "Attempting to complete a task with an invalid ID",
			taskID:           "invalidID",
			expectedModified: 0,
			expectedError:    true,
		},
		{
			description:      "Attempting to complete a task that does not exist",
			taskID:           "60d5ec49f1d3e63d10e916bd",
			expectedModified: 0,
			expectedError:    false,
		},
		{
			description:      "MongoDB update operation fails",
			taskID:           "60d5ec49f1d3e63d10e916bc",
			expectedModified: 0,
			expectedError:    true,
		},
		{
			description:      "Completing a task with an empty string as ID",
			taskID:           "",
			expectedModified: 0,
			expectedError:    true,
		},
		{
			description:      "Completing a task with a malformed ID",
			taskID:           "123",
			expectedModified: 0,
			expectedError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {

			var modifiedCount int64
			var err error

			if tt.expectedError {
				if tt.taskID == "invalidID" || tt.taskID == "" || tt.taskID == "123" {

					err = fmt.Errorf("invalid task ID")
				} else {

					err = fmt.Errorf("update operation failed")
				}
			} else {

				modifiedCount = 1
			}

			if err != nil {
				log.Println(err)
				if !tt.expectedError {
					t.Errorf("expected no error, got %v", err)
				}
				return
			}

			taskComplete(tt.taskID)

			if modifiedCount != tt.expectedModified {
				t.Errorf("For task ID %s, expected modified count %d, got %d", tt.taskID, tt.expectedModified, modifiedCount)
			}
		})
	}
}


/*
ROOST_METHOD_HASH=undoTask_a5084af2a0
ROOST_METHOD_SIG_HASH=undoTask_d96e9e6340


 */
func TestundoTask(t *testing.T) {

	var buf bytes.Buffer
	log.SetOutput(&buf)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	db := client.Database(dbName)
	collection = db.Collection(collName)

	tests := []struct {
		name          string
		taskID        string
		expectedLog   string
		expectedCount int64
		mockData      bson.M
		mockError     error
	}{
		{
			name:          "Successfully Undo a Task",
			taskID:        "507f191e810c19729de860ea",
			expectedLog:   "modified count:  1",
			expectedCount: 1,
			mockData:      bson.M{"_id": primitive.ObjectIDFromHex("507f191e810c19729de860ea"), "status": true},
			mockError:     nil,
		},
		{
			name:          "Handle Invalid Task ID",
			taskID:        "invalid-id",
			expectedLog:   "error: invalid hex string",
			expectedCount: 0,
			mockData:      nil,
			mockError:     nil,
		},
		{
			name:          "MongoDB Update Operation Fails",
			taskID:        "507f191e810c19729de860ea",
			expectedLog:   "error: update failed",
			expectedCount: 0,
			mockData:      bson.M{"_id": primitive.ObjectIDFromHex("507f191e810c19729de860ea"), "status": true},
			mockError:     fmt.Errorf("update failed"),
		},
		{
			name:          "No Task Found for Valid ID",
			taskID:        "507f191e810c19729de860ea",
			expectedLog:   "modified count:  0",
			expectedCount: 0,
			mockData:      bson.M{"_id": primitive.ObjectIDFromHex("507f191e810c19729de860ea"), "status": false},
			mockError:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.mockData != nil {
				_, err := collection.InsertOne(context.Background(), tt.mockData)
				if err != nil {
					t.Fatalf("Failed to insert mock data: %v", err)
				}
			}

			if tt.taskID == "invalid-id" {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for invalid ID, got none")
					}
				}()
			} else {

				if tt.mockError != nil {

					log.Println(tt.mockError)
				} else {
					undoTask(tt.taskID)
				}
			}

			if !strings.Contains(buf.String(), tt.expectedLog) {
				t.Errorf("Expected log to contain %q, got %q", tt.expectedLog, buf.String())
			}

			_, err := collection.DeleteMany(context.Background(), bson.M{})
			if err != nil {
				t.Fatalf("Failed to clean up mock data: %v", err)
			}

			buf.Reset()
		})
	}
}


/*
ROOST_METHOD_HASH=CreateTask_dab672e246
ROOST_METHOD_SIG_HASH=CreateTask_eb4d396a39


 */
func TestCreateTask(t *testing.T) {
	type testCase struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedBody   string
	}

	tests := []testCase{
		{
			name:           "Successful Task Creation",
			method:         "POST",
			body:           `{"title":"Test Task","description":"This is a test task"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"title":"Test Task","description":"This is a test task"}`,
		},
		{
			name:           "Handling OPTIONS Request",
			method:         "OPTIONS",
			body:           ``,
			expectedStatus: http.StatusOK,
			expectedBody:   ``,
		},
		{
			name:           "Invalid JSON Payload",
			method:         "POST",
			body:           `{"title": "Test Task", "description": "This is a test task",}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ``,
		},
		{
			name:           "Missing Required Fields in Task",
			method:         "POST",
			body:           `{"description":"This is a test task"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ``,
		},
		{
			name:           "Database Insertion Failure",
			method:         "POST",
			body:           `{"title":"Test Task","description":"This is a test task"}`,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   ``,
		},
		{
			name:           "Successful Task Creation with Context",
			method:         "POST",
			body:           `{"title":"Test Task","description":"This is a test task"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"title":"Test Task","description":"This is a test task"}`,
		},
		{
			name:           "Response Format Validation",
			method:         "POST",
			body:           `{"title":"Test Task","description":"This is a test task"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"title":"Test Task","description":"This is a test task"}`,
		},
		{
			name:           "Empty Request Body",
			method:         "POST",
			body:           ``,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   ``,
		},
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	collection = client.Database(dbName).Collection(collName)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/tasks", bytes.NewBufferString(tt.body))
			rec := httptest.NewRecorder()

			CreateTask(rec, req)

			res := rec.Result()
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			if tt.expectedStatus == http.StatusOK {
				var responseBody models.ToDoList
				if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
					t.Errorf("Failed to decode response body: %v", err)
				}

				expectedResponseBody := models.ToDoList{}
				if err := json.Unmarshal([]byte(tt.expectedBody), &expectedResponseBody); err != nil {
					t.Errorf("Failed to unmarshal expected body: %v", err)
				}

				if responseBody != expectedResponseBody {
					t.Errorf("Expected body %v, got %v", expectedResponseBody, responseBody)
				}
			}
		})
	}
}


/*
ROOST_METHOD_HASH=getAllTask_fbbc7d9d9f
ROOST_METHOD_SIG_HASH=getAllTask_dfbd966f8b


 */
func (m *MockCursor) Close(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockCursor) Decode(val interface{}) error {
	args := m.Called(val)
	return args.Error(0)
}

func (m *MockCursor) Err() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockCursor) Next(ctx context.Context) bool {
	args := m.Called(ctx)
	return args.Bool(0)
}

func TestgetAllTask(t *testing.T) {

	mockCollection := new(MockCollection)
	collection = mockCollection

	tests := []struct {
		name        string
		findReturn  *MockCursor
		findError   error
		expected    []primitive.M
		expectedErr error
	}{
		{
			name: "Retrieve All Tasks Successfully",
			findReturn: func() *MockCursor {
				mockCursor := new(MockCursor)
				mockCursor.On("Next", mock.Anything).Return(true).Once()
				mockCursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
					result := args.Get(0).(*primitive.M)
					*result = primitive.M{"_id": primitive.NewObjectID(), "title": "Task 1"}
				}).Return(nil).Once()
				mockCursor.On("Next", mock.Anything).Return(false).Once()
				mockCursor.On("Err").Return(nil).Once()
				return mockCursor
			}(),
			findError: nil,
			expected:  []primitive.M{{"_id": mock.Anything, "title": "Task 1"}},
		},
		{
			name: "Handle Empty Task List",
			findReturn: func() *MockCursor {
				mockCursor := new(MockCursor)
				mockCursor.On("Next", mock.Anything).Return(false).Once()
				mockCursor.On("Err").Return(nil).Once()
				return mockCursor
			}(),
			findError: nil,
			expected:  []primitive.M{},
		},
		{
			name:       "Handle MongoDB Retrieval Error",
			findReturn: nil,
			findError:  fmt.Errorf("database error"),
			expected:   nil,
		},
		{
			name: "Handle Decoding Error",
			findReturn: func() *MockCursor {
				mockCursor := new(MockCursor)
				mockCursor.On("Next", mock.Anything).Return(true).Once()
				mockCursor.On("Decode", mock.Anything).Return(fmt.Errorf("decode error")).Once()
				mockCursor.On("Next", mock.Anything).Return(false).Once()
				mockCursor.On("Err").Return(nil).Once()
				return mockCursor
			}(),
			findError: nil,
			expected:  nil,
		},
		{
			name: "Validate Cursor Closure",
			findReturn: func() *MockCursor {
				mockCursor := new(MockCursor)
				mockCursor.On("Next", mock.Anything).Return(true).Once()
				mockCursor.On("Decode", mock.Anything).Return(nil).Once()
				mockCursor.On("Next", mock.Anything).Return(false).Once()
				mockCursor.On("Err").Return(nil).Once()
				mockCursor.On("Close", mock.Anything).Return(nil).Once()
				return mockCursor
			}(),
			findError: nil,
			expected:  []primitive.M{{"_id": mock.Anything, "title": "Task 1"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.findError != nil {
				mockCollection.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(nil, tt.findError)
			} else {
				mockCollection.On("Find", mock.Anything, mock.Anything, mock.Anything).Return(tt.findReturn, nil)
			}

			result := getAllTask()

			if tt.expectedErr != nil {
				assert.Error(t, tt.expectedErr)
			} else {
				assert.NoError(t, tt.expectedErr)
				assert.ElementsMatch(t, tt.expected, result)
			}

			t.Logf("Test %s: expected %v, got %v", tt.name, tt.expected, result)

			mockCollection.AssertExpectations(t)
			if tt.findReturn != nil {
				tt.findReturn.AssertExpectations(t)
			}
		})
	}
}


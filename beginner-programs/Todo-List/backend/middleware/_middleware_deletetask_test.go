package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type ExpectedExec struct {
	queryBasedExpectation
	result driver.Result
	delay  time.Duration
}

type Route struct {
	// Request handler for the route.
	handler http.Handler
	// If true, this route never matches: it is only used to build URLs.
	buildOnly bool
	// The name used to build URLs.
	name string
	// Error resulted from building a route.
	err error

	// "global" reference to all named routes
	namedRoutes map[string]*Route

	// config possibly passed in from `Router`
	routeConf
}

type Router struct {
	// Configurable Handler to be used when no route matches.
	NotFoundHandler http.Handler

	// Configurable Handler to be used when the request method does not match the route.
	MethodNotAllowedHandler http.Handler

	// Routes to be matched, in order.
	routes []*Route

	// Routes by name for URL building.
	namedRoutes map[string]*Route

	// If true, do not clear the request context after handling the request.
	//
	// Deprecated: No effect, since the context is stored on the request itself.
	KeepContext bool

	// Slice of middlewares to be called after a match is found
	middlewares []middleware

	// configuration shared with `Route`
	routeConf
}




type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0

	// Header maps header keys to values. If the response had multiple
	// headers with the same key, they may be concatenated, with comma
	// delimiters.  (RFC 7230, section 3.2.2 requires that multiple headers
	// be semantically equivalent to a comma-delimited sequence.) When
	// Header values are duplicated by other fields in this struct (e.g.,
	// ContentLength, TransferEncoding, Trailer), the field values are
	// authoritative.
	//
	// Keys in the map are canonicalized (see CanonicalHeaderKey).
	Header Header

	// Body represents the response body.
	//
	// The response body is streamed on demand as the Body field
	// is read. If the network connection fails or the server
	// terminates the response, Body.Read calls return an error.
	//
	// The http Client and Transport guarantee that Body is always
	// non-nil, even on responses without a body or responses with
	// a zero-length body. It is the caller's responsibility to
	// close Body. The default HTTP client's Transport may not
	// reuse HTTP/1.x "keep-alive" TCP connections if the Body is
	// not read to completion and closed.
	//
	// The Body is automatically dechunked if the server replied
	// with a "chunked" Transfer-Encoding.
	//
	// As of Go 1.12, the Body will also implement io.Writer
	// on a successful "101 Switching Protocols" response,
	// as used by WebSockets and HTTP/2's "h2c" mode.
	Body io.ReadCloser

	// ContentLength records the length of the associated content. The
	// value -1 indicates that the length is unknown. Unless Request.Method
	// is "HEAD", values >= 0 indicate that the given number of bytes may
	// be read from Body.
	ContentLength int64

	// Contains transfer encodings from outer-most to inner-most. Value is
	// nil, means that "identity" encoding is used.
	TransferEncoding []string

	// Close records whether the header directed that the connection be
	// closed after reading Body. The value is advice for clients: neither
	// ReadResponse nor Response.Write ever closes a connection.
	Close bool

	// Uncompressed reports whether the response was sent compressed but
	// was decompressed by the http package. When true, reading from
	// Body yields the uncompressed content instead of the compressed
	// content actually set from the server, ContentLength is set to -1,
	// and the "Content-Length" and "Content-Encoding" fields are deleted
	// from the responseHeader. To get the original response from
	// the server, set Transport.DisableCompression to true.
	Uncompressed bool

	// Trailer maps trailer keys to values in the same
	// format as Header.
	//
	// The Trailer initially contains only nil values, one for
	// each key specified in the server's "Trailer" header
	// value. Those values are not added to Header.
	//
	// Trailer must not be accessed concurrently with Read calls
	// on the Body.
	//
	// After Body.Read has returned io.EOF, Trailer will contain
	// any trailer values sent by the server.
	Trailer Header

	// Request is the request that was sent to obtain this Response.
	// Request's Body is nil (having already been consumed).
	// This is only populated for Client requests.
	Request *Request

	// TLS contains information about the TLS connection on which the
	// response was received. It is nil for unencrypted responses.
	// The pointer is shared between responses and should not be
	// modified.
	TLS *tls.ConnectionState
}

type ResponseRecorder struct {
	// Code is the HTTP response code set by WriteHeader.
	//
	// Note that if a Handler never calls WriteHeader or Write,
	// this might end up being 0, rather than the implicit
	// http.StatusOK. To get the implicit value, use the Result
	// method.
	Code int

	// HeaderMap contains the headers explicitly set by the Handler.
	// It is an internal detail.
	//
	// Deprecated: HeaderMap exists for historical compatibility
	// and should not be used. To access the headers returned by a handler,
	// use the Response.Header map as returned by the Result method.
	HeaderMap http.Header

	// Body is the buffer to which the Handler's Write calls are sent.
	// If nil, the Writes are silently discarded.
	Body *bytes.Buffer

	// Flushed is whether the Handler called Flush.
	Flushed bool

	result      *http.Response // cache of Result's return value
	snapHeader  http.Header    // snapshot of HeaderMap at first Write
	wroteHeader bool
}

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
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

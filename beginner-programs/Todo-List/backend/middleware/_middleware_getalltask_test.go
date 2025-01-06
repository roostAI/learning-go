package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)





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
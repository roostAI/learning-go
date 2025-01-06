package middleware

import (
	"context"
	"fmt"
	"testing"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCollection struct {
	mock.Mock
}


type Call struct {
	Parent *Mock

	// The name of the method that was or will be called.
	Method string

	// Holds the arguments of the method.
	Arguments Arguments

	// Holds the arguments that should be returned when
	// this method is called.
	ReturnArguments Arguments

	// Holds the caller info for the On() call
	callerInfo []string

	// The number of times to return the return arguments when setting
	// expectations. 0 means to always return the value.
	Repeatability int

	// Amount of times this call has been called
	totalCalls int

	// Call to this method can be optional
	optional bool

	// Holds a channel that will be used to block the Return until it either
	// receives a message or is closed. nil means it returns immediately.
	WaitFor <-chan time.Time

	waitTime time.Duration

	// Holds a handler used to manipulate arguments content that are passed by
	// reference. It's useful when mocking methods such as unmarshalers or
	// decoders.
	RunFn func(Arguments)

	// PanicMsg holds msg to be used to mock panic on the function call
	//  if the PanicMsg is set to a non nil string the function call will panic
	// irrespective of other settings
	PanicMsg *string
}

type Mock struct {
	// Represents the calls that are expected of
	// an object.
	ExpectedCalls []*Call

	// Holds the calls that were made to this mocked object.
	Calls []Call

	// test is An optional variable that holds the test struct, to be used when an
	// invalid mock call was made.
	test TestingT

	// TestData holds any data that might be useful for testing.  Testify ignores
	// this data completely allowing you to do whatever you like with it.
	testData objx.Map

	mutex sync.Mutex
}




type Cursor struct {
	// Current contains the BSON bytes of the current change document. This property is only valid until the next call
	// to Next or TryNext. If continued access is required, a copy must be made.
	Current bson.Raw

	bc            batchCursor
	batch         *bsoncore.DocumentSequence
	batchLength   int
	registry      *bsoncodec.Registry
	clientSession *session.Client

	err error
}

type FindOptions struct {
	// AllowDiskUse specifies whether the server can write temporary data to disk while executing the Find operation.
	// This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.2 will report an error if this option
	// is specified. For server versions < 3.2, the driver will return a client-side error if this option is specified.
	// The default value is false.
	AllowDiskUse *bool

	// AllowPartial results specifies whether the Find operation on a sharded cluster can return partial results if some
	// shards are down rather than returning an error. The default value is false.
	AllowPartialResults *bool

	// BatchSize is the maximum number of documents to be included in each batch returned by the server.
	BatchSize *int32

	// Collation specifies a collation to use for string comparisons during the operation. This option is only valid for
	// MongoDB versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation

	// A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is nil, which means that no comment will be included in the logs.
	Comment *string

	// CursorType specifies the type of cursor that should be created for the operation. The default is NonTailable, which
	// means that the cursor will be closed by the server when the last batch of documents is retrieved.
	CursorType *CursorType

	// Hint is the index to use for the Find operation. This should either be the index name as a string or the index
	// specification as a document. The driver will return an error if the hint parameter is a multi-key map. The default
	// value is nil, which means that no hint will be sent.
	Hint interface{}

	// Limit is the maximum number of documents to return. The default value is 0, which means that all documents matching the
	// filter will be returned. A negative limit specifies that the resulting documents should be returned in a single
	// batch. The default value is 0.
	Limit *int64

	// Max is a document specifying the exclusive upper bound for a specific index. The default value is nil, which means that
	// there is no maximum value.
	Max interface{}

	// MaxAwaitTime is the maximum amount of time that the server should wait for new documents to satisfy a tailable cursor
	// query. This option is only valid for tailable await cursors (see the CursorType option for more information) and
	// MongoDB versions >= 3.2. For other cursor types or previous server versions, this option is ignored.
	MaxAwaitTime *time.Duration

	// MaxTime is the maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	//
	// NOTE(benjirewis): MaxTime will be deprecated in a future release. The more general Timeout option may be used in its
	// place to control the amount of time that a single operation can run before returning an error. MaxTime is ignored if
	// Timeout is set on the client.
	MaxTime *time.Duration

	// Min is a document specifying the inclusive lower bound for a specific index. The default value is 0, which means that
	// there is no minimum value.
	Min interface{}

	// NoCursorTimeout specifies whether the cursor created by the operation will not timeout after a period of inactivity.
	// The default value is false.
	NoCursorTimeout *bool

	// OplogReplay is for internal replication use only and should not be set.
	//
	// Deprecated: This option has been deprecated in MongoDB version 4.4 and will be ignored by the server if it is
	// set.
	OplogReplay *bool

	// Project is a document describing which fields will be included in the documents returned by the Find operation. The
	// default value is nil, which means all fields will be included.
	Projection interface{}

	// ReturnKey specifies whether the documents returned by the Find operation will only contain fields corresponding to the
	// index used. The default value is false.
	ReturnKey *bool

	// ShowRecordID specifies whether a $recordId field with a record identifier will be included in the documents returned by
	// the Find operation. The default value is false.
	ShowRecordID *bool

	// Skip is the number of documents to skip before adding documents to the result. The default value is 0.
	Skip *int64

	// Snapshot specifies whether the cursor will not return a document more than once because of an intervening write operation.
	// The default value is false.
	//
	// Deprecated: This option has been deprecated in MongoDB version 3.6 and removed in MongoDB version 4.0.
	Snapshot *bool

	// Sort is a document specifying the order in which documents should be returned.  The driver will return an error if the
	// sort parameter is a multi-key map.
	Sort interface{}

	// Let specifies parameters for the find expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let interface{}
}


type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
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

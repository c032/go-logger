package logger

// Buffered is a logger that might not persist the logged messages immediately,
// e.g. because doing so may be too inefficient.
//
// The `Flush` method must be called from time to time, to ensure all pending
// writes are persisted.
type Buffered interface {
	Logger

	// Flush writes any pending messages.
	Flush() error
}

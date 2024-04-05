package errorutil

// Causer represents this error was from other error
type Causer interface {
	// Cause returns the cause error
	Cause() error
}

package log

import "testing"

func TestLogging(t *testing.T) {
	worker := NewWorker()
	worker.doWork()
}

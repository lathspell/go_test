package log

import (
	"log"
	"os"
)

type Worker struct {
	log *log.Logger
}

func NewWorker() *Worker {
	// Time, Date and UTC can be configured using flags.
	// Custom formatting seems not to be supported.
	return &Worker{log.New(os.Stderr, "worker: ", log.LstdFlags)}
}

func (l *Worker) doWork() {
	l.log.Printf("doing work")
}


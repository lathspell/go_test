package lang

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// "select" will block until one of the channels is ready to be read/written to
func multichannelProducer(out, quit chan int) {
	for {
		select {
		case out <- 1:
			println("Wrote 1 to out channel")
		case <-quit:
			println("Read stop signal from quit channel")
			return
		default:
			println("All channels are blocked, waiting...")
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func TestChannelsSelect(t *testing.T) {
	out := make(chan int)
	quit := make(chan int)

	go multichannelProducer(out, quit)

	sum := 0
	sum += <-out
	sum += <-out
	sum += <-out
	quit <- 1

	assert.Equal(t, 3, sum)
}

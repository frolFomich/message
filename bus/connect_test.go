package bus

import (
	"fmt"
	"github.com/frolFomich/message"
	"sync"
	"testing"
	"time"
)

const (
	testStream = "TEST"
	testSubject = "TEST.testing"
)

func TestBus(t *testing.T) {
	err := AddStream(testStream, testSubject)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	Subscribe(testSubject, func(msg message.Message) bool {
		wg.Done()
		fmt.Printf("Received message: %s\n", msg)
		return true
	})
	StartSubscriptions("TESTING")
	err = PublishMessage(
		testSubject,
		message.New(
			message.Id("test-id-1"),
			message.Type("/test/type"),
			message.Timestamp(time.Now()),
			message.DataType("application/Json"),
			message.Data(map[string]interface{}{
				"A": "B",
				"C": 100.0,
				"D": true,
			})))
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	if waitTimeout(&wg, 10 * time.Second) {
		t.Errorf("Timeout while waiting message")
		return
	}
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})

	go func() {
		defer close(c)
		wg.Wait()
	}()

	select {
	case <- c:
		return false
	case <-time.After(timeout):
		return true
	}
}
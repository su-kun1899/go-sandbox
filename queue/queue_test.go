package queue_test

import (
	"testing"
	"github.com/su-kun1899/go-sandbox/queue"
)

func Test_Queue_Enqueue(t *testing.T) {
	// given
	q := queue.NewQueue(5)
	arg := "Hello, queue!"

	// when
	q.Enqueue(arg)

	// then
	element, err := q.Dequeue()
	if err != nil {
		t.Error("unexpected error:", err)
	}
	if element != arg {
		t.Errorf("Dequeue want %s but got %s", arg, element)
	}
}

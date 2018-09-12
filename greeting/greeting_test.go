package greeting_test

import (
	"testing"
	"github.com/su-kun1899/go-sandbox/greeting"
	"time"
	"bytes"
)

func TestGreeting_Do(t *testing.T) {
	g := greeting.Greeting{
		Clock: greeting.ClockFunc(func() time.Time {
			return time.Date(2018, 8, 31, 6, 0, 0, 0, time.Local)
		}),
	}

	var buf bytes.Buffer
	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error:", err)
	}

	if expected, actual := "おはよう", buf.String(); expected != actual {
		t.Errorf("greeting message want %s but got %s", expected, actual)
	}
}

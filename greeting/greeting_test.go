package greeting_test

import (
	"testing"
	"github.com/su-kun1899/go-sandbox/greeting"
	"time"
	"bytes"
)

func TestGreeting_Do(t *testing.T) {
	g := greeting.Greeting{
		Clock: mockClock(t, "2018/08/31 06:00:00"),
	}

	var buf bytes.Buffer
	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error:", err)
	}

	if expected, actual := "おはよう", buf.String(); expected != actual {
		t.Errorf("greeting message want %s but got %s", expected, actual)
	}
}

func mockClock(t *testing.T, v string) greeting.Clock {
	t.Helper()
	now, err := time.Parse("2006/01/02 15:04:05", v)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	return greeting.ClockFunc(func() time.Time {
		return now
	})
}

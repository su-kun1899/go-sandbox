package greeting_test

import (
	"testing"
	"github.com/su-kun1899/go-sandbox/greeting"
	"time"
	"bytes"
	"io"
)

func TestGreeting_Do(t *testing.T) {
	cases := map[string]struct {
		writer io.Writer
		clock  greeting.Clock

		msg       string
		expectErr bool
	}{
		"04時": {
			writer: new(bytes.Buffer),
			clock:  mockClock(t, "2018/08/31 04:00:00"),
			msg:    "おはよう",
		},
		"09時": {
			writer: new(bytes.Buffer),
			clock:  mockClock(t, "2018/08/31 09:00:00"),
			msg:    "おはよう",
		},
		"10時": {
			writer: new(bytes.Buffer),
			clock:  mockClock(t, "2018/08/31 10:00:00"),
			msg:    "こんにちは",
		},
		"16時": {
			writer: new(bytes.Buffer),
			clock:  mockClock(t, "2018/08/31 16:00:00"),
			msg:    "こんにちは",
		},
		"17時": {
			writer: new(bytes.Buffer),
			clock:  mockClock(t, "2018/08/31 17:00:00"),
			msg:    "こんばんは",
		},
		"03時": {
			writer: new(bytes.Buffer),
			clock:  mockClock(t, "2018/08/31 03:00:00"),
			msg:    "こんばんは",
		},
	}

	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			g := greeting.Greeting{
				Clock: tc.clock,
			}

			switch err := g.Do(tc.writer); true {
			// エラーを期待しているのにエラーが発生していない
			case err == nil && tc.expectErr:
				t.Error("expected error did not occur")
				// エラーは期待していないのにエラーが発生した
			case err != nil && !tc.expectErr:
				t.Error("unexpected error:", err)
			}

			// tc.writerが*byte.Bufferだったら値もチェック
			if buff, ok := tc.writer.(*bytes.Buffer); ok {
				msg := buff.String()
				if msg != tc.msg {
					t.Errorf("greeting message want %s but got %s", tc.msg, msg)
				}
			}
		})
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

package binance

import (
	"net/url"
	"testing"
	"time"
)

func TestNewQuery(t *testing.T) {
	cases := []struct {
		options  []QueryFunc
		expected string
	}{
		{[]QueryFunc{Limit(10)}, "limit=10"},
		{[]QueryFunc{Limit(10), FromID(10)}, "fromId=10&limit=10"},
		{[]QueryFunc{FromID(10)}, "fromId=10"},
		{[]QueryFunc{FromID(10), StartTime(FromTime(time.Time{}))}, "fromId=10&startTime=-6795364578871"},
		{[]QueryFunc{EndTime(FromTime(time.Time{}))}, "endTime=-6795364578871"},
	}

	for i, c := range cases {
		q := newQuery(c.options)
		v := url.Values{}
		q.params()(v)
		result := v.Encode()
		if result != c.expected {
			t.Errorf("case %d encoded as '%s', expected '%s'", i, result, c.expected)
		}
	}
}

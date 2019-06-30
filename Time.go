package binance

import (
	"strconv"
	"time"
)

// Time is a type that matches Binance's milliseonc precision timestamps. It's
// embedding time.Time so all the usual time.Time methods should work as
// expected.
type Time struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler while reading a
// Javascript-style millisecond timestamp.
func (t *Time) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		// Try de decode as a regular timestamp. Sigh.
		return t.Time.UnmarshalJSON(data)
	}

	sec := i / 1000
	usec := (i % 1000) * 1000 * 1000

	// We don't know the timezone. Let's assume UTC, the user of this package
	// can change it if she desires.
	t.Time = time.Unix(sec, usec).UTC()

	return nil
}

// MarshalJSON implements json.Marshaler.
func (t Time) MarshalJSON() ([]byte, error) {
	// Get timestamp in milliseconds.
	ms := t.Time.UnixNano() / (int64(time.Millisecond))

	// Convert to string.
	str := strconv.FormatInt(ms, 10)

	// Return as bytes for marshaler.
	return []byte(str), nil
}

// FromTime will take a regular Go timestamp and convert it to a type suitable
// for Binance API consumption.
func FromTime(t time.Time) Time {
	return Time{
		Time: t,
	}
}

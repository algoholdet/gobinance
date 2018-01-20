package binance

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimeJSON(t *testing.T) {
	cases := []string{
		"0",
		"1516408822000",
		"123456789",
	}

	for _, c := range cases {
		tim := Time{}
		err := json.Unmarshal([]byte(c), &tim)
		if err != nil {
			t.Errorf("Failed to unmarshal %s", c)
		}

		out, err := json.Marshal(&tim)
		if err != nil {
			t.Errorf("Failed to marshal %v", tim.Time)
		}

		if string(out) != c {
			t.Errorf("%s failed JSON-roundtrip test. Got '%s' back, time is %s", c, string(out), tim.Time.String())
		}

		err = json.Unmarshal([]byte("\""+tim.Format(time.RFC3339Nano)+"\""), &tim)
		if err != nil {
			t.Errorf("Failed to unmarshal regular Go timestamp: %s", err.Error())
		}
	}
}

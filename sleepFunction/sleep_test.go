package sleepFunction

import "testing"

func TestSleepFunction(t *testing.T) {

	time := Sleep_UsingTimeAfter(10)

	if time != 10 {
		t.Error("Expected 10, got", time)
	}

}

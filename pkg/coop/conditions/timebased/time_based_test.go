package timebased

import (
	"testing"
	"time"
)

const latitude = 43.525776
const longitude = 1.327727

func TestGetOpeningTime(t *testing.T) {
	tbc := NewTimeBasedCondition(8, 0)

	if tbc.GetOpeningTime() != time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 0, 0, 0, time.Local) {
		t.Errorf("Time is incorrect ! Should be : %s. It is : %s", tbc.GetOpeningTime(), time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 0, 0, 0, time.Local))
		t.Fail()
	}
}

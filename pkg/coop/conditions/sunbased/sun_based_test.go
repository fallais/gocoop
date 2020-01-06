package sunbased

import (
	"testing"
	"time"

	"github.com/cpucycle/astrotime"
)

const latitude = 43.525776
const longitude = 1.327727

func TestGetOpeningTime(t *testing.T) {
	offset := 45 * time.Minute
	sbc := NewSunBasedCondition(offset, latitude, longitude)
	sunrise := astrotime.CalcSunrise(time.Now(), latitude, longitude)

	if sbc.GetOpeningTime() != sunrise.Add(offset) {
		t.Errorf("Time is incorrect ! Should be : %s. It is : %s", sbc.GetOpeningTime(), sunrise.Add(offset))
		t.Fail()
	}

	offset = -45 * time.Minute
	sbc = NewSunBasedCondition(offset, latitude, longitude)
	sunrise = astrotime.CalcSunrise(time.Now(), latitude, longitude)

	if sbc.GetOpeningTime() != sunrise.Add(offset) {
		t.Errorf("Time is incorrect ! Should be : %s. It is : %s", sbc.GetOpeningTime(), sunrise.Add(offset))
		t.Fail()
	}
}

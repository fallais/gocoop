package conditions

import (
	"testing"
	"time"

	"github.com/cpucycle/astrotime"
)

const latitude = 43.525776
const longitude = 1.327727

func TestGetTime(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Paris")
	now := time.Now().In(loc)

	offset := 45 * time.Minute
	sbc := NewSunBasedCondition(offset, latitude, longitude, loc)
	sunrise := astrotime.CalcSunrise(now, latitude, longitude)

	if sbc.GetTime() != sunrise.Add(offset) {
		t.Errorf("Time is incorrect ! Should be : %s. It is : %s", sbc.GetTime(), sunrise.Add(offset))
		t.Fail()
	}

	offset = -45 * time.Minute
	sbc = NewSunBasedCondition(offset, latitude, longitude, loc)
	sunrise = astrotime.CalcSunrise(now, latitude, longitude)

	if sbc.GetTime() != sunrise.Add(offset) {
		t.Errorf("Time is incorrect ! Should be : %s. It is : %s", sbc.GetTime(), sunrise.Add(offset))
		t.Fail()
	}
}

func TestGetNextTime(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Paris")
	now := time.Now().In(loc)

	offset := 45 * time.Minute
	sbc := NewSunBasedCondition(offset, latitude, longitude, loc)
	sunrise := astrotime.CalcSunrise(now, latitude, longitude)

	if sbc.GetNextTime() != sunrise.Add(offset).AddDate(0, 0, 1) {
		t.Errorf("Time is incorrect ! Should be : %s. It is : %s", sbc.GetTime(), sunrise.Add(offset).AddDate(0, 0, 1))
		t.Fail()
	}
}

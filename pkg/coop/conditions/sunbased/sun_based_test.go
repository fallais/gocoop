package sunbased

import (
	"testing"
	"time"

	"github.com/cpucycle/astrotime"
)

const latitude = 43.525776
const longitude = 1.327727

func TestSunBasedCondition(t *testing.T) {
	_, err := NewSunBasedCondition("faya", latitude, longitude)
	if err == nil {
		t.Fatal("should error")
	}
}

func TestGetOpeningTime(t *testing.T) {
	sbc, err := NewSunBasedCondition("45m", latitude, longitude)
	if err != nil {
		t.Fatalf("should not error")
	}
	sunrise := astrotime.CalcSunrise(time.Now(), latitude, longitude)

	if sbc.OpeningTime() != sunrise.Add(45*time.Minute) {
		t.Fatalf("Time is incorrect ! Should be : %s. It is : %s", sbc.OpeningTime(), sunrise.Add(45*time.Minute))
	}
}

func TestGetClosingTime(t *testing.T) {
	sbc, err := NewSunBasedCondition("-45m", latitude, longitude)
	if err != nil {
		t.Fatalf("should not error")
	}
	sunset := astrotime.CalcSunset(time.Now(), latitude, longitude)

	if sbc.ClosingTime() != sunset.Add(-45*time.Minute) {
		t.Fatalf("Time is incorrect ! Should be : %s. It is : %s", sbc.ClosingTime(), sunset.Add(-45*time.Minute))
	}
}

package timebased

import (
	"testing"
	"time"
)

func TestTimeBasedCondition(t *testing.T) {
	_, err := NewTimeBasedCondition("0888h00")
	if err == nil {
		t.Fatal("should error")
	}

	tbc, _ := NewTimeBasedCondition("08h00")
	if tbc.Mode() != "time_based" {
		t.Fatalf("should be time_based, it is %s", tbc.Mode())
	}
	if tbc.Value() != "08h00" {
		t.Fatalf("should be 08h00, it is %s", tbc.Value())
	}
}

func TestGetOpeningTime(t *testing.T) {
	tbc, err := NewTimeBasedCondition("08h00")
	if err != nil {
		t.Fatal("should not error")
	}

	if tbc.OpeningTime() != time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 0, 0, 0, time.Local) {
		t.Fatalf("Time is incorrect ! Should be : %s. It is : %s", tbc.OpeningTime(), time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 0, 0, 0, time.Local))
	}
}

func TestGetClosingTime(t *testing.T) {
	tbc, err := NewTimeBasedCondition("18h30")
	if err != nil {
		t.Fatal("should not error")
	}

	if tbc.ClosingTime() != time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 18, 30, 0, 0, time.Local) {
		t.Fatalf("Time is incorrect ! Should be : %s. It is : %s", tbc.ClosingTime(), time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 18, 30, 0, 0, time.Local))
	}
}

package timebased

import "testing"

func TestParseTime(t *testing.T) {
	t1 := "DDhEE"
	h, m, err := parseTime(t1)
	if err == nil {
		t.Errorf("Should raise an error")
		t.Fail()
	}

	t2 := "08h30"
	h, m, err = parseTime(t2)
	if err != nil {
		t.Fatal("Should not raise an error")
	}
	if h != 8 {
		t.Fatal("Hours should be 8")
	}
	if m != 30 {
		t.Fatal("Minutes should be 30")
	}

	t4 := "754h845"
	h, m, err = parseTime(t4)
	if err == nil {
		t.Errorf("Should raise an error")
		t.Fail()
	}

	t5 := "26h00"
	h, m, err = parseTime(t5)
	if err == nil {
		t.Errorf("Should raise an error")
		t.Fail()
	}

	t6 := "15h72"
	h, m, err = parseTime(t6)
	if err == nil {
		t.Errorf("Should raise an error")
		t.Fail()
	}
}

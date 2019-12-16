package join

import "testing"

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple, and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Error("Expected ", want, "but got", got)
	}
}

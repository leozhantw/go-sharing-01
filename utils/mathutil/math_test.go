package mathutil

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		giveA, giveB, want int
	}{
		{1, 2 ,3},
		{0, 1, 1},
	}
	for _, c := range cases {
		got := Add(c.giveA, c.giveB)
		if got != c.want {
			t.Errorf("Add(%d, %d) == %d, want %d", c.giveA, c.giveB, got, c.want)
		}
	}
}
package stringutil

import "testing"

func TestGetCopy(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"dicky", "dicky"},
	}
	for _, val := range cases {
		got := GetCopy(val.in)
		if got != val.want {
			t.Errorf("wrong answer expect %v got %v\n", val.want, got)
		}
	}
}

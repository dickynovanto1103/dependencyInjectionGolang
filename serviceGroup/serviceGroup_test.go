package serviceGroup

import (
	"testing"
)

func TestGetAll(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"d", "a"}, "Get info from url dGet info from url a"},
	}
	serviceGroup := CreateServiceGroup()
	for _, val := range cases {
		got := serviceGroup.GetAll(val.in...)
		if got != val.want {
			t.Errorf("wrong answer got %v want %v\n", got, val.want)
		}
	}
}

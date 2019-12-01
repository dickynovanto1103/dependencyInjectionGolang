package httpClient

import (
	"testing"

	"github.com/dickynovanto1103/dependencyInjectionGolang/logger"
)

func TestGet(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"https://google.com", "Get info from url https://google.com"},
		{"", "Get info from url "},
		{"sadf", "Get info from url sadf"},
	}
	logger := &logger.Logger{}
	httpClient := NewHttpClient(logger)
	for _, val := range cases {
		got := httpClient.Get(val.in)
		if got != val.want {
			t.Errorf("wrong answer got %v want %v\n", got, val.want)
		}
	}
}

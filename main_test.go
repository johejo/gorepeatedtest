package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		args []string
	}{
		{args: []string{"--", "-v", "testing"}},
		{args: []string{"--", "testing"}},
		{args: []string{"-v", "--", "testing"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if err := run(tt.args); err != nil {
				t.Fatal(err)
			}
		})
	}
}

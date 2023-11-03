package gointtoletters

import (
	"testing"
)

func TestIntToLetters(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want string
	}{
		{
			name: "correct first letter",
			arg:  1,
			want: "A",
		},
		{
			name: "two letters result",
			arg:  27,
			want: "AA",
		},
		{
			name: "three letters result",
			arg:  16158,
			want: "WWL",
		},
		{
			name: "five letters result",
			arg:  787428,
			want: "ARTUR",
		},
		{
			name: "six letters result",
			arg:  100011111111,
			want: "CYFSHNC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToLetters(tt.arg)
			if got != tt.want {
				t.Errorf("IntToLetters got= %v, want %v", got, tt.want)
			}
		})
	}
}

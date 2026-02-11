package main_test

import (
	main "example/test"
	"testing"
)

func TestTest1(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		strNum  string
		want    int
		wantErr bool
	}{
		{
			name:    "example test 1",
			strNum:  "16",
			want:    7,
			wantErr: false,
		},
		{
			name:    "example test 2",
			strNum:  "942",
			want:    6,
			wantErr: false,
		},
		{
			name:    "example test 3",
			strNum:  "132189",
			want:    6,
			wantErr: false,
		},
		{
			name:    "example test 4",
			strNum:  "493193",
			want:    2,
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := main.Test1(tt.strNum)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Test1() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Test1() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Test1() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_getPassword(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		grid    [3][3]string
		command []string
		want    string
		wantErr bool
	}{
		{
			name: "case 1",
			grid: [3][3]string{
				{"a", "x", "c"}, // row index 0
				{"g", "l", "t"}, // row index 1
				{"o", "v", "e"}, // row index 2
			},
			command: []string{"downT", "down", "leftT", "rightT", "rightT", "upT"},
			want:    "lovet",
			wantErr: false,
		},
		{
			name: "case 2",
			grid: [3][3]string{
				{"p", "x", "m"}, // row index 0
				{"a", "$", "$"}, // row index 1
				{"k", "i", "t"}, // row index 2
			},
			command: []string{"leftT", "downT", "rightT", "rightT"},
			want:    "pa$$",
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := getPassword(tt.grid, tt.command)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("getPassword() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("getPassword() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("getPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

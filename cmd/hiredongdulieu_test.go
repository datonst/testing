package main

import (
	"testing"

	"github.com/datonst/testing_assignment/internal/assert"
)

func TestHireEmployee_DongDK(t *testing.T) {
	// Create a slice of anonymous structs containing the test case name,
	// input to our hireEmployee() function (the t field, the deal field), and expected output (the want field).
	tests := []struct {
		name string
		t    int
		deal int
		want string
	}{
		{
			name: "1",
			t:    -1,
			deal: 10,
			want: "Invalid",
		},
		{
			name: "2",
			t:    65,
			deal: 10,
			want: "Không thuê",
		},
		{
			name: "3",
			t:    18,
			deal: 800,
			want: "Thuê toàn thời gian",
		},
		{
			name: "4",
			t:    65,
			deal: 10,
			want: "Không thuê",
		},
		{
			name: "5",
			t:    20,
			deal: 2000,
			want: "Không thuê",
		},
		{
			name: "6",
			t:    65,
			deal: 10,
			want: "Không thuê",
		},
		{
			name: "7",
			t:    17,
			deal: 100,
			want: "Thuê bán thời gian",
		},
		{
			name: "8",
			t:    18,
			deal: 800,
			want: "Thuê toàn thời gian",
		},
		{
			name: "9",
			t:    20,
			deal: 100,
			want: "Thuê toàn thời gian",
		},
		{
			name: "10",
			t:    20,
			deal: 2000,
			want: "Không thuê",
		},
		{
			name: "11",
			t:    65,
			deal: 10,
			want: "Không thuê",
		},
		{
			name: "12",
			t:    -1,
			deal: 10,
			want: "Invalid",
		},
		{
			name: "13",
			t:    17,
			deal: 100,
			want: "Thuê bán thời gian",
		},
		{
			name: "14",
			t:    20,
			deal: 100,
			want: "Thuê toàn thời gian",
		},
	}
	// Loop over the test cases.
	for _, tt := range tests {
		// Use the t.Run() function to run a sub-test for each test case. The
		// first parameter to this is the name of the test (which is used to
		// identify the sub-test in any log output) and the second parameter is
		// and anonymous function containing the actual test for each case.
		t.Run(tt.name, func(t *testing.T) {
			hd := hireEmployee(tt.t, tt.deal)
			// t.Errorf("hd: %q\n", hd)
			assert.Equal(t, hd, tt.want)
		})
	}
}

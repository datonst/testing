package main

import (
	"testing"

	"github.com/datonst/testing_assignment/internal/assert"
)

func TestHireEmployee(t *testing.T) {
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
			t:    0,
			deal: 50,
			want: "Invalid",
		},
		{
			name: "2",
			t:    16,
			deal: 350,
			want: "Thuê bán thời gian",
		},
		{
			name: "3",
			t:    16,
			deal: 600,
			want: "Không thuê",
		},
		{
			name: "4",
			t:    20,
			deal: 700,
			want: "Thuê toàn thời gian",
		},
		{
			name: "5",
			t:    20,
			deal: 1200,
			want: "Không thuê",
		},
		{
			name: "6",
			t:    56,
			deal: 400,
			want: "Không thuê",
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

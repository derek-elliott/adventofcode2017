package handlers

import "testing"

type DaySixTestCase struct {
	Data    []int
	Target  int
	Message string
}

type DaySixTestCases []DaySixTestCase

var cases = DaySixTestCases{
	DaySixTestCase{
		[]int{0, 2, 7, 0},
		5,
		"Expected 5",
	},
}

func TestDaySix(t *testing.T) {
	for _, test := range cases {
		testValue := DaySix(test.Data)
		if testValue != test.Target {
			t.Error("Test Failed! Returned: ", testValue, test.Message)
		}
	}
}

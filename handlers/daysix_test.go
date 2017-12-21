package handlers

import "testing"

type DaySixTestCase struct {
	Data    []int
	Target  []int
	Message string
}

type DaySixTestCases []DaySixTestCase

var daySixCases = DaySixTestCases{
	DaySixTestCase{
		[]int{0, 2, 7, 0},
		[]int{5, 4},
		"Expected 5, 4",
	},
}

func TestDaySix(t *testing.T) {
	for _, test := range daySixCases {
		testValue1, testValue2 := DaySix(test.Data)
		if testValue1 != test.Target[0] || testValue2 != test.Target[1] {
			t.Errorf("Test Failed! Returned: (%d, %d), %s", testValue1, testValue2, test.Message)
		}
	}
}

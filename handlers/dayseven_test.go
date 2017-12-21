package handlers

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	log "github.com/sirupsen/logrus"
)

type DaySevenTestCase struct {
	Data    DaySevenInputs
	Target  string
	Message string
}

type DaySevenTestCases []DaySevenTestCase

func TestDaySevenPartOne(t *testing.T) {
	var daySevenCases DaySevenTestCases
	testFile, err := ioutil.ReadFile("../testdata/dayseven.json")
	if err != nil {
		log.WithError(err).Fatal("Error reading file")
	}
	var input DaySevenInputs
	if err = json.Unmarshal(testFile, &input); err != nil {
		log.WithError(err).Fatal("Error deserializing test data")
	}
	daySevenCases = DaySevenTestCases{
		DaySevenTestCase{
			input,
			"tknk",
			"Expected tknk",
		},
	}

	for _, test := range daySevenCases {
		testValue := DaySevenPartOne(test.Data)
		if testValue != test.Target {
			t.Errorf("Test failed!  Returned: %s, %s", testValue, test.Message)
		}
	}
}

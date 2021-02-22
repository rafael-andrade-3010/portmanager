package api

import (
	"strings"
	"testing"
)

func TestParseJson(t *testing.T) {
	body := `{
	  "AEAJM": {
		"name": "Ajman",
		"city": "Ajman",
		"country": "United Arab Emirates",
		"alias": [],
		"regions": [],
		"coordinates": [
		  55.5136433,
		  25.4052165
		],
		"province": "Ajman",
		"timezone": "Asia/Dubai",
		"unlocs": [
		  "AEAJM"
		],
		"code": "52000"
	  },
	  "AEAUH": {
		"name": "Abu Dhabi",
		"coordinates": [
		  54.37,
		  24.47
		],
		"city": "Abu Dhabi",
		"province": "Abu Z¸aby [Abu Dhabi]",
		"country": "United Arab Emirates",
		"alias": [],
		"regions": [],
		"timezone": "Asia/Dubai",
		"unlocs": [
		  "AEAUH"
		],
		"code": "52001"
	  }
	}`
	ports, err := parseJson(strings.NewReader(body))
	if err != nil {
		t.Errorf("No error was expected here, got %v", err.Error())
	}
	if len(ports) != 2 {
		t.Errorf("Expected %v ports, got %v", 2, len(ports))
	}
	for i, p := range ports {
		if i == 0 {
			if p.Key != "AEAJM" {
				t.Errorf("Expected port with key %v got %v", "AEAJM", p.Key)
			}
			if p.Name != "Ajman" {
				t.Errorf("Expected port with name %v got %v", "Ajman", p.Name)
			}
		}
	}
}

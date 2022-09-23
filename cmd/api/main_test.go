package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestSuccessfulServerRequest(t *testing.T) {
	cfg := &config{
		4000, "development ",
	}
	go startServer(cfg)

	tests := []struct {
		name string
		url  string
		want []string
	}{
		{
			"request time without timezone",
			fmt.Sprintf("http://localhost:%d/api/time", cfg.port),
			[]string{"current_time"},
		},
		{
			"request time with timezone",
			fmt.Sprintf("http://localhost:%d/api/time?tz=America/New_York,Asia/Kolkata", cfg.port),
			[]string{"America/New_York", "Asia/Kolkata"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := http.Get(test.url)

			if err != nil {
				t.Errorf("test failed with error from endpoint for %s", test.name)
			}

			if result.StatusCode != http.StatusOK {
				t.Errorf("test failed to return successful http status code %d for %s", result.StatusCode, test.name)
			}

			data, err := io.ReadAll(result.Body)
			if err != nil {
				t.Errorf("test failed due to failure to read body for %s", test.name)
			}

			var responseObject map[string]string
			json.Unmarshal(data, &responseObject)

			if len(responseObject) <= 0 {
				t.Errorf("test failed with no response for %s", test.name)
			}

			for _, value := range test.want {
				if responseObject[value] == "" {
					t.Errorf("test failed with expected want : %s not found for %s", value, test.name)
				}
			}

		})
	}
}

func TestInvalidTimeZoneRequest(t *testing.T) {

	const testName = "invalid time zone"
	cfg := &config{
		4000, "development ",
	}
	go startServer(cfg)

	t.Run(testName, func(t *testing.T) {
		result, err := http.Get(fmt.Sprintf("http://localhost:%d/api/time?tz=America/New_York,Ajdshjdshjfsjolkata", cfg.port))

		if err != nil {
			t.Error("test failed with error from endpoint ")
		}

		if result.StatusCode != http.StatusNotFound {
			t.Errorf("test failed to return not found http status code but returned  %d ", result.StatusCode)
		}
	})
}

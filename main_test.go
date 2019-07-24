package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

var testDescriptions = []string{
	"Sampe Timer",
	"NoSpace",
	"",
	"1",
	"\n",
	"\\",
	"'",
	"}",
	",",
	"{},,.[]",
}

func TestPostingTimer(t *testing.T) {
	for _, tt := range testDescriptions {
		t.Run(tt, func(t *testing.T) {
			router := SetupRouter()
			ts := httptest.NewServer(router)
			store = make(map[uuid.UUID]*Timer)
			defer ts.Close()

			client := ts.Client()
			expected := tt

			payload := make(map[string]string)
			payload["description"] = expected
			body, err := json.Marshal(payload)

			if err != nil {
				panic(err)
			}

			bodyBuffer := bytes.NewBuffer(body)

			_, err = client.Post(ts.URL, "application/json", bodyBuffer)
			if err != nil {
				panic(err)
			}

			resp, err := client.Get(ts.URL + "/current")

			respBody, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				panic(err)
			}

			var currentTimer Timer

			err = json.Unmarshal(respBody, &currentTimer)

			if err != nil {
				panic(err)
			}

			got := currentTimer.Description

			if got != expected {
				t.Errorf("Wanted: %v, Got: %v\n", expected, got)
			}
		})
	}
}

func TestStartStopTimer(t *testing.T) {
	for _, tt := range testDescriptions {
		t.Run(tt, func(t *testing.T) {
			store = make(map[uuid.UUID]*Timer)
			router := SetupRouter()
			ts := httptest.NewServer(router)
			defer ts.Close()

			client := ts.Client()

			// Start timer

			expected := tt
			payload := make(map[string]string)
			payload["description"] = expected
			body, err := json.Marshal(payload)

			if err != nil {
				panic(err)
			}

			bodyBuffer := bytes.NewBuffer(body)

			_, err = client.Post(ts.URL, "application/json", bodyBuffer)
			if err != nil {
				panic(err)
			}

			// Stop timer
			var stoppedTimer Timer

			req, err := http.NewRequest("PUT", ts.URL+"/stop", nil)

			if err != nil {
				panic(err)
			}

			resp, err := client.Do(req)

			if err != nil {
				panic(err)
			}

			respBody, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				panic(err)
			}

			json.Unmarshal(respBody, &stoppedTimer)
			got := stoppedTimer.Description

			if got != expected {
				t.Errorf("Expected %v, Got: %v", expected, got)
			}
		})
	}
}

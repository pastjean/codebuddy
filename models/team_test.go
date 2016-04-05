package models

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestTeams(t *testing.T) {
	src := []byte("{\"Avenger\":[\"Ironman\",\"Batman\"],\"Parliment\":[\"Harper\",\"Trudeau\"]}")
	teams := &TeamsMap{}
	err := json.Unmarshal(src, teams)
	if err != nil {
		t.Error(err)
	}

	result, err := json.Marshal(teams)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(src, result) {
		t.Error(string(src), "!=", string(result))
	}
}

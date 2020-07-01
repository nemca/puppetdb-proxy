package main

import "encoding/json"

type v3EventCount struct {
	Subject     json.RawMessage `json:"subject"`
	SubjectType json.RawMessage `json:"subject_type"`
	Failures    json.RawMessage `json:"failures"`
	Successes   json.RawMessage `json:"successes"`
	Noops       json.RawMessage `json:"noops"`
	Skips       json.RawMessage `json:"skips"`
}

type v3EventCounts []v3EventCount

type aliasEventCount struct {
	Subject     json.RawMessage `json:"subject"`
	SubjectType json.RawMessage `json:"subject-type"`
	Failures    json.RawMessage `json:"failures"`
	Successes   json.RawMessage `json:"successes"`
	Noops       json.RawMessage `json:"noops"`
	Skips       json.RawMessage `json:"skips"`
}

type v3AggregateEventCount struct {
	Successes json.RawMessage `json:"successes"`
	Failures  json.RawMessage `json:"failures"`
	Noops     json.RawMessage `json:"noops"`
	Skips     json.RawMessage `json:"skips"`
	Total     json.RawMessage `json:"total"`
}

type v3AggregateEventCounts []v3AggregateEventCount

func (ec *v3EventCount) MarshalJSON() ([]byte, error) {
	var a = aliasEventCount(*ec)
	return json.Marshal(&a)
}

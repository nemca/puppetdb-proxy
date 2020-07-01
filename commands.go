package main

import "encoding/json"

type v3Commands struct {
	Command string          `json:"command"`
	Version int             `json:"version"`
	Payload json.RawMessage `json:"payload"`
}

type v3CommandsFact struct {
	Name              json.RawMessage `json:"name"`
	Environment       json.RawMessage `json:"environment"`
	Values            json.RawMessage `json:"values"`
	ProducerTimestamp string          `json:"producer_timestamp"`
	Producer          string          `json:"producer"`
}

type v4Commands struct {
	Command string          `json:"command"`
	Version int             `json:"version"`
	Payload json.RawMessage `json:"payload"`
}

type v4CommandsFact struct {
	Name              json.RawMessage `json:"certname"`
	Environment       json.RawMessage `json:"environment"`
	Values            json.RawMessage `json:"values"`
	ProducerTimestamp string          `json:"producer_timestamp"`
	Producer          string          `json:"producer"`
}

type v4CommandsDeacticate struct {
	Name              json.RawMessage `json:"certname"`
	ProducerTimestamp string          `json:"producer_timestamp"`
}

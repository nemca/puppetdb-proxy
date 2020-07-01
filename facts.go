package main

import (
	"encoding/json"
	"time"
)

type v3Facts struct {
	Certname    string          `json:"name"`
	Environment string          `json:"environment"`
	Values      json.RawMessage `json:"values"`
}

type v3FactsGet struct {
	Certname string          `json:"certname"`
	Name     string          `json:"name"`
	Value    json.RawMessage `json:"value"`
}

type v4Facts struct {
	Certname          string          `json:"certname"`
	Environment       string          `json:"environment"`
	Values            json.RawMessage `json:"values"`
	ProducerTimestamp string          `json:"producer_timestamp"`
	Producer          string          `json:"producer"`
}

func v3toV4FactsConv(v3f v3Facts) v4Facts {
	var v4f v4Facts
	v4f.Certname = v3f.Certname
	v4f.Environment = opts.Environment
	v4f.Producer = opts.Producer
	v4f.ProducerTimestamp = time.Now().Format(time.RFC3339)
	v4f.Values = v3f.Values

	return v4f
}

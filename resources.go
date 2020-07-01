package main

import "encoding/json"

type v3Resource struct {
	Certname   json.RawMessage `json:"certname"`
	Resource   json.RawMessage `json:"resource"`
	Type       json.RawMessage `json:"type"`
	Title      json.RawMessage `json:"title"`
	Tags       json.RawMessage `json:"tags"`
	File       json.RawMessage `json:"file"`
	Line       json.RawMessage `json:"line"`
	Parameters json.RawMessage `json:"parameters"`
	Exported   json.RawMessage `json:"exported"`
}

type v3Resources []v3Resource

package main

import "encoding/json"

type v3Node struct {
	Name             json.RawMessage `json:"certname"`
	Deactivated      json.RawMessage `json:"deactivated"`
	CatalogTimestamp json.RawMessage `json:"catalog_timestamp"`
	FactsTimestamp   json.RawMessage `json:"facts_timestamp"`
	ReportTimestamp  json.RawMessage `json:"report_timestamp"`
}

type v3Nodes []v3Node

type aliasNode struct {
	Name             json.RawMessage `json:"name"`
	Deactivated      json.RawMessage `json:"deactivated"`
	CatalogTimestamp json.RawMessage `json:"catalog_timestamp"`
	FactsTimestamp   json.RawMessage `json:"facts_timestamp"`
	ReportTimestamp  json.RawMessage `json:"report_timestamp"`
}

func (n *v3Node) MarshalJSON() ([]byte, error) {
	var a = aliasNode(*n)
	return json.Marshal(&a)
}

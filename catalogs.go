package main

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type v3Catalog struct {
	Name              string           `json:"name"`
	Version           string           `json:"version"`
	Environment       string           `json:"environment,omitempty"`
	TransactionUUID   string           `json:"transaction-uuid"`
	ProducerTimestamp string           `json:"producer-timestamp,omitempty"`
	Edges             catalogEdges     `json:"edges"`
	Resources         catalogResources `json:"resources"`
}

type catalogEdges []catalogEdge

type catalogEdge struct {
	Relationship string       `json:"relationship"`
	Source       resourceSpec `json:"source"`
	Target       resourceSpec `json:"target"`
}

type resourceSpec struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

type catalogResources []catalogResource

type catalogResource struct {
	Resource   string          `json:"resource"`
	Type       string          `json:"type"`
	Title      string          `json:"title"`
	Tags       []string        `json:"tags"`
	File       string          `json:"file"`
	Line       int             `json:"line"`
	Parameters json.RawMessage `json:"parameters"`
	Exported   bool            `json:"exported"`
}

type v4Catalog struct {
	Certname          string           `json:"certname"`
	Version           string           `json:"version"`
	Environment       string           `json:"environment"`
	TransactionUUID   string           `json:"transaction_uuid"`
	ProducerTimestamp string           `json:"producer_timestamp"`
	Producer          string           `json:"producer"`
	CatalogUUID       string           `json:"catalog_uuid"`
	CodeID            json.RawMessage  `json:"code_id"`
	Edges             catalogEdges     `json:"edges"`
	Resources         catalogResources `json:"resources"`
}

type v4CatalogEdges []v4CatalogEdge

type v4CatalogResources struct {
	Data catalogResources `json:"data"`
}

type v4CatalogEdge struct {
	SourceType   string `json:"source_type"`
	SourceTitle  string `json:"source_title"`
	TargetType   string `json:"target_type"`
	TargetTitle  string `json:"target_title"`
	Relationship string `json:"relationship"`
}

type v3CatalogData struct {
	Name              json.RawMessage `json:"name"`
	Version           json.RawMessage `json:"version"`
	Environment       json.RawMessage `json:"environment,omitempty"`
	TransactionUUID   json.RawMessage `json:"transaction-uuid"`
	ProducerTimestamp string          `json:"producer-timestamp,omitempty"`
	// Edges             v3CatalogEdges     `json:"edges"`
	// Resources         v3CatalogResources `json:"resources"`
}

func v3toV4CatalogConv(v3c v3Catalog) v4Catalog {
	var v4c v4Catalog
	v4c.Certname = v3c.Name
	v4c.Version = v3c.Version
	v4c.Environment = opts.Environment
	v4c.TransactionUUID = v3c.TransactionUUID
	v4c.Producer = opts.Producer
	v4c.ProducerTimestamp = time.Now().Format(time.RFC3339)
	v4c.CatalogUUID = uuid.New().String()
	v4c.Edges = v3c.Edges
	v4c.Resources = v3c.Resources

	return v4c
}

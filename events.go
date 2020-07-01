package main

import "encoding/json"

type v3Event struct {
	Certname          json.RawMessage `json:"certname"`
	OldValue          json.RawMessage `json:"old-value"`
	Property          json.RawMessage `json:"property"`
	Timestamp         json.RawMessage `json:"timestamp"`
	ResourceType      json.RawMessage `json:"resource-type"`
	ResourceTitle     json.RawMessage `json:"resource-title"`
	NewValue          json.RawMessage `json:"new-value"`
	Message           json.RawMessage `json:"message"`
	Report            json.RawMessage `json:"report"`
	Status            json.RawMessage `json:"status"`
	File              json.RawMessage `json:"file"`
	Line              json.RawMessage `json:"line"`
	ContainmentPath   json.RawMessage `json:"containment-path"`
	ContainingClass   json.RawMessage `json:"containing-class"`
	RunStartTime      json.RawMessage `json:"run-start-time"`
	RunEndTime        json.RawMessage `json:"run-end-time"`
	ReportReceiveTime json.RawMessage `json:"report-receive-time"`
}

type v3Events []v3Event

type aliasEvent struct {
	Certname          json.RawMessage `json:"certname"`
	OldValue          json.RawMessage `json:"old_value"`
	Property          json.RawMessage `json:"property"`
	Timestamp         json.RawMessage `json:"timestamp"`
	ResourceType      json.RawMessage `json:"resource_type"`
	ResourceTitle     json.RawMessage `json:"resource_title"`
	NewValue          json.RawMessage `json:"new_value"`
	Message           json.RawMessage `json:"message"`
	Report            json.RawMessage `json:"report"`
	Status            json.RawMessage `json:"status"`
	File              json.RawMessage `json:"file"`
	Line              json.RawMessage `json:"line"`
	ContainmentPath   json.RawMessage `json:"containment_path"`
	ContainingClass   json.RawMessage `json:"containing_class"`
	RunStartTime      json.RawMessage `json:"run_start_time"`
	RunEndTime        json.RawMessage `json:"run_end_time"`
	ReportReceiveTime json.RawMessage `json:"report_receive_time"`
}

type aliasEvents []aliasEvent

func (e *v3Event) MarshalJSON() ([]byte, error) {
	var a = aliasEvent(*e)
	return json.Marshal(&a)
}

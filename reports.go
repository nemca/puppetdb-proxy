package main

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type v3Report struct {
	Certname             string           `json:"certname"`
	Hash                 string           `json:"hash,omitempty"`
	Environment          string           `json:"environment,omitempty"`
	Status               string           `json:"status,omitempty"`
	PuppetVersion        string           `json:"puppet-version"`
	ReportFormat         int              `json:"report-format"`
	ConfigurationVersion string           `json:"configuration-version"`
	StartTime            string           `json:"start-time"`
	EndTime              string           `json:"end-time"`
	ReceiveTime          string           `json:"receive-time"`
	TransactionUUID      string           `json:"transaction-uuid"`
	ResourceEvents       v3ResourceEvents `json:"resource-events,omitempty"`
}

type v3Reports []v3Report

type v4Report struct {
	Certname             string          `json:"certname"`
	PuppetVersion        string          `json:"puppet_version"`
	ReportFormat         int             `json:"report_format"`
	ConfigurationVersion string          `json:"configuration_version"`
	StartTime            string          `json:"start_time"`
	EndTime              string          `json:"end_time"`
	ProducerTimestamp    string          `json:"producer_timestamp"`
	Producer             string          `json:"producer"`
	CorrectiveChange     bool            `json:"corrective_change"`
	Resources            v4Resources     `json:"resources"`
	Noop                 bool            `json:"noop"`
	NoopPending          bool            `json:"noop_pending"`
	TransactionUUID      string          `json:"transaction_uuid"`
	CatalogUUID          string          `json:"catalog_uuid"`
	CodeID               json.RawMessage `json:"code_id"`
	JobID                json.RawMessage `json:"job_id"`
	CachedCatalogStatus  string          `json:"cached_catalog_status"`
	Metrics              v4Metrics       `json:"metrics"`
	Logs                 v4Logs          `json:"logs"`
	Environment          string          `json:"environment,omitempty"`
	Status               string          `json:"status,omitempty"`
}

type v4Reports []v4Report

type v3ResourceEvent struct {
	ResourceType    string          `json:"resource-type"`
	ResourceTitle   string          `json:"resource-title"`
	Property        string          `json:"property"`
	TimeStamp       string          `json:"timestamp"`
	Status          string          `json:"status"`
	OldValue        json.RawMessage `json:"old-value"`
	NewValue        json.RawMessage `json:"new-value"`
	Message         string          `json:"message"`
	File            string          `json:"file"`
	Line            int             `json:"line"`
	ContainmentPath []string        `json:"containment-path"`
}

type v3ResourceEvents []v3ResourceEvent

type v4Resource struct {
	ResourceType     string                    `json:"resource_type"`
	ResourceTitle    string                    `json:"resource_title"`
	TimeStamp        string                    `json:"timestamp"`
	Skipped          bool                      `json:"skipped"`
	File             string                    `json:"file"`
	Line             int                       `json:"line"`
	Events           []v4ResourceEventExpanded `json:"events"`
	ContainmentPath  []string                  `json:"containment_path"`
	CorrectiveChange bool                      `json:"corrective_change"`
}

type v4Resources []v4Resource

type v4ResourceEvents struct {
	Data v4ResourceEventData `json:"data"`
	Href string              `json:"href,omitempty"`
}

// type v4ResourceEvents []v4ResourceEvent
type v4ResourceEventData []v4ResourceEventExpanded

type v4ResourceEventExpanded struct {
	Status           string          `json:"status"`
	TimeStamp        string          `json:"timestamp"`
	ResourceType     string          `json:"resource_type,omitempty"`
	ResourceTitle    string          `json:"resource_title,omitempty"`
	Property         string          `json:"property"`
	OldValue         json.RawMessage `json:"old_value"`
	NewValue         json.RawMessage `json:"new_value"`
	Message          string          `json:"message"`
	File             string          `json:"file,omitempty"`
	Line             int             `json:"line,omitempty"`
	ContainmentPath  []string        `json:"containment_path,omitempty"`
	CorrectiveChange bool            `json:"corrective_change"`
}

type v4Metric struct {
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Value    float64 `json:"value"`
}

type v4Metrics []v4Metric

type v4MetricsGet struct {
	Data []v4Metric `json:"data,omitempty"`
	Href string     `json:"href,omitempty"`
}

type v4Log struct {
	File    string   `json:"file"`
	Line    int      `json:"line"`
	Level   string   `json:"level"`
	Message string   `json:"message"`
	Source  string   `json:"source"`
	Tags    []string `json:"tags"`
	Time    string   `json:"time"`
}

type v4Logs []v4Log

type v4LogsGet struct {
	Data []v4Log `json:"data,omitempty"`
	Href string  `json:"href,omitempty"`
}

type v4ReportGet struct {
	Certname             string           `json:"certname"`
	PuppetVersion        string           `json:"puppet_version"`
	ReportFormat         int              `json:"report_format"`
	ConfigurationVersion string           `json:"configuration_version"`
	StartTime            string           `json:"start_time"`
	EndTime              string           `json:"end_time"`
	ProducerTimestamp    string           `json:"producer_timestamp"`
	Producer             string           `json:"producer"`
	CorrectiveChange     json.RawMessage  `json:"corrective_change"`
	Resources            []string         `json:"resources"`
	Noop                 bool             `json:"noop"`
	NoopPending          bool             `json:"noop_pending"`
	TransactionUUID      string           `json:"transaction_uuid"`
	CatalogUUID          string           `json:"catalog_uuid"`
	CodeID               json.RawMessage  `json:"code_id"`
	JobID                json.RawMessage  `json:"job_id"`
	CachedCatalogStatus  string           `json:"cached_catalog_status"`
	Metrics              v4MetricsGet     `json:"metrics"`
	Logs                 v4LogsGet        `json:"logs"`
	Environment          string           `json:"environment,omitempty"`
	Status               string           `json:"status,omitempty"`
	Hash                 string           `json:"hash,omitempty"`
	ReceiveTime          string           `json:"receive_time,omitempty"`
	ResourceEvents       v4ResourceEvents `json:"resource_events"`
}

type v4ReportsGet []v4ReportGet

func v3toV4ReportConv(v3r v3Report) v4Report {
	var v4r v4Report
	v4r.Certname = v3r.Certname
	v4r.Environment = opts.Environment
	v4r.Status = v3r.Status
	v4r.PuppetVersion = v3r.PuppetVersion
	v4r.ReportFormat = 8
	v4r.ProducerTimestamp = time.Now().Format(time.RFC3339)
	v4r.Producer = opts.Producer
	v4r.ConfigurationVersion = v3r.ConfigurationVersion
	v4r.StartTime = v3r.StartTime
	v4r.EndTime = v3r.EndTime
	v4r.CatalogUUID = uuid.New().String()
	v4r.CachedCatalogStatus = "not_used"
	v4r.TransactionUUID = v3r.TransactionUUID
	v4r.CorrectiveChange = false
	// Fake log
	var v4l v4Log
	v4l.Level = "notice"
	v4l.Tags = append(v4l.Tags, "notice")
	v4l.Source = "Puppet"
	v4l.Time = time.Now().Format(time.RFC3339)
	v4l.Message = "Puppet agent v3 does not send messages"
	v4r.Logs = append(v4r.Logs, v4l)

	var eventsStatus = make(map[string]float64)

	for _, v3res := range v3r.ResourceEvents {
		var v4res v4Resource
		v4res.ResourceType = v3res.ResourceType
		v4res.ResourceTitle = v3res.ResourceTitle
		v4res.TimeStamp = v3res.TimeStamp
		v4res.File = v3res.File
		v4res.Line = v3res.Line
		v4res.ContainmentPath = v3res.ContainmentPath
		v4res.CorrectiveChange = false

		var v4ree v4ResourceEventExpanded
		v4ree.Status = v3res.Status
		eventsStatus[v3res.Status]++
		eventsStatus["total"]++
		v4ree.TimeStamp = v3res.TimeStamp
		v4ree.NewValue = v3res.NewValue
		v4ree.OldValue = v3res.OldValue
		v4ree.Message = v3res.Message
		v4ree.CorrectiveChange = false
		v4ree.Property = v3res.Property
		v4res.Events = append(v4res.Events, v4ree)

		v4r.Resources = append(v4r.Resources, v4res)
	}

	for k, v := range eventsStatus {
		var v4m v4Metric
		v4m.Category = "events"
		v4m.Name = k
		v4m.Value = v
		v4r.Metrics = append(v4r.Metrics, v4m)
	}

	return v4r
}

func v4toV3ReportsConv(v4rs v4ReportsGet) v3Reports {
	var v3rs v3Reports
	for _, v4r := range v4rs {
		var v3r v3Report
		v3r.Certname = v4r.Certname
		v3r.Environment = opts.Environment
		v3r.Hash = v4r.Hash
		v3r.Status = v4r.Status
		v3r.PuppetVersion = v4r.PuppetVersion
		v3r.ReportFormat = 4
		v3r.ConfigurationVersion = v4r.ConfigurationVersion
		v3r.StartTime = v4r.StartTime
		v3r.EndTime = v4r.EndTime
		v3r.ReceiveTime = v4r.ReceiveTime
		v3r.TransactionUUID = v4r.TransactionUUID
		v3rs = append(v3rs, v3r)
	}

	return v3rs
}

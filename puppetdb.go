package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	oldPuppetTimeLayout = "2006-01-02T15:04:05-07:00"
	newPuppetTimeLayout = "2006-01-02T15:04:05.999Z"
)

func getNodes(vs url.Values) (v3Nodes, error) {
	body, err := getWithData(vs, "nodes")
	if err != nil {
		return v3Nodes{}, err
	}

	var nodes v3Nodes
	err = json.Unmarshal(body, &nodes)
	if err != nil {
		return v3Nodes{}, err
	}

	return nodes, nil
}

func getNodeByName(vs url.Values, name string) (v3Node, error) {
	body, err := getWithData(vs, "nodes/"+name)
	if err != nil {
		return v3Node{}, err
	}

	var node v3Node
	err = json.Unmarshal(body, &node)
	if err != nil {
		return v3Node{}, err
	}

	return node, nil
}

func getFacts(vs url.Values) ([]v3FactsGet, error) {
	body, err := getWithData(vs, "facts")
	if err != nil {
		return nil, err
	}

	var facts []v3FactsGet
	err = json.Unmarshal(body, &facts)
	if err != nil {
		return nil, err
	}

	return facts, nil
}

func getFactsByName(vs url.Values, fact string) ([]v3FactsGet, error) {
	body, err := getWithData(vs, "facts/"+fact)
	if err != nil {
		return nil, err
	}

	var facts []v3FactsGet
	err = json.Unmarshal(body, &facts)
	if err != nil {
		return nil, err
	}

	return facts, nil
}

func getFactsByNameAndValue(vs url.Values, fact, value string) ([]v3Facts, error) {
	body, err := getWithData(vs, "facts/"+fact+"/"+value)
	if err != nil {
		return nil, err
	}

	var facts []v3Facts
	err = json.Unmarshal(body, &facts)
	if err != nil {
		return nil, err
	}

	return facts, nil
}

func getNodeFactsByName(vs url.Values, name string) ([]v3FactsGet, error) {
	body, err := getWithData(vs, "nodes/"+name+"/facts")
	if err != nil {
		return nil, err
	}

	var facts []v3FactsGet
	err = json.Unmarshal(body, &facts)
	if err != nil {
		return nil, err
	}

	return facts, nil
}

func getFactsByNameAndFact(vs url.Values, name, fact string) ([]v3Facts, error) {
	body, err := getWithData(vs, "nodes/"+name+"/facts/"+fact)
	if err != nil {
		return nil, err
	}

	var facts []v3Facts
	err = json.Unmarshal(body, &facts)
	if err != nil {
		return nil, err
	}

	return facts, nil
}

func getFactsByNameAndFactValue(vs url.Values, name, fact, value string) (v3Facts, error) {
	body, err := getWithData(vs, "nodes/"+name+"/facts/"+fact+"/"+value)
	if err != nil {
		return v3Facts{}, err
	}

	var facts v3Facts
	err = json.Unmarshal(body, &facts)
	if err != nil {
		return v3Facts{}, err
	}

	return facts, nil
}

func getResourcesByNode(vs url.Values, name string) (v3Resources, error) {
	body, err := getWithData(vs, "nodes/"+name+"/resources")
	if err != nil {
		return v3Resources{}, err
	}

	var resources v3Resources
	err = json.Unmarshal(body, &resources)
	if err != nil {
		return v3Resources{}, err
	}

	return resources, nil
}

func getResourcesByNodeAndType(vs url.Values, name, t string) (v3Resources, error) {
	body, err := getWithData(vs, "nodes/"+name+"/resources/"+t)
	if err != nil {
		return v3Resources{}, err
	}

	var resources v3Resources
	err = json.Unmarshal(body, &resources)
	if err != nil {
		return v3Resources{}, err
	}

	return resources, err
}

func getResourcesByNodeAndTypeAndTitle(vs url.Values, name, t, title string) (v3Resources, error) {
	body, err := getWithData(vs, "nodes/"+name+"/resources/"+t+"/"+title)
	if err != nil {
		return v3Resources{}, err
	}

	var resources v3Resources
	err = json.Unmarshal(body, &resources)
	if err != nil {
		return v3Resources{}, err
	}

	return resources, err
}

// func v3toV4CatalogConv() {}
func v4toV3CatalogConv(v4c v4Catalog) v3CatalogData {
	var v3c v3CatalogData

	/*
		v3c.Name = v4c.Certname
		v3c.Version = v4c.Version
		v3c.TransactionUUID = v4c.TransactionUUID
		for _, v4e := range v4c.Edges.Data {
			var v3e v3CatalogEdge
			v3e.Relationship = v4e.Relationship
			v3e.Source.Type = v4e.SourceType
			v3e.Source.Title = v4e.SourceTitle
			v3e.Target.Type = v4e.TargetType
			v3e.Target.Title = v4e.TargetTitle
			v3c.Edges = append(v3c.Edges, v3e)
		}
		v3c.Resources = v4c.Resources.Data
	*/

	return v3c
}

func getReports(vs url.Values) (v3Reports, error) {
	body, err := getWithData(vs, "reports")
	if err != nil {
		return v3Reports{}, err
	}

	var reports v4ReportsGet
	err = json.Unmarshal(body, &reports)
	if err != nil {
		return v3Reports{}, err
	}

	var v3r = v4toV3ReportsConv(reports)

	return v3r, nil
}

func getEventCounts(vs url.Values) (v3EventCounts, error) {
	body, err := getWithData(vs, "event-counts")
	if err != nil {
		return v3EventCounts{}, err
	}

	var ec v3EventCounts
	err = json.Unmarshal(body, &ec)
	if err != nil {
		return v3EventCounts{}, err
	}

	return ec, nil
}

func getAggregateEventCounts(vs url.Values) (v3AggregateEventCount, error) {
	body, err := getWithData(vs, "aggregate-event-counts")
	if err != nil {
		return v3AggregateEventCount{}, err
	}

	var aecs v3AggregateEventCounts
	err = json.Unmarshal(body, &aecs)
	if err != nil {
		return v3AggregateEventCount{}, err
	}

	return aecs[0], nil
}

func getResources(vs url.Values) (v3Resources, error) {
	body, err := getWithData(vs, "resources")
	if err != nil {
		return v3Resources{}, err
	}

	var r v3Resources
	err = json.Unmarshal(body, &r)
	if err != nil {
		return v3Resources{}, err
	}

	return r, nil
}

func getResourcesByType(vs url.Values, t string) (v3Resources, error) {
	body, err := getWithData(vs, "resources/"+t)
	if err != nil {
		return v3Resources{}, err
	}

	var r v3Resources
	err = json.Unmarshal(body, &r)
	if err != nil {
		return v3Resources{}, err
	}

	return r, nil
}

func getResourcesByTypeAndTitle(vs url.Values, t, title string) (v3Resources, error) {
	body, err := getWithData(vs, "resources/"+t+"/"+title)
	if err != nil {
		return v3Resources{}, err
	}

	var r v3Resources
	err = json.Unmarshal(body, &r)
	if err != nil {
		return v3Resources{}, err
	}

	return r, nil
}

func getEvents(vs url.Values) (v3Events, error) {
	body, err := getWithData(vs, "events")
	if err != nil {
		return v3Events{}, err
	}

	var e v3Events
	err = json.Unmarshal(body, &e)
	if err != nil {
		return v3Events{}, err
	}

	return e, nil
}

func getWithData(vs url.Values, uri string) ([]byte, error) {
	client := http.Client{}
	data := ioutil.NopCloser(strings.NewReader(vs.Encode()))
	req, err := http.NewRequest(http.MethodGet, opts.PuppetDBURL+"/pdb/query/v4/"+uri, data)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []byte(`[ ]`), nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func postWithData(body []byte, values url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", opts.PuppetDBURL+"/pdb/cmd/v1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Add URL query params
	req.URL.RawQuery = valuesToString(values)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func valuesToString(values url.Values) string {
	var ret string
	count := 1
	for k, v := range values {
		if k == "producer-timestamp" {
			t, err := time.Parse(oldPuppetTimeLayout, v[0])
			if err == nil {
				v[0] = t.Format(newPuppetTimeLayout)
				v[0] = strings.Replace(v[0], "Z", ".000Z", 1)
			}
		}
		if count == 1 {
			ret += fmt.Sprintf("%s=%s", k, v[0])
			count++
			continue
		}
		ret += fmt.Sprintf("&%s=%s", k, v[0])
	}

	return ret
}

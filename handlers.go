// handlers.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *server) initRoutes() {
	// v3 API
	v3 := s.Router.PathPrefix("/v3").Subrouter()
	v3.HandleFunc("/nodes", s.v3nodesHandler).Methods(http.MethodGet)
	v3.HandleFunc("/nodes/{name}", s.v3nodeWithNameHandler).Methods(http.MethodGet)
	v3.HandleFunc("/nodes/{name}/facts", s.v3nodeFactsHandler).Methods(http.MethodGet)
	v3.HandleFunc("/nodes/{name}/facts/{fact}", s.v3nodeFactHandler).Methods(http.MethodGet)
	v3.HandleFunc("/nodes/{name}/facts/{fact}/{value}", s.v3nodeFactValueHandler).Methods(http.MethodGet)
	v3.HandleFunc("/nodes/{name}/resources", s.v3nodeResourcesHandler).Methods(http.MethodGet)
	v3.HandleFunc("/nodes/{name}/resources/{type}", s.v3nodeResourcesByTypeHandler).Methods(http.MethodGet)
	v3.HandleFunc("/nodes/{name}/resources/{type}/{title}", s.v3nodeResourcesByTypeAndTitleHandler).Methods(http.MethodGet)
	v3.HandleFunc("/facts", s.v3factsHandler).Methods(http.MethodGet)
	v3.HandleFunc("/facts/{fact}", s.v3factByName).Methods(http.MethodGet)
	v3.HandleFunc("/facts/{fact}/{value}", s.v3factByNameAndValue).Methods(http.MethodGet)
	v3.HandleFunc("/fact-names", s.v3factNamesHandler).Methods(http.MethodGet)
	v3.HandleFunc("/catalogs/{name}", s.v3catalogsByNameHandler).Methods(http.MethodGet)
	v3.HandleFunc("/resources", s.v3resourcesHandler).Methods(http.MethodGet)
	v3.HandleFunc("/resources/{type}", s.v3resourcesByTypeHandler).Methods(http.MethodGet)
	v3.HandleFunc("/resources/{type}/{title}", s.v3resourcesByTypeAndTitleHandler).Methods(http.MethodGet)
	v3.HandleFunc("/reports", s.v3reportsHandler).Methods(http.MethodGet)
	v3.HandleFunc("/events", s.v3eventsHandler).Methods(http.MethodGet)
	v3.HandleFunc("/event-counts", s.v3eventCountsHandler).Methods(http.MethodGet)
	v3.HandleFunc("/aggregate-event-counts", s.v3aggregateEventCountsHandler).Methods(http.MethodGet)
	v3.HandleFunc("/server-time", s.v3serverTimeHandler).Methods(http.MethodGet)
	v3.HandleFunc("/version", s.v3versionHandler).Methods(http.MethodGet)

	// Commands
	v3.HandleFunc("/commands", s.v3commandsHandler).Methods(http.MethodPost)

	// Prometheus
	s.Router.Handle("/metrics", promhttp.Handler())
}

func (s *server) v3nodesHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	nodes, err := getNodes(vs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.Log.Errorf("failed to get nodes: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nodes)
}

func (s *server) v3nodeWithNameHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]

	node, err := getNodeByName(vs, name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.Log.Errorf("failed to get nodes by name: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(node)
}

func (s *server) v3nodeFactsHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]

	facts, err := getNodeFactsByName(vs, name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get node facts by node name: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(facts)
}

func (s *server) v3nodeFactHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]
	fact := vars["fact"]

	facts, err := getFactsByNameAndFact(vs, name, fact)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get node facts by nodeand fact names: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(facts)
}

func (s *server) v3nodeFactValueHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]
	fact := vars["fact"]
	value := vars["value"]

	facts, err := getFactsByNameAndFactValue(vs, name, fact, value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get node facts by node and fact names and fact value: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(facts)
}

func (s *server) v3nodeResourcesHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}
	// s.Log.Debug(vs)

	vars := mux.Vars(r)
	name := vars["name"]

	resorces, err := getResourcesByNode(vs, name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("falied to get node resources by node name: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resorces)
}

func (s *server) v3nodeResourcesByTypeHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]
	t := vars["type"]

	resorces, err := getResourcesByNodeAndType(vs, name, t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get node resources by node name and resource type: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resorces)
}

func (s *server) v3nodeResourcesByTypeAndTitleHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]
	t := vars["type"]
	title := vars["title"]

	resorces, err := getResourcesByNodeAndTypeAndTitle(vs, name, t, title)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get node resources by node name and resource type and title: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resorces)
}

func (s *server) v3factsHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	facts, err := getFacts(vs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get facts: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(facts)
}

func (s *server) v3factByName(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	vars := mux.Vars(r)
	fact := vars["fact"]

	facts, err := getFactsByName(vs, fact)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get facts by name: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(facts)
}

func (s *server) v3factByNameAndValue(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	vars := mux.Vars(r)
	fact := vars["fact"]
	value := vars["value"]

	facts, err := getFactsByNameAndValue(vs, fact, value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed get facts by name and value: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(facts)
}

func (s *server) v3factNamesHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(opts.PuppetDBURL + "/pdb/query/v4/fact-names")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed get fact names: %v", err)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, resp.Body)
}

func (s *server) v3catalogsByNameHandler(w http.ResponseWriter, r *http.Request) {
	return
	/*
		vars := mux.Vars(r)
		name := vars["name"]

		resp, err := http.Get(puppetDB + "/pdb/query/v4/catalogs/" + name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var v4c v4Catalog
		var v3c v3Catalog
		err = json.Unmarshal(body, &v4c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		v3c.Data = v4toV3CatalogConv(v4c)
		v3c.Metadata.ApiVersion = 1

		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(v3c)
	*/
}

func (s *server) v3resourcesHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}
	s.Log.Trace(vs)

	resources, err := getResources(vs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get resources: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resources)
}

func (s *server) v3resourcesByTypeHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}
	s.Log.Trace(vs)

	vars := mux.Vars(r)
	t := vars["type"]

	resources, err := getResourcesByType(vs, t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get resources by type: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resources)
}

func (s *server) v3resourcesByTypeAndTitleHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}
	s.Log.Trace(vs)

	vars := mux.Vars(r)
	t := vars["type"]
	title := vars["title"]

	resources, err := getResourcesByTypeAndTitle(vs, t, title)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed get resources by type and title: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resources)
}

func (s *server) v3eventsHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	events, err := getEvents(vs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get events: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func (s *server) v3reportsHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	reports, err := getReports(vs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get reports: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reports)
}

func (s *server) v3eventCountsHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	ec, err := getEventCounts(vs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get event counts: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ec)
}

func (s *server) v3aggregateEventCountsHandler(w http.ResponseWriter, r *http.Request) {
	vs, err := parseForm(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed parse request body: %v", err)
		return
	}

	aec, err := getAggregateEventCounts(vs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get aggregate event counts: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aec)
}

func (s *server) v3serverTimeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(opts.PuppetDBURL + "/pdb/meta/v1/server-time")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get server time: %v", err)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, resp.Body)
}

func (s *server) v3versionHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(opts.PuppetDBURL + "/pdb/meta/v1/version")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to get server version: %v", err)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, resp.Body)
}

func (s *server) v3commandsHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var v3c v3Commands
	var v4c v4Commands
	var values url.Values

	err := decoder.Decode(&v3c)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to decode request body: %v", err)
		return
	}

	v4c.Command = v3c.Command
	switch v4c.Command {
	case "replace facts":
		v4c.Version = 5
		v4c.Payload, values, err = getV4FactsPayload(v3c.Payload)
		if err != nil {
			body, _ := ioutil.ReadAll(r.Body)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			s.Log.Errorf("failed parse to payload for replace facts command: %v", err)
			s.Log.Trace(string(body))
			return
		}
	case "replace catalog":
		v4c.Version = 9
		v4c.Payload, values, err = getV4CatalogPayload(v3c.Payload)
		if err != nil {
			body, _ := ioutil.ReadAll(r.Body)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			s.Log.Errorf("failed parse to payload for replace catalog command: %v", err)
			s.Log.Trace(string(body))
			return
		}
	case "store report":
		v4c.Version = 8
		v4c.Payload, values, err = getV4ReportPayload(v3c.Payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			s.Log.Errorf("failed parse to payload for store report command: %v", err)
			s.Log.Trace(string(v3c.Payload))
			return
		}
	case "deactivate node":
		v4c.Version = 3
		v4c.Payload, values, err = getV4DeactivatePayload(v3c.Payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			s.Log.Errorf("failed to parse payload for deactivate node command: %v", err)
			return
		}
	default:
		s.Log.Error("unable to unmarshal JSON data")
	}

	// Add URL parameters
	values.Set("command", strings.Replace(v4c.Command, " ", "_", -1))
	values.Set("version", strconv.Itoa(v4c.Version))

	body, err := json.Marshal(v4c.Payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to marshal v4c: %v", err)
		return
	}
	resp, err := postWithData(body, values)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to do POST request with data: %v", err)
		return
	}
	var data response
	err = json.Unmarshal(resp, &data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		s.Log.Errorf("failed to unmarshal response: %v", err)
		s.Log.Trace(string(resp))
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func getV4FactsPayload(v3payload json.RawMessage) (json.RawMessage, url.Values, error) {
	var v3f v3Facts
	if err := json.Unmarshal(v3payload, &v3f); err != nil {
		return nil, nil, err
	}
	if opts.DumpFacts && opts.DumpHostname != "" {
		if v3f.Certname == opts.DumpHostname {
			file, _ := os.OpenFile("/tmp/"+opts.DumpHostname+"-reports.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			defer file.Close()
			fmt.Fprintln(file, string(v3payload))
		}
	}
	var v4f = v3toV4FactsConv(v3f)

	v := url.Values{}
	v.Set("certname", v4f.Certname)
	v.Set("producer-timestamp", v4f.ProducerTimestamp)

	j, err := json.Marshal(&v4f)
	return j, v, err
}

func getV4ReportPayload(v3payload json.RawMessage) (json.RawMessage, url.Values, error) {
	var v3r v3Report
	if err := json.Unmarshal(v3payload, &v3r); err != nil {
		return nil, nil, err
	}
	if opts.DumpReport && opts.DumpHostname != "" {
		if v3r.Certname == opts.DumpHostname {
			file, _ := os.OpenFile("/tmp/"+opts.DumpHostname+"-reports.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			defer file.Close()
			fmt.Fprintln(file, string(v3payload))
		}
	}
	var report = v3toV4ReportConv(v3r)

	v := url.Values{}
	v.Set("certname", report.Certname)
	v.Set("producer-timestamp", report.ProducerTimestamp)

	j, err := json.Marshal(&report)
	return j, v, err
}

func getV4CatalogPayload(v3payload json.RawMessage) (json.RawMessage, url.Values, error) {
	var v3c v3Catalog
	if err := json.Unmarshal(v3payload, &v3c); err != nil {
		return nil, nil, err
	}
	if opts.DumpCatalog && opts.DumpHostname != "" {
		if v3c.Name == opts.DumpHostname {
			file, _ := os.OpenFile("/tmp/"+opts.DumpHostname+"-reports.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			defer file.Close()
			fmt.Fprintln(file, string(v3payload))
		}
	}
	var catalog = v3toV4CatalogConv(v3c)

	v := url.Values{}
	v.Set("certname", catalog.Certname)
	v.Set("producer-timestamp", catalog.ProducerTimestamp)

	j, err := json.Marshal(&catalog)
	return j, v, err
}

func getV4DeactivatePayload(v3payload json.RawMessage) (json.RawMessage, url.Values, error) {
	var v4c v4CommandsDeacticate
	v4c.Name = v3payload
	v4c.ProducerTimestamp = time.Now().Format(time.RFC3339)

	name, err := v4c.Name.MarshalJSON()
	if err != nil {
		return nil, nil, err
	}

	v := url.Values{}
	v.Set("certname", string(name))
	v.Set("producer-timestamp", v4c.ProducerTimestamp)

	j, err := json.Marshal(&v4c)
	return j, v, err
}

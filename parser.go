package main

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func parseForm(r *http.Request) (vs url.Values, err error) {
	var reader io.Reader = r.Body
	maxFormSize := int64(1<<63 - 1)
	if _, ok := r.Body.(*maxBytesReader); !ok {
		maxFormSize = int64(10 << 20) // 10 MB is a lot of text.
		reader = io.LimitReader(r.Body, maxFormSize+1)
	}
	b, e := ioutil.ReadAll(reader)
	if e != nil {
		if err == nil {
			err = e
		}
	}
	if int64(len(b)) > maxFormSize {
		err = errors.New("http: POST too large")
		return
	}
	vs, e = url.ParseQuery(string(b))
	if err == nil {
		err = e
	}

	urlQuery, e := url.ParseQuery(r.URL.RawQuery)
	if err == nil {
		err = e
	}
	query := urlQuery.Get("query")
	if query != "" {
		vs = urlQuery
	}

	// Replacing for v3 API compatible
	for k := range vs {
		if strings.Contains(k, "-") {
			n := strings.ReplaceAll(k, "-", "_")
			vs.Set(n, vs.Get(k))
			vs.Del(k)
		}
	}
	return
}

type maxBytesReader struct {
	w   http.ResponseWriter
	r   io.ReadCloser // underlying reader
	n   int64         // max bytes remaining
	err error         // sticky error
}

func (l *maxBytesReader) Close() error {
	return l.r.Close()
}

func (l *maxBytesReader) Read(p []byte) (n int, err error) {
	if l.err != nil {
		return 0, l.err
	}
	if len(p) == 0 {
		return 0, nil
	}
	// If they asked for a 32KB byte read but only 5 bytes are
	// remaining, no need to read 32KB. 6 bytes will answer the
	// question of the whether we hit the limit or go past it.
	if int64(len(p)) > l.n+1 {
		p = p[:l.n+1]
	}
	n, err = l.r.Read(p)

	if int64(n) <= l.n {
		l.n -= int64(n)
		l.err = err
		return n, err
	}

	n = int(l.n)
	l.n = 0

	// The server code and client code both use
	// maxBytesReader. This "requestTooLarge" check is
	// only used by the server code. To prevent binaries
	// which only using the HTTP Client code (such as
	// cmd/go) from also linking in the HTTP server, don't
	// use a static type assertion to the server
	// "*response" type. Check this interface instead:
	type requestTooLarger interface {
		requestTooLarge()
	}
	if res, ok := l.w.(requestTooLarger); ok {
		res.requestTooLarge()
	}
	l.err = errors.New("http: request body too large")
	return n, l.err
}

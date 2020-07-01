package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/jessevdk/go-flags"
)

var (
	version string
	build   string
)

var opts struct {
	ListenAddress string `short:"a" long:"listen.address" default:"127.0.0.1" description:"Listen address"`
	ListenPort    int    `short:"p" long:"port" default:"8088" description:"Listen port"`
	PuppetDBURL   string `short:"u" long:"puppetdb.url" default:"https://puppetdb.example.com" description:"URL for connection to PuppetDB"`
	Environment   string `short:"e" long:"environment" default:"production" description:"Change 'environment' field"`
	Producer      string `short:"P" long:"producer" default:"puppet.example.com" description:"Change 'producer' field"`
	Insecure      bool   `short:"k" long:"insecure" description:"Disable verify the server's certificate chain and hostname"`
	LogFile       string `short:"L" long:"log.file" default:"/var/log/puppetdb-proxy.log" description:"Path to logfile"`
	LogLevel      int    `short:"V" long:"log.level" default:"4" description:"Log level (0-6)"`
	Version       bool   `short:"v" long:"version" description:"Show version number and quit"`
	DumpHostname  string `short:"H" long:"dump.hostname" description:"Hostname of Puppet node for dumping the commands payload to file /tmp/$hostname-$command.json (use with -C|-F|-R options)"`
	DumpReport    bool   `short:"R" long:"dump.report" description:"Dump the command store report payload to file (use with -H option)"`
	DumpFacts     bool   `short:"F" long:"dump.facts" description:"Dump the command replace facts payload to file (use with -H option)"`
	DumpCatalog   bool   `short:"C" long:"dump.catalog" description:"Dump the command replace catalog payload to file (use with -H option)"`
}

func main() {
	if _, err := flags.Parse(&opts); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	if opts.Version {
		fmt.Printf("version: %s\nbuild:   %s\n", version, build)
		os.Exit(0)
	}

	if opts.Insecure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	s := newServer()

	addr := fmt.Sprintf("%s:%d", opts.ListenAddress, opts.ListenPort)
	s.run(addr)
}

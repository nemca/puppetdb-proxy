# puppetdb-proxy

The proxy service for Puppet for store data from old Puppet agents to new PuppetDB.
Testing on a Puppet agent v3.7.2 and PuppetDB v6.2.0.

## Build
```sh
# build binary
git clone https://github.com/nemca/puppetdb-proxy.git
cd puppetdb-proxy
make
```
```sh
# build Debian package
git clone https://github.com/nemca/puppetdb-proxy.git
cd puppetdb-proxy
make fpm-deb
```

## Usage
All variables can be added to `/etc/default/puppetdb-proxy`
```sh
./puppetdb-proxy --help
Usage:
  puppetdb-proxy [OPTIONS]

Application Options:
  -a, --listen.address= Listen address (default: 127.0.0.1)
  -p, --port=           Listen port (default: 8088)
  -u, --puppetdb.url=   URL for connection to PuppetDB (default: https://puppetdb.example.com)
  -e, --environment=    Change 'environment' field (default: production)
  -P, --producer=       Change 'producer' field (default: puppet.example.com)
  -k, --insecure        Disable verify the server's certificate chain and hostname
  -L, --log.file=       Path to logfile (default: /var/log/puppetdb-proxy.log)
  -V, --log.level=      Log level (0-6) (default: 4)
  -v, --version         Show version number and quit
  -H, --dump.hostname=  Hostname of Puppet node for dumping the commands payload to file /tmp/$hostname-$command.json (use with -C|-F|-R|-Q options)
  -R, --dump.report     Dump the command store report payload to file (use with -H option)
  -F, --dump.facts      Dump the command replace facts payload to file (use with -H option)
  -C, --dump.catalog    Dump the command replace catalog payload to file (use with -H option)
  -Q, --dump.query      Dump the query (use with -H option)

Help Options:
  -h, --help            Show this help message
```

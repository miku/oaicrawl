Haprot
======

Command line client for the Open Archives Initiative Protocol for Metadata
Harvesting (OAI-PMH).

```shell
$ haprot fetch -set x -prefix y http://export.arxiv.org/oai2
...
$ haprot fetch -set x -prefix y http://export.arxiv.org/oai2
...
```

Improvements over metha
-----------------------

* second resolution
* harvest by id, monthly chunks
* slice and filter by xpath expressions

Examples
--------

```shell
$ haprot cat -since 1.weeks.ago http://export.arxiv.org/oai2
...
$ haprot cat -since yesterday http://export.arxiv.org/oai2
...
$ haprot cat -since 2017-01-01 http://export.arxiv.org/oai2
...
```

Layout
------

There is one base directory for caching:

```shell
$ tree ~/.haprot
...
```

Each cache directory is the checksum of the harvesting options:

```shell
~/.haprot/c/sha256(Endpoint URL + Format + Set)
~/.haprot/c/sha256(Endpoint URL + Format + Set).json
```

Each Harvest comes with a json file, that contains various information about
the harvest. A list of identifiers, the date of the last fetch and a checksum
of the harvested files, so we can check, if descriptor and directory are
actually in sync.

Strategies
----------

* yearly harvesting windows
* montly harvesting windows
* daily harvesting windows
* hourly harvesting windows
* by identifier

All files are written to a temporary location first and then assembled in a
unified way for caching.

Usage as library
----------------

```go
endpoint := haprot.Endpoint{URL: "http://export.arxiv.org/oai2"}
ping, err := endpoint.Ping()
if err != nil {
    log.Fatal(err)
}
fmt.Println(ping)
```

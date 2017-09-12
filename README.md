oaicrawl
========

Harvest an OAI endpoint by fetching records one by one. This is a different
strategy from the windowed and cached approach used in
[metha](https://github.com/miku/metha).

Install via *go get* or [releases](https://github.com/miku/oaicrawl/releases).

oaicrawl does not cache anything and will write the raw responses directly to
standard output.

```shell
$ oaicrawl -f oai_dc -verbose http://www.academicpub.org/wapoai/OAI.aspx > harvest.data
...
DEBU[0025] fetched 1395 identifiers with 1 requests in 25.140568087s
```

Use the `-b` flag to crawl in a best effort way (continue in the presence of
errors):

```shell
$ oaicrawl -b -f oai_dc -verbose http://www.academicpub.org/wapoai/OAI.aspx > harvest.data
...
... worker-11 backoff [3]: ..ertasacademica.com/5318&metadataPrefix=oai_dc

```

This crawler was written for working with endpoints that are slightly
off-standard and cannot be harvested easily in chunks.

Test it yourself (might take a day to harvest completely):

```shell
$ oaicrawl -f mets -b -verbose http://zvdd.de/oai2/
...
```

Usage
-----

```shell
$ oaicrawl -h
Usage of oaicrawl:
  -b    create best effort data set
  -e duration
        max elapsed time (default 10s)
  -f string
        format (default "oai_dc")
  -retry int
        max number of retries (default 3)
  -verbose
        more logging
  -version
        show version
  -w int
        number of parallel connections (default 16)
```

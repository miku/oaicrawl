oaicrawl
========

Harvest an OAI endpoint by fetching records one by one. This is a different
strategy from the windowed and cached approach used in
[metha](https://github.com/miku/metha).

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
$ oaicrawl -f oai_dc -verbose http://www.academicpub.org/wapoai/OAI.aspx > harvest.data
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

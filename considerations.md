# Considerations

Should you use TypeDB in a project instead of any other graph databases currently available?

It depends on the nature of the project. For research projects that do not require 24/7 operations, but would benefit
from the unique type system, then TypeDB might be a good idea. However, for commercial projects that must perform 24/7 at scale, I most certainly prefer any one of the following alternatives:

* [AnzoGraph](https://cambridgesemantics.com/anzograph/)
* [ArangoDB](https://www.arangodb.com/)
* [InfiniteGraph](https://infinitegraph.com/)
* [Heavy.ai (Omnisci)](https://www.heavy.ai/product/overview)
* [GridGain](https://www.gridgain.com/)

For massive internet data that require multiple data models i.e. table, KV, Graph, I would look into:

* [FaunaDB](https://fauna.com/)
* [MacroMeta GDN](https://macrometa.com/)

For graph based machine learning, I would look into:

* [Deep Graph Library](http://dgl.ai/)

I ended up shortlisting FaunaDB & Macrometa. FaunaDB clearly has the more advanced product, but it's only available in Europe & the US. Macrometa, on the other hand, may not have GraphQL, but it's globally available with an average P90 latency of under 50ms. Because my primary application serves South East
Asia with a cluster colocated in Tokio, I ended up selecting the Macrometa GDN. In a simple benchmark, I actually got an
average P90 latency of 20 ms for Tokio. At this point, the GDN data access over network was faster than my previous
in-cluster database. For South East Asia, average latency was about one order of magnitude faster than FaunaDB. That said, for use cases that require only the US, EU, or both regions, FaunaDB and Macromete perform roughly equal in terms of performance and latency so that would be a different decision. 

## TypeDB strength:

* Very innovative
* Unique type system, type-inheritance, n-ary relations & typed query language
* Reasoner & inference already build into TypeDB.

## TypeDB weaknesses:

### Operations:

* No Helm package / Kubernetes deployment for OSS version. Only in paid enterprise cluster.
  See [issue](https://github.com/vaticle/typedb/issues/6455)
* No RBAC. Not even in the cluster and no open issue...
* No rolling updates. Not even in the enterprise cluster...

### Non-Determinstic:

* Identical query returns different answers. See [issue](https://github.com/vaticle/typedb/issues/6336)
* Non-deterministic query results when using parallelization. See [issue](https://github.com/vaticle/typedb/issues/6349)
* Non-deterministic error handling for multiple transactions. See [issue](https://github.com/vaticle/typedb/issues/6146)

### Performance:

* Poor query performance when using parallelization. See [issue](https://github.com/vaticle/typedb/issues/6226)
* Poor performance on multi-hop queries. See [issue](https://github.com/vaticle/typedb/issues/6183)
* Reasoner inefficient & slow. See [issue](https://github.com/vaticle/typedb/issues/6467)
  , [issue](https://github.com/vaticle/typedb/issues/6453)

### Security:

* Non resolution of known security problems caused by old dependencies. Possibly worse, but unknown.
  See [issue](https://github.com/vaticle/typedb/issues/6301)
* No information about Log4J vulnerability.
* No disclosure if any code security actually happens.

### Missing features

* Scalability concerns prevent large knowledge graphs.
  See [issue](https://github.com/vaticle/kglib/issues/157#issuecomment-1092742408)
* No analytics / OLAP. See [issue](https://github.com/vaticle/typedb/issues/6394)
* No working graph machine learning. KGLib hasn't been maintained for a year.
  See [issue](https://github.com/vaticle/kglib/issues/152)
* Lack of vector type prevents, among others, using the Stanford OGB graph datasets.
  See [issue](https://github.com/vaticle/typedb/issues/6327)
* Lack of all collection types i.e. List, array, map, set...

## Reflections

After having implemented the Golang client, I was finally able to implement the call data demo in Golang, which also
concluded my evaluation. The type based data modelling clearly elevates TypeDB and actually accelerates data engineering
with only little training required. Query performance, and especially latency, however, remains poor for all examples.
For operations, well, the many open issues clearly speak for themselves.

It is incomprehensible to me which problem TypeDB solves that cannot be done with RDF graphs. According to their own
words, they do all that RDF does, but [mysteriously better](https://towardsdatascience.com/comparing-grakn-to-semantic-web-technologies-part-1-3-3558c447214a). However, unlike mature RDF systems, like AnzoGraph, TypeDB doesn't work well, doesn't scale, doesn't come with machine
learning, and can't even remotely compete on performance. To put things into perspective, the toy data example requires
about 4 seconds to complete 8 queries from maybe 50 data records, which results in average latency of about 500ms per
query. From my experience with TypeDB, 250 ms is a "good" latency of a trivial query. GridGain delivers about 50 ms latency on
the exact same hardware when working on 1 million data records. AnzoGraph returns arbitrarily complex queries
consistently below 100 ms when working on over 50 million data records, albeit in a 3 node GCP cluster. Omnisci performs
equally on sub second latency on [1 billion data records](https://tech.marksblogg.com/benchmarks.html) even on a
2019 [laptop](https://tech.marksblogg.com/omnisci-macos-macbookpro-mbp.html).

### Final thoughts

When Vaticle reached out to me in March 2022, I was reluctant to even give TypeDB a try because the use-case wasn't
clear to me, but ultimately I decided to evaluate TypeDB because of the type based data modelling assuming this would be
backed by a reasonable operational database. Speaking of TypeDB's primary application in life science, there is real
value to model complex biological data as types and relations. Outside this very specialized domain, however, there is
little that TypeDB brings to the table while adding a lot of troubles to operations.

Summing up my experience of implementing TypeDB protocol and working with TypeDB for about a month, I couldn't help but
realizing that this protocol is overly complex and error-prone relative to the very little expressiveness the query
language offers. However, the truly nagging question really became what does TypeDB actually offers that cannot be done
otherwise?

I was unable to answer that question.


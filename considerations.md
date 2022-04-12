# Considerations

Should you use TypeDB in a project instead of any other graph databases currently available?

It depends on the nature of the project. For research projects that do not require 24/7 operations, but would benefit
from the unique type system, then TypeDB might be a good idea.

For commercial projects that must perform 24/7 at scale, I most certainly prefer any one of the following alternatives:

* [AnzoGraph](https://cambridgesemantics.com/anzograph/)
* [ArangoDB](https://www.arangodb.com/)
* [InfiniteGraph](https://infinitegraph.com/)
* [Heavy.ai (Omnisci)](https://www.heavy.ai/product/overview)
* [GridGain](https://www.gridgain.com/)

## TypeDB strength:

* Very innovative
* Unique type system, type-inheritance, n-ary relations & typed query language
* Reasoner & inference already build into TypeDB.

## TypeDB weaknesses:

### Operations:

* No Helm package / Kubernetes deployment for OSS version. Only in paid enterprise cluster.
  See [issue](https://github.com/vaticle/typedb/issues/6455)
* No RBAC. Not even in the cluster and no open issue...
* No rolling updates. Not in the cluster...

### Determinism:

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
words, they do all that RDF does,
but [mysteriously better](https://towardsdatascience.com/comparing-grakn-to-semantic-web-technologies-part-1-3-3558c447214a)
. However, unlike mature RDF systems, like AnzoGraph, TypeDB doesn't work well, doesn't scale, doesn't come with machine
learning, and can't even remotely compete on performance. To put things into perspective, the toy data example requires
about 4 seconds to complete 8 queries from maybe 50 data records, which results in average latency of about 500ms per
query. from my own experience, 250 ms is a "good" latency of a trivial query. GridGain delivers about 50 ms latency on
the exact same hardware when working on 1 million data records. AnzoGraph returns arbitrarily complex queries
consistently below 100 ms when working on over 50 million data records, albeit in a 3 node GCP cluster. Omnisci performs
equally on sub second latency, but on [1 billion data records](https://tech.marksblogg.com/benchmarks.html) even on a
2019 [laptop](https://tech.marksblogg.com/omnisci-macos-macbookpro-mbp.html).

Summing up my experience of implementing TypeDB protocol, I couldn't help but thinking that this protocol is overly
complex and error-prone relative to the very little expressiveness the query language offers. However, the truly nagging
question really became what does TypeDB actually does that cannot be done otherwise?

I was unable to answer that question.

### Final thoughts

When Vaticle reached out to me in March 2022, I was reluctant to even give TypeDB a try because the use-case wasn't
clear to me, but ultimately I decided to evaluate TypeDB because of the type based data modelling assuming this would be
backed by a reasonable operational database. Speaking of TypeDB's primary application in life science, there is real
value to model complex biological data as types and relations. Outside this very specialized domain, however, there is
little that TypeDB brings to the table while adding a lot of troubles to operations. Looking back, it becomes an even
greater mystery to me why Vaticle even attempts to reach people outside their primary niche. 

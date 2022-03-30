# typedb-client-go

Go client for [TypeDB](https://vaticle.com/typedb)

Status: Early stage work in progress! 

Supports: TypeDB [2.7.1](https://github.com/vaticle/typedb/releases/tag/2.7.1)

Protocol: [2.6.1](https://github.com/vaticle/typedb-protocol/releases/tag/2.6.1)

NOT supported:
* V. 2.8
* Cluster

DB Admin

- [x] GetAllDatabases
- [x] CreateDatabase
- [x] CheckDatabaseExists 
- [x] DeleteDatabase

Schema:
- [x] CreateDatabaseSchema
- [x] GetDatabaseSchema

Transaction (TransactionManager):
- [x] NewTransaction
- [x] OpenTransaction
- [x] ExecuteTransaction
- [x] CommitTransaction
- [x] RollbackTransaction
- [x] CloseTransaction

Query:
- [x] RunInsertQuery
- [x] RunUpdateQuery
- [x] RunDeleteQuery
- [x] RunExplainQuery
- [x] RunDefineQuery
- [x] RunUnDefineQuery
- [x] RunMatchQuery
- [x] RunMatchGroupQuery
- [x] RunMatchAggregateQuery
- [x] RunMatchGroupAggregateQuery

## Make reference

```bash 
Setup: 
    make check                  Checks whether all project requirements are present.
     
Dev: 
    make build                  Builds the code base incrementally (fast).
    make rebuild                Rebuilds dependencies, build files, & code base (slow). Use after go mod changes.
    make tests                  Runs all tests. See scripts/dev/test.shell for details.
    make stats                  Shows the latest project stats. 
```

# Author

* Marvin Hansen 
* GPG key ID: 210D39BC
* Github key ID: 369D5A0B210D39BC
* GPG Fingerprint: 4B18 F7B2 04B9 7A72 967E  663E 369D 5A0B 210D 39BC 
* Public key: [key](pubkey.txt)

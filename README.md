# typedb-client-go

go client for [TypeDB](https://vaticle.com/typedb)

Supports: TypeDB [2.8](https://github.com/vaticle/typedb/releases/tag/2.8.0)

Protocol: [2.6.1](https://github.com/vaticle/typedb-protocol/releases/tag/2.6.1)

DB Admin

- [x] GetAllDatabases
- [x] CreateDatabase
- [x] CheckDatabaseExists 
- [x] DeleteDatabase

## Make reference

```bash 
Setup: 
    make check                  Checks whether all project requirements are present.
     
Dev: 
    make build                  Builds the code base incrementally (fast).
    make rebuild                Rebuilds all dependencies & the code base (slow). Use after go mod changes. 
    make stats                  Shows the latest project stats. 
```

# Author

* Marvin Hansen 
* GPG key ID: 369D5A0B210D39BC
* GPG Fingerprint: 4B18 F7B2 04B9 7A72 967E  663E 369D 5A0B 210D 39BC 
* Public key: [key](pubkey.txt)

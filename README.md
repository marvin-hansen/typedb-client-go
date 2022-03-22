# go-typedb

go client for [TypeDB](https://vaticle.com/typedb)

Supports: TypeDB 2.7

Implemented:

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
    make stats                  Crunches & shows the latest project stats. 
```
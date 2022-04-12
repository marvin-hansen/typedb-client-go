# Dev status

##  Fully implemented & fully Tested

DBManager

- [x] GetAllDatabases
- [x] CreateDatabase
- [x] CheckDatabaseExists
- [x] DeleteDatabase
- [x] CreateDatabaseSchema
- [x] GetDatabaseSchema

## Fully implemented & tested indirectly

TransactionManager:

- [x] NewTransaction
- [x] OpenTransaction
- [x] ExecuteTransaction
- [x] CommitTransaction
- [x] RollbackTransaction
- [x] CloseTransaction

SessionManager:

- [x] NewSession
- [x] GetSession
- [x] ResetSession
- [x] CloseSession
- [x] MonitorSession (handles heartbeat automatically when open & close a session)

## Implemented - Not fully tested yet

QueryManager:

- [x] RunInsertQuery
- [x] RunInsertBulkQuery
- [x] RunUpdateQuery
- [x] RunMatchQuery
- [x] RunDeleteQuery
- [x] RunExplainQuery
- [x] RunDefineQuery
- [x] RunUnDefineQuery
- [x] RunMatchGroupQuery
- [x] RunMatchAggregateQuery
- [x] RunMatchGroupAggregateQuery


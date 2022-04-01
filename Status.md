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


## Partly implemented & NOT tested 

SessionManager:
- [x] Monitor each session 

## Implemented - needs more testing 

Query:
- [x] RunUpdateQuery

## Implemented - No testing / unknown working 

- [x] RunDeleteQuery
- [x] RunExplainQuery
- [x] RunDefineQuery
- [x] RunUnDefineQuery
- [x] RunMatchQuery
- [x] RunMatchGroupQuery
- [x] RunMatchAggregateQuery
- [x] RunMatchGroupAggregateQuery

## Broken 

- [x] RunInsertQuery

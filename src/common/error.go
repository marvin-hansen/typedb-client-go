package common

import (
	"github.com/pkg/errors"
)

type CustomErrorMessage struct {
	CodePrefix    string
	CodeNumber    uint
	MessagePrefix string
	MessageBody   string
}

func (s CustomErrorMessage) Error() error {
	return errors.Errorf("CodePrefix %v, CodeNumber: %v, MessagePrefix: %v, MessageBody: %v",
		s.CodePrefix, s.CodeNumber, s.MessagePrefix, s.MessageBody)
}

type ClientErrorMessage = CustomErrorMessage

func NewClientErrorMessage(CodeNumber uint, MessageBody string) *ClientErrorMessage {
	return &ClientErrorMessage{
		CodePrefix:    "CLI",
		MessagePrefix: "Client Error: ",
		CodeNumber:    CodeNumber,
		MessageBody:   MessageBody,
	}

}

func TypeDBClientError(ce ClientErrorMessage) error {
	return ce.Error()
}

var (
	CLIENT_CLOSE                      = NewClientErrorMessage(1, "The client has been closed and no further operation is allowed.")
	SESSION_CLOSED                    = NewClientErrorMessage(2, "The session has been closed and no further operation is allowed.")
	TRANSACTION_CLOSED                = NewClientErrorMessage(3, "The transaction has been closed and no further operation is allowed.")
	TRANSACTION_CLOSED_WITH_ERRORS    = NewClientErrorMessage(4, "The transaction has been closed with error(s):\n%s.")
	UNABLE_TO_CONNECT                 = NewClientErrorMessage(5, "Unable to connect to TypeDB server.")
	NEGATIVE_VALUE_NOT_ALLOWED        = NewClientErrorMessage(6, "Value cannot be less than 1, was: '%d'.")
	MISSING_DB_NAME                   = NewClientErrorMessage(7, "Database name cannot be empty.")
	DB_DOES_NOT_EXIST                 = NewClientErrorMessage(8, "The database '%s' does not exist.")
	MISSING_RESPONSE                  = NewClientErrorMessage(9, "Unexpected empty response for request ID '%s'.")
	UNKNOWN_REQUEST_ID                = NewClientErrorMessage(10, "Received a response with unknown request id '%s':\n%s")
	CLUSTER_NO_PRIMARY_REPLICA_YET    = NewClientErrorMessage(11, "No replica has been marked as the primary replica for latest known term '%d'.")
	CLUSTER_UNABLE_TO_CONNECT         = NewClientErrorMessage(12, "Unable to connect to TypeDB Cluster. Attempted connecting to the cluster members, but none are available: '%s'.")
	CLUSTER_REPLICA_NOT_PRIMARY       = NewClientErrorMessage(13, "The replica is not the primary replica.")
	CLUSTER_ALL_NODES_FAILED          = NewClientErrorMessage(14, "Attempted connecting to all cluster members, but the following errors occurred: \n%s")
	CLUSTER_USER_DOES_NOT_EXIST       = NewClientErrorMessage(15, "The user '%s' does not exist.")
	CLUSTER_TOKEN_CREDENTIAL_INVALID  = NewClientErrorMessage(16, "Invalid token credential.")
	CLUSTER_INVALID_ROOT_CA_PATH      = NewClientErrorMessage(17, "The provided Root CA path '%s' does not exist.")
	CLUSTER_CLIENT_CALLED_WITH_STRING = NewClientErrorMessage(18, "The first argument of TypeDBClient.cluster() must be a List of server addresses to connect to. It was called with a string, not a List, which is not allowed.")
)

type ConceptErrorMessage = CustomErrorMessage

func NewConceptErrorMessage(CodeNumber uint, MessageBody string) *ConceptErrorMessage {
	return &ClientErrorMessage{
		CodePrefix:    "CLI",
		MessagePrefix: "Concept Error: ",
		CodeNumber:    CodeNumber,
		MessageBody:   MessageBody,
	}
}

func TypeDBConceptError(ce ConceptErrorMessage) error {
	return ce.Error()
}

var (
	INVALID_CONCEPT_CASTING           = NewConceptErrorMessage(1, "Invalid concept conversion from '%s' to '%s'.")
	MISSING_TRANSACTION               = NewConceptErrorMessage(2, "Transaction cannot be null.")
	MISSING_IID                       = NewConceptErrorMessage(3, "IID cannot be null or empty.")
	MISSING_LABEL                     = NewConceptErrorMessage(4, "Label cannot be null or empty.")
	BAD_ENCODING                      = NewConceptErrorMessage(5, "The encoding '%s' was not recognised.")
	BAD_VALUE_TYPE                    = NewConceptErrorMessage(6, "The value type '%s' was not recognised.")
	BAD_ATTRIBUTE_VALUE               = NewConceptErrorMessage(7, "The attribute value '%s' was not recognised.")
	NONEXISTENT_EXPLAINABLE_CONCEPT   = NewConceptErrorMessage(8, "The concept identified by '%s' is not explainable.")
	NONEXISTENT_EXPLAINABLE_OWNERSHIP = NewConceptErrorMessage(9, "The ownership by owner '%s' of attribute '%s' is not explainable.")
	GET_HAS_WITH_MULTIPLE_FILTERS     = NewConceptErrorMessage(10, "Only one filter can be applied at a time to get_has. The possible filters are: [attribute_type, attribute_types, only_key]")
)

type QueryErrorMessage = CustomErrorMessage

func NewQueryErrorMessage(CodeNumber uint, MessageBody string) *QueryErrorMessage {
	return &ClientErrorMessage{
		CodePrefix:    "CLI",
		MessagePrefix: "Query Error: ",
		CodeNumber:    CodeNumber,
		MessageBody:   MessageBody,
	}
}

func TypeDBQueryError(ce QueryErrorMessage) error {
	return ce.Error()
}

var (
	VARIABLE_DOES_NOT_EXIST = NewQueryErrorMessage(1, "The variable '%s' does not exist.")
	NO_EXPLANATION          = NewQueryErrorMessage(2, "No explanation was found.")
	BAD_ANSWER_TYPE         = NewQueryErrorMessage(3, "The answer type '%s' was not recognised.")
	MISSING_ANSWER          = NewQueryErrorMessage(4, "The required field 'answer' of type was not set.")
)

type InternalErrorMessage = CustomErrorMessage

func NewInternalErrorMessage(CodeNumber uint, MessageBody string) *InternalErrorMessage {
	return &ClientErrorMessage{
		CodePrefix:    "CLI",
		MessagePrefix: "Internal Error: ",
		CodeNumber:    CodeNumber,
		MessageBody:   MessageBody,
	}
}

func TypeDBInternalErrorMessage(ce InternalErrorMessage) error {
	return ce.Error()
}

var (
	ILLEGAL_STATE    = NewInternalErrorMessage(2, "Illegal state has been reached!")
	ILLEGAL_ARGUMENT = NewInternalErrorMessage(3, "Illegal argument provided")
	ILLEGAL_CAST     = NewInternalErrorMessage(4, "Illegal casting operation.")
)

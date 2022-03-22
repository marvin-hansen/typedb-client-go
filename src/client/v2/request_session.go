package v2

import "github.com/marvin-hansen/go-typedb/proto/common"

// Session
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

func getSessionOpenReq(dbName string, sessionType common.Session_Type, options *common.Options) (req *common.Session_Open_Req) {
	return &common.Session_Open_Req{
		Database: dbName,
		Type:     sessionType,
		Options:  options,
	}
}

func getSessionPulseReq(sessionID []byte) (req *common.Session_Pulse_Req) {
	return &common.Session_Pulse_Req{SessionId: sessionID}
}

func getSessionCloseReq(sessionID []byte) (req *common.Session_Close_Req) {
	return &common.Session_Close_Req{SessionId: sessionID}
}

package v2

import (
	common2 "github.com/marvin-hansen/go-typedb/common"
)

// Session
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

func getSessionOpenReq(dbName string, sessionType common2.Session_Type, options *common2.Options) (req *common2.Session_Open_Req) {
	return &common2.Session_Open_Req{
		Database: dbName,
		Type:     sessionType,
		Options:  options,
	}
}

func getSessionPulseReq(sessionID []byte) (req *common2.Session_Pulse_Req) {
	return &common2.Session_Pulse_Req{SessionId: sessionID}
}

func getSessionCloseReq(sessionID []byte) (req *common2.Session_Close_Req) {
	return &common2.Session_Close_Req{SessionId: sessionID}
}

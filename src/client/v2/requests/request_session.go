// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package requests

import (
	"github.com/marvin-hansen/typedb-client-go/common"
)

// Session
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

func GetSessionOpenReq(dbName string, sessionType common.Session_Type, options *common.Options) (req *common.Session_Open_Req) {
	return &common.Session_Open_Req{
		Database: dbName,
		Type:     sessionType,
		Options:  options,
	}
}

func GetSessionPulseReq(sessionID []byte) (req *common.Session_Pulse_Req) {
	return &common.Session_Pulse_Req{SessionId: sessionID}
}

func GetSessionCloseReq(sessionID []byte) (req *common.Session_Close_Req) {
	return &common.Session_Close_Req{SessionId: sessionID}
}

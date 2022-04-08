package utils

import "encoding/hex"

func ByteToString(sessionID []byte) string {
	return hex.EncodeToString(sessionID)
}

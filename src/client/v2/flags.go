// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

const (
	debug   = false
	main    = "TypeDBClient: "
	connErr = main + "error: Connection is nil; can't create client"
)

func dbgPrint(mtd, msg string) {
	if debug {
		println("[" + mtd + "]: " + msg)
	}
}

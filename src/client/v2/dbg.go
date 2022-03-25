// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"google.golang.org/grpc"
	"log"
)

func dbgPrint(dbg bool, msg string) {
	if dbg {
		log.Println(msg)
	}
}

func dbgCheckPrintLog(dbg bool, main, mtd, msg string, err error) {
	if dbg {
		if err != nil {
			log.Print("Main: ", main)
			log.Print("Mtd: ", mtd)
			log.Print("Msg: ", msg)
			log.Print("Error: ", err.Error())
		}
	}
}

func checkPrintErr(err error, errorMsg string) {
	if err != nil {
		log.Println("Error:", err.Error())
		log.Println("Error Message: ", errorMsg)
	}
}

func checkConnection(conn *grpc.ClientConn, err error) {
	if err != nil {
		checkPrintErr(err, connErr)
	}
	if conn == nil {
		log.Fatal(connErr)
	}
}

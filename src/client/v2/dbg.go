package v2

import (
	"google.golang.org/grpc"
	"log"
	"os"
)

func DbgPrint(dbg bool, msg string) {
	if dbg {
		log.Println(msg)
	}
}

func DbgCheckPrintLog(dbg bool, main, mtd, msg string, err error) {
	if dbg {
		if err != nil {
			log.Print("Main: ", main)
			log.Print("Mtd: ", mtd)
			log.Print("Msg: ", msg)
			log.Print("Error: ", err.Error())
		}
	}
}

func CheckPrintErr(err error, errorMsg string) {
	if err != nil {
		log.Println("Error:", err.Error())
		log.Println("Error Message: ", errorMsg)
	}
}

func CheckPrintErrStop(err error, errorMsg string) {
	if err != nil {
		log.Println("Error:", err.Error())
		log.Println("Error Message: ", errorMsg)
		os.Exit(42)
	}
}

func CheckConnection(conn *grpc.ClientConn, err error) {
	if err != nil {
		CheckPrintErr(err, connErr)
	}
	if conn == nil {
		log.Fatal(connErr)
	}
}

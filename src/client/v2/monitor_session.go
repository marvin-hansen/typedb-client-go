package v2

import (
	"fmt"
	"time"
)

func (s SessionManager) stopMonitorSession() {

}

func (s SessionManager) startMonitorSession() {
	//https://medium.com/@marcus.murray/go-closing-routines-f90097f35223

}

func heartbeat() {
	go func() {
		timer := time.After(time.Second * 10)
		<-timer
		fmt.Println("heartbeat happened!")
	}()
}

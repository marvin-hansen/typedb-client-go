package v2

import (
	"fmt"
	"time"
)

func (s SessionManager) MonitorSession() {

}

func heartbeat() {
	go func() {
		timer := time.After(time.Second * 10)
		<-timer
		fmt.Println("heartbeat happened!")
	}()
}

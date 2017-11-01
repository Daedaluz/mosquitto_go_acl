package main

import(
	"fmt"
	"time"
)

var (
	exit chan bool
)

func Ticker() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-exit:
			fmt.Println("Time to quit!")
			return
		case tick := <-ticker.C:
			fmt.Println("Tick!", tick)
		}
	}
}

func PluginInit(opts map[string]string) bool {
	Log(MOSQ_LOG_ERR, "Hello world!")
	for key, value := range opts {
		fmt.Println(key, value)
	}
	exit = make(chan bool)
	go Ticker()
	return true
}

func PluginCleanup(opts map[string]string) bool {
	fmt.Println("PluginCleanup")
	exit <- true
	return true
}

func SecurityInit(opts map[string]string, reload bool) bool {
	fmt.Println("SecurityInit")
	return true
}

func SecurityCleanup(opts map[string]string, reload bool) bool {
	fmt.Println("SecurityCleanup")
	return true
}

func ACLCheck(client_id, username, topic string, access int) bool {
	fmt.Println("ACL Check", client_id, username, topic, access)
	return true
}

func UnpwdCheck(username, password string) int {
	fmt.Println("Unpwd Check", username, password)
	return MOSQ_ERR_SUCCESS
}


package main

import ()

var ()

func PluginInit(argopts map[string]string) bool {
	Logf(MOSQ_LOG_INFO, "Hello from PluginInit")
	for option, value := range argopts {
		Logf(MOSQ_LOG_INFO, "Option %s = %s", option, value)
	}
	return true
}

func PluginCleanup(opts map[string]string) bool {
	Logf(MOSQ_LOG_INFO, "Bye from PluginCleanup")
	return true
}

func ACLCheck(client *Client, access Access, topic string, payload []byte, qos int, retain bool) bool {
	Logf(MOSQ_LOG_INFO, "ACLCheck %s: %v %v (q%v,r%v)", access, topic, payload, qos, retain)
	switch(access) {
	case MOSQ_ACL_READ:
	case MOSQ_ACL_WRITE:
	case MOSQ_ACL_SUBSCRIBE:
	}
	return true
}

func UnpwdCheck(client *Client, username, password string) int {
	Logf(MOSQ_LOG_INFO, "User %s was granted acces", username)
	return MOSQ_ERR_SUCCESS
	// return MOSQ_ERR_AUTH
}

func SecurityInit(opts map[string]string, reload bool) bool {
	Logf(MOSQ_LOG_INFO, "Security init, reload = %v", reload)
	return true
}

func SecurityCleanup(opts map[string]string, reload bool) bool {
	Logf(MOSQ_LOG_INFO, "Security cleanup, reload = %v", reload)
	return true
}

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

func ACLCheck(client *Client, access int, topic string, payload []byte, qos int, retain bool) bool {
	Logf(MOSQ_LOG_INFO, "ACL Check %s", topic)
	Logf(MOSQ_LOG_INFO, "%v (%v;%v)", payload, qos, retain)
	Logf(MOSQ_LOG_INFO, "IP: %v", client.Address())
	Logf(MOSQ_LOG_INFO, "Clean: %v", client.CleanSession())
	Logf(MOSQ_LOG_INFO, "ClientId: %v", client.ClientId())
	Logf(MOSQ_LOG_INFO, "KeepAlive: %v", client.KeepAlive())
	Logf(MOSQ_LOG_INFO, "Protocol: %v", client.Protocol())
	Logf(MOSQ_LOG_INFO, "SubCount: %v", client.SubCount())
	Logf(MOSQ_LOG_INFO, "Username: %v", client.Username())
	Logf(MOSQ_LOG_INFO, "--------------------")
	return true
}

func UnpwdCheck(client *Client, username, password string) int {
	Logf(MOSQ_LOG_INFO, "User %s was granted acces", username)
	Logf(MOSQ_LOG_INFO, "IP: %v", client.Address())
	Logf(MOSQ_LOG_INFO, "Clean: %v", client.CleanSession())
	Logf(MOSQ_LOG_INFO, "ClientId: %v", client.ClientId())
	Logf(MOSQ_LOG_INFO, "KeepAlive: %v", client.KeepAlive())
	Logf(MOSQ_LOG_INFO, "Protocol: %v", client.Protocol())
	Logf(MOSQ_LOG_INFO, "SubCount: %v", client.SubCount())
	Logf(MOSQ_LOG_INFO, "Username: %v", client.Username())
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

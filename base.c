#define bool _Bool
#include <mosquitto_plugin.h>
#include "auth.h"
int mosquitto_auth_acl_check(void* udata, const char* client_id, const char* username, const char* topic, int access) {
	return go_mosquitto_auth_acl_check(udata, (char*)client_id, (char*)username, (char*)topic, access);
}

int mosquitto_auth_unpwd_check(void* udata, const char* username, const char* password) {
	return go_mosquitto_auth_unpwd_check(udata, (char*)username, (char*)password);
}

int mosquitto_auth_psk_key_get(void* udata, const char* hint, const char* identity, char* key, int max_key_len) {
	return go_mosquitto_auth_psk_key_get(udata, (char*)hint, (char*)identity, key, max_key_len);
}

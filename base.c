#define bool _Bool
#include <stdio.h>
#include <mosquitto.h>
#include <mosquitto_plugin.h>

extern void mosquitto_log_printf(int, const char*, ...);

extern int go_mosquitto_auth_acl_check(int, struct mosquitto*, struct mosquitto_acl_msg*);
extern int go_mosquitto_auth_unpwd_check(struct mosquitto*, char*, char*);
extern int go_mosquitto_auth_psk_key_get(struct mosquitto*, char*, char*, char*, int);

void mosquitto_log(int lvl, char* str) {
	mosquitto_log_printf(lvl, "%s", str);
}

bool topic_match(char* sub, char* topic) {
	bool res;
	mosquitto_topic_matches_sub(sub, topic, &res);
	return res;
}

int mosquitto_auth_acl_check(void* udata, int access, const struct mosquitto* client, const struct mosquitto_acl_msg *msg) {
	return go_mosquitto_auth_acl_check(access, (struct mosquitto*)client, (struct mosquitto_acl_msg*)msg);
}

int mosquitto_auth_unpwd_check(void* udata, const struct mosquitto *client, const char* username, const char* password) {
	return go_mosquitto_auth_unpwd_check((struct mosquitto*) client, (char*)username, (char*)password);
}

int mosquitto_auth_psk_key_get(void* udata, const struct mosquitto *client, const char* hint, const char* identity, char* key, int max_key_len) {
	return go_mosquitto_auth_psk_key_get((struct mosquitto*)client, (char*)hint, (char*)identity, key, max_key_len);
}


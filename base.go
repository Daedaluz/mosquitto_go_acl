package main
/*
#define bool _Bool
#include <mosquitto.h>
#include <mosquitto_plugin.h>
*/
import "C"

import(
	"unsafe"
	"reflect"
)

const(
	MOSQ_ACL_READ = C.MOSQ_ACL_READ
	MOSQ_ACL_WRITE = C.MOSQ_ACL_WRITE
	MOSQ_ACL_NONE = C.MOSQ_ACL_NONE
)

const (
	MOSQ_ERR_CONN_PENDING = C.MOSQ_ERR_CONN_PENDING
	MOSQ_ERR_SUCCESS = C.MOSQ_ERR_SUCCESS
	MOSQ_ERR_NOMEM = C.MOSQ_ERR_NOMEM
	MOSQ_ERR_PROTOCOL = C.MOSQ_ERR_PROTOCOL
	MOSQ_ERR_INVAL = C.MOSQ_ERR_INVAL
	MOSQ_ERR_NO_CONN = C.MOSQ_ERR_NO_CONN
	MOSQ_ERR_CONN_REFUSED = C.MOSQ_ERR_CONN_REFUSED
	MOSQ_ERR_NOT_FOUND = C.MOSQ_ERR_NOT_FOUND
	MOSQ_ERR_CONN_LOST = C.MOSQ_ERR_CONN_LOST
	MOSQ_ERR_TLS = C.MOSQ_ERR_TLS
	MOSQ_ERR_PAYLOAD_SIZE = C.MOSQ_ERR_PAYLOAD_SIZE
	MOSQ_ERR_NOT_SUPPORTED = C.MOSQ_ERR_NOT_SUPPORTED
	MOSQ_ERR_AUTH = C.MOSQ_ERR_AUTH
	MOSQ_ERR_ACL_DENIED = C.MOSQ_ERR_ACL_DENIED
	MOSQ_ERR_UNKNOWN = C.MOSQ_ERR_UNKNOWN
	MOSQ_ERR_ERRNO = C.MOSQ_ERR_ERRNO
	MOSQ_ERR_EAI = C.MOSQ_ERR_EAI
	MOSQ_ERR_PROXY = C.MOSQ_ERR_PROXY
)

//export mosquitto_auth_plugin_version
func mosquitto_auth_plugin_version() C.int {
	return C.MOSQ_AUTH_PLUGIN_VERSION
}

func toOptSlice(opts *C.struct_mosquitto_auth_opt, optcount C.int) []C.struct_mosquitto_auth_opt {
	var Opts []C.struct_mosquitto_auth_opt

	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&Opts)))
	sliceHeader.Cap = int(optcount)
	sliceHeader.Len = int(optcount)
	sliceHeader.Data = uintptr(unsafe.Pointer(opts))
	return Opts
}

func optSliceToMap(slice []C.struct_mosquitto_auth_opt) map[string]string {
	optmap := map[string]string{}
	for _, opt := range slice {
		optmap[C.GoString(opt.key)] = C.GoString(opt.value)
	}
	return optmap
}

//export mosquitto_auth_plugin_init
func mosquitto_auth_plugin_init(userdata *unsafe.Pointer, opts *C.struct_mosquitto_auth_opt, optcount C.int) C.int {
	optslice := toOptSlice(opts, optcount)
	optmap := optSliceToMap(optslice)

	x := PluginInit(optmap)
	if x {
		return C.int(0)
	}
	return C.int(1)
}

//export mosquitto_auth_plugin_cleanup
func mosquitto_auth_plugin_cleanup(userdata unsafe.Pointer, opts *C.struct_mosquitto_auth_opt, optcount C.int) C.int {
	optslice := toOptSlice(opts, optcount)
	optmap := optSliceToMap(optslice)
	x := PluginCleanup(optmap)
	if x {
		return C.int(0)
	}
	return C.int(1)
}

//export mosquitto_auth_security_init
func mosquitto_auth_security_init(userdata unsafe.Pointer, opts *C.struct_mosquitto_auth_opt, optcount C.int, reload C.bool) C.int {
	optslice := toOptSlice(opts, optcount)
	optmap := optSliceToMap(optslice)

	x := SecurityInit(optmap, bool(reload))
	if x {
		return C.int(0)
	}
	return C.int(1)
}

//export mosquitto_auth_security_cleanup
func mosquitto_auth_security_cleanup(userdata unsafe.Pointer, opts *C.struct_mosquitto_auth_opt, optcount C.int, reload C.bool) C.int {
	optslice := toOptSlice(opts, optcount)
	optmap := optSliceToMap(optslice)

	x := SecurityCleanup(optmap, bool(reload))
	if x {
		return C.int(0)
	}
	return C.int(1)
}

//export go_mosquitto_auth_acl_check
func go_mosquitto_auth_acl_check(userdata unsafe.Pointer, client_id, username, topic *C.char, access C.int) C.int {
	x := ACLCheck(C.GoString(client_id), C.GoString(username), C.GoString(topic), int(access))
	if x {
		return C.int(0)
	}
	return C.int(1)
}

//export go_mosquitto_auth_unpwd_check
func go_mosquitto_auth_unpwd_check(userdata unsafe.Pointer, username, password *C.char) C.int {
	x := UnpwdCheck(C.GoString(username), C.GoString(password))
	return C.int(x)
}

//export go_mosquitto_auth_psk_key_get
func go_mosquitto_auth_psk_key_get(userdata unsafe.Pointer, hint, identity, key *C.char, max_key_len C.int) C.int {
	panic("AUTH_PSK_KEY_GET IS UNIMPLEMENTED")
	return C.int(0)
}

func main() {}

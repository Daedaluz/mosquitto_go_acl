package main

//#cgo LDFLAGS: -Wl,-unresolved-symbols=ignore-all -lmosquitto
/*
#define bool _Bool
#include <malloc.h>
#include <mosquitto.h>
#include <mosquitto_plugin.h>

void mosquitto_log(int lvl, char* msg);
bool topic_match(char* sub, char* topic);

extern char* mosquitto_client_address(const struct mosquitto *client);
extern bool mosquitto_client_clean_session(const struct mosquitto *client);
extern char *mosquitto_client_id(const struct mosquitto *client);
extern int mosquitto_client_keepalive(const struct mosquitto *client);
extern int mosquitto_client_protocol(const struct mosquitto *client);
extern int mosquitto_client_sub_count(const struct mosquitto *client);
extern char* mosquitto_client_username(const struct mosquitto *context);

*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

const (
	MOSQ_ACL_READ      = C.MOSQ_ACL_READ
	MOSQ_ACL_WRITE     = C.MOSQ_ACL_WRITE
	MOSQ_ACL_NONE      = C.MOSQ_ACL_NONE
	MOSQ_ACL_SUBSCRIBE = C.MOSQ_ACL_SUBSCRIBE
)

var (
	_aclMap = map[int]string{
		MOSQ_ACL_READ:      "READ",
		MOSQ_ACL_WRITE:     "WRITE",
		MOSQ_ACL_NONE:      "NONE",
		MOSQ_ACL_SUBSCRIBE: "SUB",
	}
)

type Access int

func (a Access) String() string {
	return _aclMap[int(a)]
}

const (
	MOSQ_ERR_SUCCESS       = C.MOSQ_ERR_SUCCESS
	MOSQ_ERR_PROTOCOL      = C.MOSQ_ERR_PROTOCOL
	MOSQ_ERR_NOT_SUPPORTED = C.MOSQ_ERR_NOT_SUPPORTED
	MOSQ_ERR_AUTH          = C.MOSQ_ERR_AUTH
	MOSQ_ERR_ACL_DENIED    = C.MOSQ_ERR_ACL_DENIED
	MOSQ_ERR_PLUGIN_DEFER  = C.MOSQ_ERR_PLUGIN_DEFER
	MOSQ_ERR_UNKNOWN       = C.MOSQ_ERR_UNKNOWN
)

const (
	MOSQ_LOG_INFO    = C.MOSQ_LOG_INFO
	MOSQ_LOG_NOTICE  = C.MOSQ_LOG_NOTICE
	MOSQ_LOG_WARNING = C.MOSQ_LOG_WARNING
	MOSQ_LOG_ERR     = C.MOSQ_LOG_ERR
	MOSQ_LOG_DEBUG   = C.MOSQ_LOG_DEBUG
)

// Helper functions

func toOptSlice(opts *C.struct_mosquitto_opt, optcount C.int) []C.struct_mosquitto_opt {
	var Opts []C.struct_mosquitto_opt

	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&Opts)))
	sliceHeader.Cap = int(optcount)
	sliceHeader.Len = int(optcount)
	sliceHeader.Data = uintptr(unsafe.Pointer(opts))
	return Opts
}

func optSliceToMap(slice []C.struct_mosquitto_opt) map[string]string {
	optmap := map[string]string{}
	for _, opt := range slice {
		optmap[C.GoString(opt.key)] = C.GoString(opt.value)
	}
	return optmap
}

// Expose log functions within mosquitto
func Log(lvl int, msg string) {
	cmsg := C.CString(msg)
	C.mosquitto_log(C.int(lvl), cmsg)
	C.free(unsafe.Pointer(cmsg))
}

func Logf(lvl int, msg string, v ...interface{}) {
	fmsg := fmt.Sprintf(msg, v...)
	cmsg := C.CString(fmsg)
	C.mosquitto_log(C.int(lvl), cmsg)
	C.free(unsafe.Pointer(cmsg))
}

// Use mosquittos own pattern matching
func Match(t, sub string) bool {
	c_t := C.CString(t)
	c_sub := C.CString(sub)
	res := C.topic_match(c_sub, c_t)
	C.free(unsafe.Pointer(c_t))
	C.free(unsafe.Pointer(c_sub))
	return bool(res)
}

// AUTH implenetation

//export mosquitto_auth_plugin_version
func mosquitto_auth_plugin_version() C.int {
	return C.MOSQ_AUTH_PLUGIN_VERSION
}

//export mosquitto_auth_plugin_init
func mosquitto_auth_plugin_init(userdata *unsafe.Pointer, opts *C.struct_mosquitto_opt, optcount C.int) C.int {
	optslice := toOptSlice(opts, optcount)
	optmap := optSliceToMap(optslice)

	x := PluginInit(optmap)
	if x {
		return C.int(0)
	}
	return C.int(1)
}

//export mosquitto_auth_plugin_cleanup
func mosquitto_auth_plugin_cleanup(userdata unsafe.Pointer, opts *C.struct_mosquitto_opt, optcount C.int) C.int {
	optslice := toOptSlice(opts, optcount)
	optmap := optSliceToMap(optslice)
	x := PluginCleanup(optmap)
	if x {
		return C.int(0)
	}
	return C.int(1)
}

//export mosquitto_auth_security_init
func mosquitto_auth_security_init(userdata unsafe.Pointer, opts *C.struct_mosquitto_opt, optcount C.int, reload C.bool) C.int {
	optslice := toOptSlice(opts, optcount)
	optmap := optSliceToMap(optslice)

	x := SecurityInit(optmap, bool(reload))
	if x {
		return C.int(0)
	}
	return C.int(1)
}

//export mosquitto_auth_security_cleanup
func mosquitto_auth_security_cleanup(userdata unsafe.Pointer, opts *C.struct_mosquitto_opt, optcount C.int, reload C.bool) C.int {
	optslice := toOptSlice(opts, optcount)
	optmap := optSliceToMap(optslice)

	x := SecurityCleanup(optmap, bool(reload))
	if x {
		return C.int(0)
	}
	return C.int(1)
}

//export go_mosquitto_auth_acl_check
func go_mosquitto_auth_acl_check(access C.int, client *C.struct_mosquitto, msg *C.struct_mosquitto_acl_msg) int {
	c := &Client{
		c: client,
	}

	var Payload []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&Payload)))
	sliceHeader.Cap = int(msg.payloadlen)
	sliceHeader.Len = int(msg.payloadlen)
	sliceHeader.Data = uintptr(unsafe.Pointer(msg.payload))

	x := ACLCheck(c, Access(access), C.GoString(msg.topic), Payload, int(msg.qos), bool(msg.retain))
	if x {
		return C.MOSQ_ERR_SUCCESS
	} else {
		return C.MOSQ_ERR_ACL_DENIED
	}
}

//export go_mosquitto_auth_unpwd_check
func go_mosquitto_auth_unpwd_check(client *C.struct_mosquitto, username, password *C.char) C.int {
	c := &Client{
		c: client,
	}
	x := UnpwdCheck(c, C.GoString(username), C.GoString(password))
	return C.int(x)
}

//export go_mosquitto_auth_psk_key_get
func go_mosquitto_auth_psk_key_get(client *C.struct_mosquitto, hint, identity, key *C.char, max_key_len C.int) C.int {
	// TODO: map call to some function in main.go
	return C.MOSQ_ERR_PLUGIN_DEFER
}

// Expose client information
type Client struct {
	c *C.struct_mosquitto
}

func (c *Client) Address() string {
	cstr := C.mosquitto_client_address(c.c)
	return C.GoString(cstr)
}

func (c *Client) CleanSession() bool {
	x := C.mosquitto_client_clean_session(c.c)
	return bool(x)
}

func (c *Client) ClientId() string {
	cstr := C.mosquitto_client_id(c.c)
	return C.GoString(cstr)
}

func (c *Client) KeepAlive() int {
	x := C.mosquitto_client_keepalive(c.c)
	return int(x)
}

// Perhaps implement mosquitto_client_certificate here.. some day...

func (c *Client) Protocol() int {
	x := C.mosquitto_client_protocol(c.c)
	return int(x)
}

func (c *Client) SubCount() int {
	x := C.mosquitto_client_sub_count(c.c)
	return int(x)
}

func (c *Client) Username() string {
	cstr := C.mosquitto_client_username(c.c)
	return C.GoString(cstr)
}

func main() {}

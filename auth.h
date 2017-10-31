/* Created by "go tool cgo" - DO NOT EDIT. */

/* package _/home/tux/mqtt_plugin */

/* Start of preamble from import "C" comments.  */


#line 2 "/home/tux/mqtt_plugin/base.go"

#define bool _Bool
#include <mosquitto.h>
#include <mosquitto_plugin.h>

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

typedef struct { const char *p; GoInt n; } GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


extern int mosquitto_auth_plugin_version();

extern int mosquitto_auth_plugin_init(void** p0, struct mosquitto_auth_opt* p1, int p2);

extern int mosquitto_auth_plugin_cleanup(void* p0, struct mosquitto_auth_opt* p1, int p2);

extern int mosquitto_auth_security_init(void* p0, struct mosquitto_auth_opt* p1, int p2, _Bool p3);

extern int mosquitto_auth_security_cleanup(void* p0, struct mosquitto_auth_opt* p1, int p2, _Bool p3);

extern int go_mosquitto_auth_acl_check(void* p0, char* p1, char* p2, char* p3, int p4);

extern int go_mosquitto_auth_unpwd_check(void* p0, char* p1, char* p2);

extern int go_mosquitto_auth_psk_key_get(void* p0, char* p1, char* p2, char* p3, int p4);

#ifdef __cplusplus
}
#endif

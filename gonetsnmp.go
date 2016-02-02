package gonetsnmp

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L/usr/local/lib -lnetsnmp
#include <stdlib.h>
#include "net-snmp-5.7.3/include/net-snmp/net-snmp-config.h"
#include "net-snmp-5.7.3/include/net-snmp/net-snmp-features.h"
#include <snmp.h>

*/
import "C"
import (
	// "fmt"
	"log"
	"unsafe"
)

const (
	MAX_TYPE_NAME_LEN int = 32
	MAX_OID_LEN       int = 128
	STR_BUF_SIZE      int = (MAX_TYPE_NAME_LEN * MAX_OID_LEN)
)

func InitSnmp(name string) {
	log.Println("INIT SNMP")

	var cmsg *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cmsg))
	C.init_snmp(cmsg)
}

func ShutdownSnmp(name string) {
	var cmsg *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cmsg))
	C.snmp_shutdown(cmsg)
}

func EnableMibWarnings(x C.int) {
	C.snmp_set_mib_warnings(x)
}

func EnableMibError(x C.int) {
	C.snmp_set_mib_errors(x)
}

func PrintTree() {
	C.print_tree()
}

func SetMibDirectory(dir string) {
	var cmsg *C.char = C.CString(dir)
	defer C.free(unsafe.Pointer(cmsg))
	C.netsnmp_set_mib_directory(cmsg)
}

func AddMibDirectory(dir string) {
	var cmsg *C.char = C.CString(dir)
	defer C.free(unsafe.Pointer(cmsg))
	var res C.int = C.add_mibdir(cmsg)
	if res > 0 {
		log.Printf("DIR ADD '%v' : SUCCESS", dir)
	} else {
		log.Printf("DIR ADD '%v' : ERROR", dir)
	}
}

func UnloadAllMibs() {
	C.unload_all_mibs()
}

func UnloadModule(name string) {
	var cmsg *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cmsg))
	var res C.int = C.netsnmp_unload_module(cmsg)
	if res > 0 {
		log.Printf("MODULE UNLOAD '%v' : SUCCESS", name)
	} else {
		log.Printf("MODULE UNLOAD '%v' : ERROR", name)
	}
}

func InitMibs() {
	C.netsnmp_init_mib()
}

func ReadModule(name string) {
	var cmsg *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(cmsg))
	C.netsnmp_read_module(cmsg)
}

func ReadAllMibs() {
	C.read_all_mibs()
}

func GetEnv(envname string) string {
	var (
		c_envname *C.char = C.CString(envname)
		c_envval  *C.char
		res       string
	)

	defer C.free(unsafe.Pointer(c_envname))
	defer C.free(unsafe.Pointer(c_envval))

	c_envval = C.netsnmp_getenv(c_envname)
	res = C.GoString(c_envval)
	return res
}

func SetEnv(envname string, envval string, x C.int) {
	var (
		c_envname *C.char = C.CString(envname)
		c_envval  *C.char = C.CString(envval)
	)

	defer C.free(unsafe.Pointer(c_envname))
	defer C.free(unsafe.Pointer(c_envval))

	C.netsnmp_setenv(c_envname, c_envval, x)
}

// todo : case for output
func DsSetInt() {
	// C.NETSNMP_OID_OUTPUT_FULL   ".iso.org.dod.internet.private.enterprises.doremi.dcp.machine.softwareVersion"
	// C.NETSNMP_OID_OUTPUT_SUFFIX "softwareVersion"
	// C.NETSNMP_OID_OUTPUT_MODULE "DOREMI-MIB::softwareVersion"

	C.netsnmp_ds_set_int(C.NETSNMP_DS_LIBRARY_ID, C.NETSNMP_DS_LIB_OID_OUTPUT_FORMAT, C.NETSNMP_OID_OUTPUT_SUFFIX)
}

//to divide to tag2oid / oid2tag
func TranslateObj(OID []uint, long bool) string {

	var buf [STR_BUF_SIZE]C.char
	var res string

	C.snprint_objid(&buf[0], C.size_t(len(buf)), (*C.oid)(unsafe.Pointer(&OID[0])), C.size_t(len(OID)))

	res = C.GoString(&buf[0])

	return res
}

func Str2oid(s string) []uint {
	var res []uint
	intRes := strings.Split(s, ".")
	for _, val := range intRes {
		v, _ := strconv.Atoi(val)
		res = append(res, uint(v))
	}

	return res[1:]
}

func Test(OID []uint) {

	// type SIZE_T C.size_t

	// var val *C.char
	// var name_length SIZE_T

	// C.read_objid(val, (*C.oid)(unsafe.Pointer(&OID[0])), &SIZE_T)
}

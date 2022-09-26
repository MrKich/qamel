package main

// #include <stdlib.h>
// #include <stdint.h>
// #include <stdbool.h>
// #include <string.h>
// #include "qamel-backend.h"
import "C"
import (
	"github.com/MrKich/qamel"
	"unsafe"
)

//export qamelBackEndConstructor
func qamelBackEndConstructor(ptr unsafe.Pointer) {
	obj := &BackEnd{}
	obj.Ptr = ptr
	qamel.RegisterObject(ptr, obj)
	obj.init()
}

//export qamelDestroyBackEnd
func qamelDestroyBackEnd(ptr unsafe.Pointer) {
	qamel.DeleteObject(ptr)
}

// getter and setter

// signals invoker

func (obj *BackEnd) timeChanged(p0 string) {
	if obj.Ptr == nil || !qamel.ObjectExists(obj.Ptr) {
		return
	}

	cP0 := C.CString(p0)
	defer C.free(unsafe.Pointer(cP0))
	C.BackEnd_TimeChanged(obj.Ptr, cP0)
}

// RegisterQmlBackEnd registers BackEnd as QML object
func RegisterQmlBackEnd(uri string, versionMajor int, versionMinor int, qmlName string) {
	cURI := C.CString(uri)
	cQmlName := C.CString(qmlName)
	cVersionMajor := C.int(int32(versionMajor))
	cVersionMinor := C.int(int32(versionMinor))
	defer func() {
		C.free(unsafe.Pointer(cURI))
		C.free(unsafe.Pointer(cQmlName))
	}()

	C.BackEnd_RegisterQML(cURI, cVersionMajor, cVersionMinor, cQmlName)
}

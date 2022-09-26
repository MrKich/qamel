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
}

//export qamelDestroyBackEnd
func qamelDestroyBackEnd(ptr unsafe.Pointer) {
	qamel.DeleteObject(ptr)
}

//export qamelBackEndGetRandomNumber
func qamelBackEndGetRandomNumber(ptr unsafe.Pointer) (result C.int) {
	obj := qamel.BorrowObject(ptr)
	defer qamel.ReturnObject(ptr)
	if obj == nil {
		return
	}

	objBackEnd, ok := obj.(*BackEnd)
	if !ok {
		return
	}

	result = C.int(int32(objBackEnd.getRandomNumber()))
	return
}

// getter and setter

// signals invoker

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

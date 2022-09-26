package qamel

// #include <stdint.h>
// #include <stdlib.h>
// #include <string.h>
// #include <stdbool.h>
// #include "qobject.h"
import "C"
import "unsafe"

type QObject struct {
	ptr unsafe.Pointer
}

func (qobject QObject) FindChild(name string) QObject {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	ptr := C.QObject_FindChild(qobject.ptr, cName)
	return QObject{ptr: ptr}
}

func (qobject QObject) GetName() string {
	return C.GoString(C.QObject_GetObjectName(qobject.ptr))
}

func (qobject QObject) SetProperty(name string, value string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	return C.QObject_SetProperty(qobject.ptr, cName, cValue) == true
}

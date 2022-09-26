#pragma once

#ifndef QAMEL_QOBJECT_H
#define QAMEL_QOBJECT_H

#ifdef __cplusplus
extern "C" {
#endif

// Methods
void* QObject_FindChild(void* ptr, char* name);
const char* QObject_GetObjectName(void* ptr);
bool QObject_SetProperty(void* ptr, char* name, char* value);

#ifdef __cplusplus
}
#endif

#endif
#pragma once

#ifndef QAMEL_BACKEND_H
#define QAMEL_BACKEND_H

#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus

// Class
class BackEnd;

extern "C" {
#endif

// Properties

// Signals

// Register
void BackEnd_RegisterQML(char* uri, int versionMajor, int versionMinor, char* qmlName);

#ifdef __cplusplus
}
#endif

#endif

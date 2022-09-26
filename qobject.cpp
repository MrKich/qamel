#include "qobject.h"
#include <QObject>
#include <QString>
#include <QVariant>


void* QObject_FindChild(void* ptr, char* name) {
    QObject *object = static_cast<QObject*>(ptr);
    return object->findChild<QObject*>(QString(name));
}

const char* QObject_GetObjectName(void* ptr) {
    QObject *object = static_cast<QObject*>(ptr);
    return object->objectName().toStdString().c_str();
}

bool QObject_SetProperty(void* ptr, char* name, char* value) {
    QObject *object = static_cast<QObject*>(ptr);
    return object->setProperty(name, QVariant(value));
}
#include <QObject>
#include <QQuickItem>
#include <QString>
#include <QByteArray>
#include <QQmlEngine>
#include <QMetaObject>

#include "_cgo_export.h"
#include "qamel-backend.h"

class BackEnd : public QQuickItem {
	Q_OBJECT

private:

public:
	BackEnd(QQuickItem* parent=Q_NULLPTR) : QQuickItem(parent) {
		qamelBackEndConstructor(this);
	}

	~BackEnd() {
		qamelDestroyBackEnd(this);
	}

signals:
	void timeChanged(QString p0);

public slots:
};

void BackEnd_TimeChanged(void* ptr, char* p0) {
	BackEnd *obj = static_cast<BackEnd*>(ptr);
	obj->timeChanged(QString(p0));
}

void BackEnd_RegisterQML(char* uri, int versionMajor, int versionMinor, char* qmlName) {
	qmlRegisterType<BackEnd>(uri, versionMajor, versionMinor, qmlName);
}

#include "moc-qamel-backend.h"

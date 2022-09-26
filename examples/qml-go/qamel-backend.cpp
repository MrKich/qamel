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

public slots:
	int getRandomNumber() {
		return qamelBackEndGetRandomNumber(this);
	}
};


void BackEnd_RegisterQML(char* uri, int versionMajor, int versionMinor, char* qmlName) {
	qmlRegisterType<BackEnd>(uri, versionMajor, versionMinor, qmlName);
}

#include "moc-qamel-backend.h"

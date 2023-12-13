package upload

import (
	"medsecurity/cmd/webservice/router/ping/handler"
	"medsecurity/config"
	"medsecurity/service/ping"

	"github.com/labstack/echo/v4"
)

func InitHandler(
	e *echo.Echo,
	fileService ping.ServiceItf,
	config config.Config,
) {
	pingHandler := handler.InitUploadHandler(e, fileService)

	pingRouterV1 := e.Group(V1PingPath)

	pingRouterV1.GET("", pingHandler.PingHandler())
}

/*
	ObjectPayment
		* CheckPayment
			* Deps => ObjectEmail

	# Metode 1

	Ketika object diinisiasi, buat object email saat inisiasi

	class GMailObject:
		def init(self):
			bla bla bla

	class PangMailObject:
		def init(self):
			bla bla bla

	class Payment:
		def init(self):
			self.mailObject = new PangMailObject()

		def CheckPayment(account):
			fsdafsdafad

	# Metode 2 => Dependency Injection

	class GMailObject:
		def init(self):
			bla bla bla

	class PangMailObject:
		def init(self):
			bla bla bla

	class Payment:
		def init(self, mailObject):
			self.mailObject = mailObject

		def CheckPayment(account):
			fsdafsdafad

	class Main:
		def init(self):
			self.mailObject = new fdasfasdfasd()
			self.payment = new Payment(self.mailObject)

*/

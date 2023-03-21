package url

import (
	"github.com/GilarYa/web-rps/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Homepage)
	// page.Get("/presensi", controller.GetPresensiBulanIni)
	page.Get("/dosen", controller.GetDosen)         //API from user whatsapp message from iteung gowa
	page.Get("/about", controller.Getabout)         //API from user whatsapp message from iteung gowa
	page.Get("/contacus", controller.Getcontacus)   //API from user whatsapp message from iteung gowa
	page.Get("/dashboard", controller.Getdashboard) //API from user whatsapp message from iteung gowa
	page.Get("/mahasiswa", controller.Getmahasiswa) //API from user whatsapp message from iteung gowa

	page.Get("/", controller.Sink)
	page.Post("/", controller.Sink)
	page.Put("/", controller.Sink)
	page.Patch("/", controller.Sink)
	page.Delete("/", controller.Sink)
	page.Options("/", controller.Sink)

}

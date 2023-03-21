package controller

import (
	"github.com/GilarYa/ckbackend"
	"github.com/GilarYa/web-rps/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/whatsauth/whatsauth"
)

var DataDosen = "dosen"

type HTTPRequest struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}

func Sink(c *fiber.Ctx) error {
	var req HTTPRequest
	req.Header = string(c.Request().Header.Header())
	req.Body = string(c.Request().Body())
	return c.JSON(req)
}

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}
func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}
func GetDosen(c *fiber.Ctx) error {
	getstatus := ckbackend.GetDataDosen("DOSEN")
	return c.JSON(getstatus)
}
func Getabout(c *fiber.Ctx) error {
	getstatus := ckbackend.GetDataAbout("Apakah kurikulum ini menyulitkan hidup?")
	return c.JSON(getstatus)
}
func Getcontacus(c *fiber.Ctx) error {
	getstatus := ckbackend.GetDataContacus("0822126722")
	return c.JSON(getstatus)
}
func Getdashboard(c *fiber.Ctx) error {
	getstatus := ckbackend.GetDataDashboard("Denmark")
	return c.JSON(getstatus)
}
func Getmahasiswa(c *fiber.Ctx) error {
	getstatus := ckbackend.GetDataMahasiswa("johndoe@flex.co")
	return c.JSON(getstatus)
}

// func GetPresensiBulanIni(c *fiber.Ctx) error {
// 	ps := presensi.GetPresensiCurrentMonth(config.Ulbimongoconn)
// 	return c.JSON(ps)
// }

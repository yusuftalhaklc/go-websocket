package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// Soket bağlantılarını tutmak için bir harita kullanıyoruz.
var connections = make(map[*websocket.Conn]bool)

func main() {
	app := fiber.New()

	// WebSocket endpoint'i oluşturuyoruz.
	app.Get("/ws", websocket.New(handleWebSocket))
	app.Get("/", func(c *fiber.Ctx) error {
		// index.html dosyasını okuyup cevap olarak döndürüyoruz
		return c.SendFile("index.html")
	})
	// Web uygulamasını servis etmek için basit bir endpoint ekleyebiliriz.
	app.Static("/", "./static")

	log.Fatal(app.Listen(":3000"))
}

func handleWebSocket(c *websocket.Conn) {
	// Soket bağlantısını haritaya ekliyoruz.
	connections[c] = true
	defer func() {
		// Soket bağlantısını kapattığımızda haritadan kaldırıyoruz.
		delete(connections, c)
		c.Close()
	}()

	for {
		// Soket üzerinden gelen verileri alıyoruz.
		msgType, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Soket üzerinden mesaj okunurken bir hata oluştu:", err)
			break
		}

		// Gelen mesajı haritadaki tüm bağlantılara gönderiyoruz.
		for conn := range connections {
			err = conn.WriteMessage(msgType, msg)
			if err != nil {
				log.Println("Mesaj gönderilirken bir hata oluştu:", err)
				conn.Close()
				delete(connections, conn)
			}
		}
	}
}

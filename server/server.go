package server

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

const (
	maxSize int = 1024 * 1024 * 50
)

func getFileContent(fileName string) ([]byte, int, error) {
	fd, err := os.Open(fileName)
	if err != nil {
		return []byte{}, 0, errors.New("file not found")
	}

	defer fd.Close()

	reader := bufio.NewReader(fd)
	info, _ := fd.Stat()
	data := make([]byte, info.Size()) // BUG: get correct size here
	count, err := io.ReadFull(reader, data)
	if err != nil {
		return []byte{}, 0, errors.New("file read failed")
	}

	log.Printf("bytes read = %d", count)
	return data, count, nil
}

// StartServer starts the media streaming server.
func StartServer() {
	app := fiber.New()

	app.Use(middleware.Logger())

	app.Get("/media", func(c *fiber.Ctx) {
		c.Send("Hello World")
	})

	app.Static("/", "./public")
	app.Listen(80)
}

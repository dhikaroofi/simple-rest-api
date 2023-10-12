package restApi

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Task interface {
	Start() error
	Shutdown() error
}

type server struct {
	app *fiber.App
}

func NewFiberServer() Task {
	app := fiber.New(fiber.Config{})
	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(recover.New())

	return &server{
		app: app,
	}
}

func (s *server) Start() error {
	if err := s.app.Listen(fmt.Sprintf(":%s", "1000")); err != nil {
		// ErrServerClosed is expected behaviour when exiting app
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("server is closed caused by: %s", err.Error())
		}

		return err
	}

	return nil
}

func (s *server) Shutdown() error {
	if err := s.app.Shutdown(); err != nil {
		return err
	}

	log.Println("http server is stopped")
	return nil
}

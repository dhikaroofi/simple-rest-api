package restApi

import (
	"errors"
	"fmt"
	"github.com/dhikaroofi/simple-rest-api/internal/presentation/restApi/common"
	"github.com/dhikaroofi/simple-rest-api/internal/usecase"
	validator2 "github.com/dhikaroofi/simple-rest-api/pkg/validator"
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
	app       *fiber.App
	validator *validator2.ValidationEngine
	useCase   *usecase.Container

	appPort string
}

func NewFiberServer(appPort string, container *usecase.Container) Task {
	app := fiber.New(fiber.Config{
		ErrorHandler: common.ErrResponse,
	})

	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(recover.New())

	validatorEngine, err := validator2.NewEngine()
	if err != nil {
		log.Fatalf("failed to set up validator")
	}

	return &server{
		app:       app,
		appPort:   appPort,
		useCase:   container,
		validator: validatorEngine,
	}
}

func (s *server) Start() error {
	s.route()
	if err := s.app.Listen(fmt.Sprintf(":%s", s.appPort)); err != nil {
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

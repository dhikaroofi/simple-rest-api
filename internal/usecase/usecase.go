package usecase

import (
	"database/sql"
	"gorm.io/gorm"
	"log"

	"github.com/dhikaroofi/simple-rest-api/internal/config"
	"github.com/dhikaroofi/simple-rest-api/internal/usecase/employee"
	gorm2 "github.com/dhikaroofi/simple-rest-api/pkg/gorm"
)

type Container struct {
	dbClient *gorm.DB
	sqlDB    *sql.DB
	Employee employee.ServicesInterfaces
}

func NewUseCase(conf *config.Config) *Container {
	dbClient, sqlDB, err := gorm2.NewGorm(conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Pass, conf.DB.Database)
	if err != nil {
		log.Fatalf("failed to connect to database | err: %s", err.Error())
	}

	return &Container{
		dbClient: dbClient,
		sqlDB:    sqlDB,
		Employee: employee.NewEmployeeServices(dbClient),
	}
}

func (c Container) StoppingAdapters() error {
	if err := c.sqlDB.Close(); err != nil {
		return err
	}

	return nil
}

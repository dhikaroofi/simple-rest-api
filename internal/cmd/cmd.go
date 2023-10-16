package cmd

import (
	"github.com/dhikaroofi/simple-rest-api/internal/config"
	restapi "github.com/dhikaroofi/simple-rest-api/internal/presentation/restApi"
	"github.com/dhikaroofi/simple-rest-api/internal/usecase"
	"log"
)

func Init(appExitChan chan bool) {
	var (
		conf   *config.Config
		cont   *usecase.Container
		server *restapi.Server
	)

	conf = config.LoadConfigFromFile("resources/config/config.yaml")
	cont = usecase.NewUseCase(conf)
	server = restapi.NewFiberServer(conf.AppPort, cont)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("failed to start the server caused by: %s", err.Error())
		}
	}()

	go func() {
		<-appExitChan
		if err := server.Shutdown(); err != nil {
			log.Printf("failed to stop the server caused by: %s \n", err.Error())
		}

		if err := cont.StoppingAdapters(); err != nil {
			log.Printf("failed to stop the server caused by: %s \n", err.Error())
		}

		log.Println("all adapters has been shutdown")
		appExitChan <- true
	}()
}

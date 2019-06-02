package integrationtest

import (
	"github.com/state-alchemists/ayanami/msgbroker"
	"github.com/state-alchemists/ayanami/service"
	"log"
)

// MainServiceHTML emulating HTML's main
func MainServiceHTML() {
	serviceName := "html"
	// define broker
	broker, err := msgbroker.NewNats()
	if err != nil {
		log.Fatal(err)
	}
	// define services
	services := service.Services{
		"pre": service.NewService(serviceName, "pre",
			[]string{"input"},
			[]string{"output"},
			WrappedPre,
		),
	}
	// consume and publish forever
	ch := make(chan bool)
	services.ConsumeAndPublish(broker, serviceName)
	<-ch
}
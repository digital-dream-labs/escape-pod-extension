package main

import (
	"github.com/digital-dream-labs/escape-pod-extension/examples/go/pkg/battery"
	"github.com/digital-dream-labs/escape-pod-extension/examples/go/pkg/genericresponse"
	"github.com/digital-dream-labs/escape-pod-extension/examples/go/pkg/openweathermap"
	"github.com/digital-dream-labs/escape-pod-extension/examples/go/pkg/selfaware"
	"github.com/digital-dream-labs/escape-pod-extension/examples/go/pkg/server"
	"github.com/digital-dream-labs/escape-pod-extension/proto/lang/go/podextension"
	grpclog "github.com/digital-dream-labs/hugh/grpc/interceptors/log"
	grpcserver "github.com/digital-dream-labs/hugh/grpc/server"
	"github.com/digital-dream-labs/hugh/log"
)

func getProcessors() (map[string]server.Processor, error) {
	// init weather
	weather, err := openweathermap.New(
		openweathermap.WithViper(),
	)
	if err != nil {
		return nil, err
	}

	// init selfaware
	ih, err := selfaware.New(
		selfaware.WithViper(),
	)
	if err != nil {
		return nil, err
	}

	b, err := battery.New(
		battery.WithViper(),
	)
	if err != nil {
		return nil, err
	}

	home, err := genericresponse.New(
		genericresponse.WithViper(),
		genericresponse.WithResponse(
			"I'm talking with the escape pod.  Stop talking to me and go write more plugins.",
		),
	)
	if err != nil {
		return nil, err
	}

	return map[string]server.Processor{
			"intent_weather_extend":        weather,
			"intent_custom_hack":           ih,
			"intent_custom_hack_battery":   b,
			"intent_custom_hack_escapepod": home,
		},
		nil
}

func main() {
	log.SetJSONFormat("2006-01-02 15:04:05")

	srv, err := grpcserver.New(
		grpcserver.WithViper(),
		grpcserver.WithLogger(log.Base()),
		grpcserver.WithReflectionService(),

		grpcserver.WithUnaryServerInterceptors(
			grpclog.UnaryServerInterceptor(),
		),

		grpcserver.WithStreamServerInterceptors(
			grpclog.StreamServerInterceptor(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	p, err := getProcessors()
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.New(
		server.WithProcessors(p),
	)
	if err != nil {
		log.Fatal(err)
	}

	podextension.RegisterPodExtensionServer(srv.Transport(), s)

	srv.Start()

	<-srv.Notify(grpcserver.Stopped)

}

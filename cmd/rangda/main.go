package main

import (
	"net/http"

	"github.com/alexflint/go-arg"
	"github.com/openware/rangda/pkg/log"
	"github.com/reconquest/karma-go"
)

var (
	// this variable is changed by runtime ldflags
	version = "[manual build]"
)

type args struct {
	Config string
	Debug  bool
}

func (args *args) Version() string {
	return version
}

func initArgs() *args {
	args := &args{}

	// Fill default values
	args.Config = "rangda.conf"

	// parses command line
	arg.MustParse(args)

	return args
}

func main() {
	args := initArgs()

	if args.Debug {
		log.SetDebug(true)
	}

	config, err := LoadConfig(args.Config)
	if err != nil {
		log.Fatalf(err, "unable to load config: %s", args.Config)
	}

	log.Infof(nil, "loaded configuration file: %s", args.Config)

	server := NewServer(config)

	server.SetupRoutes()

	log.Infof(
		karma.Describe("address", config.Address).Describe("version", version),
		"starting listener",
	)

	err = http.ListenAndServe(config.Address, server)
	if err != nil {
		log.Errorf(err, "unable to start listener at %s", config.Address)
	}
}

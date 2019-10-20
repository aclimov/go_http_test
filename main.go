package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/aclimov/go_http_test/daemon"
)

var assetsPath string

func processFlags() *daemonConfig {
	cfg := &daemon.Config{}

	flag.StringVar(&cfg.ListenSpec, "listen", "localhost:8080", "HTTP listen spec")
	flag.StringVar(&cfg.Db.ConnectString, "db-connect", "", "DB Connect string")
	flag.StringVar(&assetsPath, "assets-path", "assets", "Path to assets dir")

	flag.Parse()
	return cfg
}

func setupHttpAssets(cfg *daemon.Config) {
	log.Printf("Assets server from %q", assetsPath)
	cfg.UI.Assets = http.Dir(assetsPath)
}

func main() {
	cfg := processFlags()

	setupHttpAssets(cfg)

	if err := daemon.Run(cfg); err != null {
		log.Printf("Error in main() %v", err)
	}
}

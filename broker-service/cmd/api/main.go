package main

import (
	"log"
	"net/http"
)

const webPort = "8000"

type Config struct {}

func main() {

  app := Config{}
  log.Println("Starting on port", webPort)

  // define web server
  server := &http.Server{
	Addr: ":" + webPort,
	Handler: app.routes(),
  }

  // start web server
  err := server.ListenAndServe()
  if err != nil {
	log.Fatal(err)
  }

}

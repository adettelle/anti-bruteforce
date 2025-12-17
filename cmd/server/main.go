package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adettelle/anti-bruteforce/config"
	internalhttp "github.com/adettelle/anti-bruteforce/internal/server/http"
)

func main() {
	fmt.Println("start")

	cfg := config.New()

	server := internalhttp.NewServer(cfg)

	err := server.Srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("server failed: ", err)
	}
}

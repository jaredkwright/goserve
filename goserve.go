package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 9999, "port number")
	flag.Parse()
	portString := fmt.Sprintf(":%d", *port)
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Simple static webserver:
	fmt.Printf("Serving files on port %d\n", *port)
	log.Fatal(http.ListenAndServe(portString, http.FileServer(http.Dir(pwd))))
}

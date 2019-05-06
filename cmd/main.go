package main

import (
	"github.com/martinrocket/tengine"
	"os"
	"log"
	"net/http"
	"fmt"
)

func main(){
	log.Print("Hello world sample started.")

        http.HandleFunc("/", tengine.Handle)

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}


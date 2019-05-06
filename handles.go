package tengine

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

func Handle(w http.ResponseWriter, r *http.Request){
	log.Print("Hello world received a request.")
	target := os.Getenv("TARGET")
        if target == "" {
                target = "World"
        }
        fmt.Fprintf(w, "Hello %s!\n", target)
}



package main

import (
    "fmt"
    "log"
    "net/http"
	"multinodal-cache/handlers"
)

func main() {
    http.HandleFunc("/app", AppHandler)

    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}

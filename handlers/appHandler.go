package handlers

import (
    "fmt"
    "net/http"
)

func AppHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the query parameters
    param1 := r.URL.Query().Get("param_1")

    if param1 == "" {
        http.Error(w, "Missing parameter", http.StatusBadRequest)
        return
    }


    response := fmt.Sprintf("Received parameters: param_1 = %s" , param1)
    fmt.Fprintf(w, response)
}
package handlers

import (
    "fmt"
    "net/http"
)

func AppHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the query parameters
    param1 := r.URL.Query().Get("param_1")
    param2 := r.URL.Query().Get("param_2")

    if param1 == "" || param2 == "" {
        http.Error(w, "Missing parameters", http.StatusBadRequest)
        return
    }

    response := fmt.Sprintf("Received parameters: param_1 = %s, param_2 = %s", param1, param2)
    fmt.Fprintf(w, response)
}
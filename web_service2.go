package main

import (
    "fmt"
    "net/http"
)

// Handler untuk endpoint pertama (/hello)
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "helo ini web service saya")
}

// Handler untuk endpoint kedua (/info)
func infoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "fadil nugroho")
}

func main() {
    // Registrasi handler untuk masing-masing endpoint
    http.HandleFunc("/halow", helloHandler)
    http.HandleFunc("/infong", infoHandler)

    // Mulai server web pada port 8080
    fmt.Println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}


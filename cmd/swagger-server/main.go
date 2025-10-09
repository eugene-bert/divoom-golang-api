package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eugene-bert/divoom-golang-api/swagger"
)

func main() {
	// Swagger UI endpoint
	http.HandleFunc("/swagger", swagger.Handler())

	// OpenAPI spec endpoint
	http.HandleFunc("/openapi.yaml", swagger.SpecHandler())

	// Root redirect to swagger
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/swagger", http.StatusFound)
			return
		}
		http.NotFound(w, r)
	})

	port := ":8080"
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║      Divoom PIXOO64 API Documentation Server              ║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  Swagger UI:         http://localhost%s/swagger       ║\n", port)
	fmt.Printf("║  OpenAPI Spec:       http://localhost%s/openapi.yaml  ║\n", port)
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Press Ctrl+C to stop the server")

	log.Fatal(http.ListenAndServe(port, nil))
}

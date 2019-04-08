// Main entry point
package main

// Use standard library "net/http" to handle HTTP request
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go web development")
	})

	// Listen to port 5000
	fmt.Println(http.ListenAndServe(":5000", nil))
}

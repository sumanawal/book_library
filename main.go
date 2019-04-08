// Main entry point
package main

// Use standard library "net/http" to handle HTTP request
import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Page struct {
	Name     string
	DBStatus bool
}

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "go_library"
)

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Name: "Suman"}
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}
		p.DBStatus = db.Ping() == nil
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Listen to port 5000
	fmt.Println(http.ListenAndServe(":5000", nil))
}

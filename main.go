package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KanhaiyaKumarGupta/assignment/databases"
	"github.com/KanhaiyaKumarGupta/assignment/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"
)

func main() {
	err := databases.InitDatabase()
	if err != nil {
		log.Fatal(err)
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	routes.TaskRoutes(r)
    
	fmt.Println(http.ListenAndServe(":3000", r))
}

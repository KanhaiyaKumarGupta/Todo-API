package routes

import (
	"github.com/KanhaiyaKumarGupta/assignment/controllers"
	"github.com/go-chi/chi/v5"
)

func TaskRoutes(router *chi.Mux) {
	router.Post("/tasks/createtask", controllers.CreateTask)
	router.Get("/tasks/{taskID}", controllers.GetTask)
	router.Put("/tasks/{taskID}", controllers.UpdateTask)
	router.Delete("/tasks/{taskID}", controllers.DeleteTask)
}

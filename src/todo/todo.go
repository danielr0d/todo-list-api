package todo

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Config struc{
	*configs.Config
}

func New(configuration *configs.Config) *Config{
	return &Config{configuration}
}

func (config *Config) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoID}", config.GetATodo)
	router.Delete("/{todoID}", config.DeleteTodo)
	router.Post("/{todoID}", config.CreateTodo)
	router.Get("/", config.GetAllTodos)
	return router

}

type Todo struct {
	Slug string `json:"slug"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func (config *Config) GetATodo (w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todos := Todo{
		Slug: todoID,
		Title: "My first Todo on Go",
		Body: "Hello world from planet Neptune",
	}
	render.JSON(w, r, todos)
}

func (config *Config) DeleteTodo (w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "TODO deleted successfully"
	render.JSON(w, r, respose)
}

func (config *Config) CreateTodo (w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "TODO created successfully"
	render.JSON(w, r, respose)
}

func (config *Config) GetAllTodos(w http.ResponseWriter, r* http.Request){
	todos := []Todo{
		{
			Slug: "slug",
			Title:"My first Todo on Go",
			Body: "Hello world from planet Makonha",
		},
	}
	render.JSON(w, r, todos)
}
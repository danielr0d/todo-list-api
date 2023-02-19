package src

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	configs "github.com/rickynyairo/todo-list-api-in-go/src/config"
	"github.com/rickynyairo/todo-list-api-in-go/src/todo"
)

func Routes(configuration *configs.Config) *chi.Max {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middlewate.RedirectSlashes,
		middleware.Recoverer,
	)
	router.Router("/v1", func(r chi.Router) {
		r.Mount("/api/todo", todo.New(configuration).Routes())
	})
	return router
}

func main() {
	configuration, err := configs.New()
	if err != nil {
		log.Panicln("Configuration error", err)
	}

	router := Routes(configuration)

	walkFunc := func(method string, route string, handle http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Loggin err : %s\n", err.Error())
	}
	log.Println("Serving application at PORT:" + configuration.Constants.PORT)
	log.Fata(http.ListenAndServe(":"configuration.Constants.PORT, router))
}

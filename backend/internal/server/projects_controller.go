package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ProjectsController struct{}

func projectContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		projectID := chi.URLParam(r, "projectID")

		ctx := context.WithValue(r.Context(), projectIDContextKey, projectID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func pageContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageID := chi.URLParam(r, "pageID")

		ctx := context.WithValue(r.Context(), pageIDContextKey, pageID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (c *ProjectsController) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", c.List)
	r.Get("/search", c.Search)

	r.Post("/", c.Create)

	r.Route("/{projectID}", func(r chi.Router) {
		r.Use(projectContext)
		r.Get("/", c.Get)
		r.Put("/", c.Update)
		r.Delete("/", c.Delete)

		r.Route("/pages", func(r chi.Router) {
			r.Get("/", c.ListPages)

			r.Get("/search", c.SearchPages)

			r.Route("/{pageID}", func(r chi.Router) {
				r.Use(pageContext)
				r.Get("/", c.GetPage)
			})
		})
	})
	return r
}

func (c *ProjectsController) RouteGroup() string {
	return "projects"
}

func (c *ProjectsController) List(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("list projects")); err != nil {
		panic(err)
	}
}

func (c *ProjectsController) Search(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("search projects")); err != nil {
		panic(err)
	}
}

func (c *ProjectsController) Create(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("create project")); err != nil {
		panic(err)
	}
}

func (c *ProjectsController) Get(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("get project: " + r.Context().Value(projectIDContextKey).(string))); err != nil {
		panic(err)
	}
}

func (c *ProjectsController) Update(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("update project: " + r.Context().Value(projectIDContextKey).(string))); err != nil {
		panic(err)
	}
}

func (c *ProjectsController) Delete(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("delete project: " + r.Context().Value(projectIDContextKey).(string))); err != nil {
		panic(err)
	}
}

func (c *ProjectsController) GetPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf(
		"get project (%s) page: %s",
		r.Context().Value(projectIDContextKey).(string),
		r.Context().Value(pageIDContextKey).(string),
	)))

	if err != nil {
		panic(err)
	}
}

func (c *ProjectsController) SearchPages(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf(
		"search project (%s) pages",
		r.Context().Value(projectIDContextKey).(string),
	)))

	if err != nil {
		panic(err)
	}

}

func (c *ProjectsController) ListPages(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf(
		"list project (%s) pages",
		r.Context().Value(projectIDContextKey).(string),
	)))

	if err != nil {
		panic(err)
	}

}

func NewProjectsController() *ProjectsController {
	return &ProjectsController{}
}
package handlers

import (
	"github.com/vladyslavpavlenko/go-web-course/pkg/config"
	"github.com/vladyslavpavlenko/go-web-course/pkg/models"
	"github.com/vladyslavpavlenko/go-web-course/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

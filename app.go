package app

import (
	"TestProj/internal/actions"
	"TestProj/internal/actions/handler"
	"TestProj/pkg/config"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = actions.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/books", a.GetAllBooks)
	a.Post("/book", a.CreateBook)
	a.Get("book/{id}", a.GetBookByPage)
	a.Get("/book", a.GetBook)
	a.Put("/book", a.UpdateBook)
	a.Delete("/books", a.DeleteBook)
	a.Put("/books/{id}/disable", a.DisableBook)
	a.Put("/books/{id}/enable", a.EnableBook)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Book Data
func (a *App) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	handler.GetAllBooks(a.DB, w, r)
}

func (a *App) CreateBook(w http.ResponseWriter, r *http.Request) {
	handler.CreateBook(a.DB, w, r)
}

func (a *App) GetBook(w http.ResponseWriter, r *http.Request) {
	handler.GetBook(a.DB, w, r)
}

func (a *App) GetBookByPage(w http.ResponseWriter, r *http.Request) {
	handler.GetBookByPage(a.DB, w, r)
}

func (a *App) UpdateBook(w http.ResponseWriter, r *http.Request) {
	handler.UpdateBook(a.DB, w, r)
}

func (a *App) DeleteBook(w http.ResponseWriter, r *http.Request) {
	handler.DeleteBook(a.DB, w, r)
}

func (a *App) DisableBook(w http.ResponseWriter, r *http.Request) {
	handler.DisableBook(a.DB, w, r)
}

func (a *App) EnableBook(w http.ResponseWriter, r *http.Request) {
	handler.EnableBook(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
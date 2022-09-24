package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/FilipSolich/ci-server/configs"
	"github.com/FilipSolich/ci-server/db"
	"github.com/FilipSolich/ci-server/handlers"
	"github.com/FilipSolich/ci-server/middlewares"
	"github.com/FilipSolich/ci-server/models"
	"github.com/FilipSolich/ci-server/services"
)

func initDatabase() {
	var err error
	db.DB, err = gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	db.DB.AutoMigrate(&models.User{}, &models.OAuth2Token{}, &models.Webhook{})
}

func initGitServices() {
	if configs.GitHubService {
		services.NewGitHubOAuth2Config(configs.GitHubClientID, configs.GitHubClientSecret)
	}

	// TODO: Add GitLab service.
}

func initTemplates() {
	configs.LoadTemplates()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = configs.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	initTemplates()
	initDatabase()
	initGitServices()

	CSRF := csrf.Protect([]byte(configs.CSRFSecret))

	r := mux.NewRouter()
	r.Use(CSRF)
	r.Use(middlewares.LoggingMiddleware)
	r.Handle("/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.Index)))
	r.HandleFunc("/login", handlers.Login)
	r.HandleFunc("/logout", handlers.Logout)

	sOAuth2 := r.PathPrefix("/oauth2").Subrouter()
	sOAuth2.HandleFunc("/callback", handlers.OAuth2Callback)

	sRepos := r.PathPrefix("/repositories").Subrouter()
	sRepos.Use(middlewares.AuthMiddleware)
	sRepos.HandleFunc("", handlers.Repos)
	sRepos.HandleFunc("/register", handlers.ReposRegister).Methods(http.MethodPost)
	sRepos.HandleFunc("/unregister", handlers.ReposUnregister).Methods(http.MethodPost)
	sRepos.HandleFunc("/activate", handlers.ReposActivate).Methods(http.MethodPost)
	sRepos.HandleFunc("/deactivate", handlers.ReposDeactivate).Methods(http.MethodPost)

	server := &http.Server{
		Addr:         "" + configs.Port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Println("Server running on " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

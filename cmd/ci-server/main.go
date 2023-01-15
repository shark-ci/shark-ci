package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/shark-ci/shark-ci/ci-server/configs"
	"github.com/shark-ci/shark-ci/ci-server/db"
	"github.com/shark-ci/shark-ci/ci-server/handlers"
	"github.com/shark-ci/shark-ci/ci-server/middlewares"
	"github.com/shark-ci/shark-ci/ci-server/services"
	"github.com/shark-ci/shark-ci/ci-server/store"
	"github.com/shark-ci/shark-ci/mq"
)

func initGitServices() {
	if configs.GitHubEnabled {
		services.NewGitHubManager(configs.GitHubClientID, configs.GitHubClientSecret)
		services.Services[services.GitHub.GetServiceName()] = &services.GitHub
	}
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
	initGitServices()

	mongoStore, err := store.NewMongoStore(configs.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoStore.Close(context.TODO())

	disconnectDB, err := db.InitDatabase(configs.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
	defer disconnectDB(context.Background())

	closeMQ, err := mq.InitMQ(configs.RabbitMQHost, configs.RabbitMQPort, configs.RabbitMQUsername, configs.RabbitMQPassword)
	if err != nil {
		log.Fatal(err)
	}
	defer closeMQ()

	CSRF := csrf.Protect([]byte(configs.CSRFSecret))

	loginHandler := handlers.NewLoginHandler(mongoStore)

	r := mux.NewRouter()
	r.Use(middlewares.LoggingMiddleware)
	r.Handle("/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.IndexHandler)))
	r.HandleFunc("/login", loginHandler.Login)
	r.HandleFunc("/logout", handlers.LogoutHandler)
	r.HandleFunc(configs.EventHandlerPath+"/{service}", handlers.EventHandler).Methods(http.MethodPost)

	// OAuth2 subrouter.
	OAuth2 := r.PathPrefix("/oauth2").Subrouter()
	OAuth2.HandleFunc("/callback", handlers.OAuth2CallbackHandler)

	// Repositories subrouter.
	repos := r.PathPrefix("/repositories").Subrouter()
	repos.Use(CSRF)
	repos.Use(middlewares.AuthMiddleware)
	repos.HandleFunc("", handlers.ReposHandler)
	repos.HandleFunc("/register", handlers.ReposRegisterHandler).Methods(http.MethodPost)
	repos.HandleFunc("/unregister", handlers.ReposUnregisterHandler).Methods(http.MethodPost)
	repos.HandleFunc("/activate", handlers.ReposActivateHandler).Methods(http.MethodPost)
	repos.HandleFunc("/deactivate", handlers.ReposDeactivateHandler).Methods(http.MethodPost)

	// Jobs subrouter.
	jobs := r.PathPrefix(configs.JobsPath).Subrouter()
	jobs.Handle("/{id}", middlewares.AuthMiddleware(http.HandlerFunc(handlers.JobsTargetHandler)))
	jobs.HandleFunc(configs.JobsReportStatusHandlerPath+"/{id}", handlers.JobsReportStatusHandler).Methods(http.MethodPost)
	jobs.HandleFunc(configs.JobsPublishLogsHandlerPath+"/{id}", handlers.JobsPublishLogsHandler).Methods(http.MethodPost)

	server := &http.Server{
		Addr:    ":" + configs.Port,
		Handler: r,
		//ReadTimeout:  15 * time.Second,
		//WriteTimeout: 15 * time.Second,
		//IdleTimeout:  60 * time.Second,
	}
	log.Println("Server running on " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

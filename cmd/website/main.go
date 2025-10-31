package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/waduhek/website/internal"
	eduHandler "github.com/waduhek/website/internal/education/handler"
	expHandler "github.com/waduhek/website/internal/experience/handler"
	homeHandler "github.com/waduhek/website/internal/home/handler"
	projHandler "github.com/waduhek/website/internal/projects/handler"
	"github.com/waduhek/website/internal/telemetry"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGHUP,
	)
	defer cancel()

	shutdownOTel, err := telemetry.SetupOTelSDK(context.Background())
	if err != nil {
		slog.Error("error while setting up otel sdk", "error", err)
		os.Exit(1)
	}

	dependencies := internal.BuildDependencies(internal.TemplateNameFileMap)

	mux := http.NewServeMux()

	mux.Handle(
		"GET /static/css/",
		http.StripPrefix(
			"/static/css/",
			http.FileServer(http.Dir("static/css")),
		),
	)
	mux.Handle(
		"GET /static/js/",
		http.StripPrefix(
			"/static/js/",
			http.FileServer(http.Dir("static/js")),
		),
	)

	expHandler := expHandler.NewExperienceHandler(dependencies)
	mux.Handle("GET /experience", expHandler)

	eduHandler := eduHandler.NewEducationHandler(dependencies)
	mux.Handle("GET /education", eduHandler)

	projHandler := projHandler.NewProjectsHandler(dependencies)
	mux.Handle("GET /projects", projHandler)

	homeHandler := homeHandler.NewHomeHandler(dependencies)
	mux.Handle("GET /{$}", homeHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go server.ListenAndServe()

	<-ctx.Done()

	dependencies.Logger.Info("shutting down server")

	server.Shutdown(context.Background())
	dependencies.DbConn.Close()
	shutdownOTel(context.Background())
}

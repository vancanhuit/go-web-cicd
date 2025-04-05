package main

import (
	"encoding/json"
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	Logger *slog.Logger
}

func (app *application) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.Logger.Info("incoming request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (app *application) helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	//nolint:errcheck
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello from Go!",
	})
}

func (app *application) router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/hello", app.helloHandler)
	return mux
}

func main() {
	var addr string

	flag.StringVar(&addr, "addr", ":8080", "HTTP address")
	flag.Parse()

	app := &application{
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})),
	}

	server := &http.Server{
		Handler: app.loggerMiddleware(app.router()),
		Addr:    addr,
	}

	app.Logger.Info("starting server", "addr", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		app.Logger.Error("failed to start server", "error", err.Error())
		os.Exit(1)
	}
}

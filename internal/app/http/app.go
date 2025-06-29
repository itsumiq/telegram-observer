package httpapp

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"telegram-observer/internal/presentation/http/handler"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	srv  *http.Server
	addr string
	port int
	log  *slog.Logger
}

func New(addr string, port int, telegramHandler handler.Telegram, userHandler handler.User, log *slog.Logger) *App {
	r := mux.NewRouter()
	r.HandleFunc("/telegram/webhook", telegramHandler.HandlePostWebhook).Methods(http.MethodPost)
	r.HandleFunc("/users/{user_id}/photos", userHandler.HandlePostPhoto).Methods(http.MethodPost)
	r.HandleFunc("users/{user_id}/videos", userHandler.HandlePostVideo).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", addr, port),
		Handler: r,
	}
	return &App{srv: srv, log: log}
}

func (a *App) Run() {
	a.log.Info(fmt.Sprintf("server started on %s:%d", a.addr, a.port))
	log.Fatal(a.srv.ListenAndServe())
}

func (a *App) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		a.log.Error("server shutdown error", slog.String("error", err.Error()))
		return err
	}

	return nil
}

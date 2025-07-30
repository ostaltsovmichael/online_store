package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"online_shop/internal/configs"
)

func Server(router http.Handler) {
	var cfg = configs.Cfg

	port := cfg.Port
	host := cfg.Host

	slog.Info(fmt.Sprintf("Server started on port: %d", port))

	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router)
}

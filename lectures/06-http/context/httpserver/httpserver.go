package httpserver

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func runServer() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
		<-stop
		cancel()
	}()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler{},
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	return srv.ListenAndServe()
}

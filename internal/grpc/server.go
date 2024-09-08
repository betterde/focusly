package grpc

import (
	"errors"
	"github.com/betterde/focusly/internal/gen/svc/v1/svcv1connect"
	"github.com/betterde/focusly/internal/grpc/services"
	"github.com/betterde/focusly/internal/journal"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

var ServerInstance *http.Server

func Serve() {
	mux := http.NewServeMux()
	ServerInstance = &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	userService := &services.UserService{}

	path, handler := svcv1connect.NewUserServiceHandler(userService)
	mux.Handle(path, handler)

	go func() {
		if err := ServerInstance.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				journal.Logger.Info(err.Error())
			} else {
				journal.Logger.Error(err)
			}
		}
	}()
}

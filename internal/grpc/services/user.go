package services

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	svcv1 "github.com/betterde/focusly/internal/gen/svc/v1"
	"github.com/betterde/focusly/internal/journal"
)

type UserService struct{}

func (svc *UserService) SignInWithSSOProvider(ctx context.Context, req *connect.Request[svcv1.SignInWithSSOProviderRequest]) (*connect.Response[svcv1.SignInWithSSOProviderResponse], error) {
	journal.Logger.Infof("req: %+v", req)
	res := connect.NewResponse(&svcv1.SignInWithSSOProviderResponse{
		Payload: fmt.Sprintf("Your username is %s", req.Msg.Username),
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}
func (svc *UserService) Echo(ctx context.Context, req *connect.Request[svcv1.EchoRequest]) (*connect.Response[svcv1.EchoResponse], error) {
	res := connect.NewResponse(&svcv1.EchoResponse{
		Payload: req.Msg.Payload,
	})
	res.Header().Set("User-Version", "v1")
	return res, nil
}

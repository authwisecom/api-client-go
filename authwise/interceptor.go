package authwise

import (
	"context"
	"github.com/authwisecom/api-client-go/credentials/bearer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewDefaultBearerCredentialsInterceptor() grpc.UnaryClientInterceptor {
	return NewBearerCredentialsInterceptor(bearer.NewDefaultService())
}

func NewBearerCredentialsInterceptor(service bearer.Service) grpc.UnaryClientInterceptor {

	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		token, err := service.Get(ctx)
		if err != nil {
			return err
		}
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", "Bearer "+token))
		return invoker(ctx, method, req, reply, cc, opts...)

	}
}
